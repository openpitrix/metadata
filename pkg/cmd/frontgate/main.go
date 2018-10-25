// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// openpitrix frontgate service app package.
package frontgate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/chai2010/jsonmap"
	"github.com/urfave/cli"

	"openpitrix.io/logger"
	"openpitrix.io/metadata/pkg/frontgate"
	"openpitrix.io/metadata/pkg/frontgate/frontgateutil"
	"openpitrix.io/metadata/pkg/internal/pathutil"
	"openpitrix.io/metadata/pkg/internal/tlsutil"
	"openpitrix.io/metadata/pkg/pb/drone"
	"openpitrix.io/metadata/pkg/pb/types"
	"openpitrix.io/metadata/pkg/version"
)

func Main() {
	app := cli.NewApp()
	app.Name = "frontgate"
	app.Usage = "frontgate provides frontgate service."
	app.Version = version.GetVersionString()

	app.UsageText = `frontgate [global options] command [options] [args...]

EXAMPLE:
   frontgate gen-config
   frontgate info
   frontgate list
   frontgate ping
   frontgate getv key
   frontgate setv key value
   frontgate confd-start
   frontgate serve
   frontgate tour`

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Value:  "frontgate-config.json",
			Usage:  "frontgate config file",
			EnvVar: "OPENPITRIX_FRONTGATE_CONFIG",
		},

		cli.StringFlag{
			Name:   "pilot-server-name",
			Value:  "pilot.openpitrix.io",
			Usage:  "pilot server name",
			EnvVar: "OPENPITRIX_PILOT_SERVER_NAME",
		},
		cli.StringFlag{
			Name:   "pilot-client-crt-file",
			Value:  "tls-config/pilot-client.crt",
			Usage:  "pilot client tls crt file",
			EnvVar: "OPENPITRIX_PILOT_CLIENT_CRT_FILE",
		},
		cli.StringFlag{
			Name:   "pilot-client-key-file",
			Value:  "tls-config/pilot-client.key",
			Usage:  "pilot client tls key file",
			EnvVar: "OPENPITRIX_PILOT_CLIENT_KEY_FILE",
		},

		cli.StringFlag{
			Name:   "openpitrix-ca-crt-file",
			Value:  "tls-config/openpitrix-ca.crt",
			Usage:  "openpitrix ca crt file",
			EnvVar: "OPENPITRIX_CA_CRT_FILE",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "gen-config",
			Usage: "gen default config",

			Action: func(c *cli.Context) {
				fmt.Println(frontgate.NewDefaultConfigString())
				return
			},
		},

		{
			Name:  "info",
			Usage: "show frontgate service info",

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				client, err := frontgateutil.DialFrontgateService(
					cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer client.Close()

				info, err := client.GetFrontgateConfig(&pbtypes.Empty{})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				fmt.Println(jsonString(info))
				return
			},
		},
		{
			Name:      "list",
			Usage:     "list drone nodes",
			ArgsUsage: "[regexp]",

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				client, err := frontgateutil.DialFrontgateService(
					cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer client.Close()

				list, err := client.GetDroneList(&pbtypes.Empty{})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				re := c.Args().First()
				for _, v := range list.GetIdList() {
					if re == "" {
						fmt.Println(v)
						continue
					}
					matched, err := regexp.MatchString(re, v)
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					if matched {
						fmt.Println(v)
					}
				}
				return
			},
		},
		{
			Name:  "ping",
			Usage: "ping frontgate/pilot/drone service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "endpoint-type",
					Value: "frontgate",
					Usage: "set endpoint type (pilot/frontgate/drone)",
				},

				cli.StringFlag{
					Name:  "drone-host",
					Value: "localhost",
				},
				cli.IntFlag{
					Name:  "drone-port",
					Value: int(pbdrone.Default_Default_DroneServicePort),
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				client, err := frontgateutil.DialFrontgateService(
					cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer client.Close()

				switch s := c.String("endpoint-type"); s {
				case "frontgate":
					_, err = client.PingFrontgate(&pbtypes.Empty{})
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					fmt.Println("OK")
					return
				case "pilot":
					_, err = client.PingPilot(&pbtypes.Empty{})
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					fmt.Println("OK")

				case "drone":
					_, err = client.PingDrone(&pbtypes.DroneEndpoint{
						FrontgateId: cfg.Id,
						DroneIp:     c.String("drone-host"),
						DronePort:   int32(c.Int("drone-port")),
					})
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					fmt.Println("OK")

				default:
					logger.Criticalf(nil, "unknown endpoint type: %s\n", s)
					os.Exit(1)
				}

				return
			},
		},

		{
			Name:      "getv",
			Usage:     "get values from etcd by keys",
			ArgsUsage: "key",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "prefix",
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				if c.NArg() == 0 {
					logger.Criticalf(nil, "missing value")
					os.Exit(1)
				}

				client, err := frontgateutil.DialFrontgateService(
					cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer client.Close()

				var reply *pbtypes.StringMap
				if !c.Bool("prefix") {
					reply, err = client.GetEtcdValues(&pbtypes.StringList{
						ValueList: c.Args(),
					})
				} else {
					reply, err = client.GetEtcdValuesByPrefix(&pbtypes.String{
						Value: c.Args().First(),
					})
				}
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				var maxLen = 1
				for k, _ := range reply.GetValueMap() {
					if len(k) > maxLen {
						maxLen = len(k)
					}
				}

				for k, v := range reply.GetValueMap() {
					fmt.Printf("%-*s => %s\n", maxLen, k, v)
				}

				return
			},
		},
		{
			Name:      "setv",
			Usage:     "set value to etcd",
			ArgsUsage: "key [value | file]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "value-is-file",
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				client, err := frontgateutil.DialFrontgateService(
					cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer client.Close()

				if c.NArg() < 2 {
					logger.Criticalf(nil, "missing args")
					os.Exit(1)
				}

				kvMap := map[string]string{}

				if c.Bool("value-is-file") {
					data, err := ioutil.ReadFile(c.Args().Get(1))
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}

					var datMap map[string]interface{}
					if err := json.Unmarshal(data, &datMap); err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}

					kvMap = jsonmap.NewJsonMapFromKV(datMap, "/").ToMapString("/")

				} else {
					kvMap[c.Args().First()] = c.Args().Get(1)
				}

				_, err = client.SetEtcdValues(&pbtypes.StringMap{
					ValueMap: kvMap,
				})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				fmt.Println("Done")
				return
			},
		},

		{
			Name:  "confd-status",
			Usage: "get confd service status",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "drone-host",
					Value: "localhost",
				},
				cli.IntFlag{
					Name:  "drone-port",
					Value: int(pbdrone.Default_Default_DroneServicePort),
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				client, err := frontgateutil.DialFrontgateService(
					cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer client.Close()

				reply, err := client.IsConfdRunning(&pbtypes.ConfdEndpoint{
					DroneIp:   c.String("drone-host"),
					DronePort: int32(c.Int("drone-port")),
				})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				if reply.GetValue() {
					fmt.Printf("confd on drone(%s:%d) is running\n",
						c.String("drone-host"), c.Int("drone-port"),
					)
				} else {
					fmt.Printf("confd on drone(%s:%d) not running\n",
						c.String("drone-host"), c.Int("drone-port"),
					)
				}

				return
			},
		},
		{
			Name:  "confd-start",
			Usage: "start confd service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "drone-host",
					Value: "localhost",
				},
				cli.IntFlag{
					Name:  "drone-port",
					Value: int(pbdrone.Default_Default_DroneServicePort),
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				client, err := frontgateutil.DialFrontgateService(
					cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer client.Close()

				_, err = client.StartConfd(&pbtypes.ConfdEndpoint{
					DroneIp:   c.String("drone-host"),
					DronePort: int32(c.Int("drone-port")),
				})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				return
			},
		},
		{
			Name:  "confd-stop",
			Usage: "stop confd service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "drone-host",
					Value: "localhost",
				},
				cli.IntFlag{
					Name:  "drone-port",
					Value: int(pbdrone.Default_Default_DroneServicePort),
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				client, err := frontgateutil.DialFrontgateService(
					cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer client.Close()

				_, err = client.StopConfd(&pbtypes.ConfdEndpoint{
					DroneIp:   c.String("drone-host"),
					DronePort: int32(c.Int("drone-port")),
				})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				return
			},
		},

		{
			Name:  "report-cmd-status",
			Usage: "report cmd status",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "task-id",
					Value: "",
				},
				cli.StringFlag{
					Name:  "task-status",
					Value: "",
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				client, err := frontgateutil.DialFrontgateService(
					cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer client.Close()

				_, err = client.ReportSubTaskStatus(&pbtypes.SubTaskStatus{
					TaskId: c.String("task-id"),
					Status: c.String("task-status"),
				})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				return
			},
		},

		{
			Name:  "serve",
			Usage: "run as frontgate service",
			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := frontgateutil.MustLoadFrontgateConfig(cfgpath)

				tlsPilotClientConfig, err := tlsutil.NewClientTLSConfigFromFile(
					c.GlobalString("pilot-client-crt-file"),
					c.GlobalString("pilot-client-key-file"),
					c.GlobalString("openpitrix-ca-crt-file"),
					c.GlobalString("pilot-server-name"),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)

				}

				frontgate.Serve(
					frontgate.NewConfigManager(cfgpath, cfg), tlsPilotClientConfig,
				)
				return
			},
		},
		{
			Name:  "tour",
			Usage: "show more examples",
			Action: func(c *cli.Context) {
				fmt.Println(tourTopic)
			},
		},
	}

	app.CommandNotFound = func(ctx *cli.Context, command string) {
		fmt.Fprintf(ctx.App.Writer, "not found '%v'!\n", command)
	}

	app.Run(os.Args)
}

func jsonString(m interface{}) string {
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return ""
	}
	data = bytes.Replace(data, []byte("\n"), []byte("\r\n"), -1)
	return string(data)
}

func atoi(s string, defaultValue int) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	return defaultValue
}

const tourTopic = `
frontgate gen-config

frontgate info
frontgate list

frontgate getv /
frontgate getv /key
frontgate getv / /key

frontgate confd-start
frontgate confd-stop
frontgate confd-status

frontgate serve

GOOS=windows frontgate list
LIBCONFD_GOOS=windows frontgate list
`
