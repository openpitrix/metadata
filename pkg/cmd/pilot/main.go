// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// openpitrix pilot service app package.
package pilot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/urfave/cli"

	"openpitrix.io/logger"
	"openpitrix.io/metadata/pkg/internal/pathutil"
	"openpitrix.io/metadata/pkg/pb/drone"
	"openpitrix.io/metadata/pkg/pb/types"
	"openpitrix.io/metadata/pkg/pilot"
	"openpitrix.io/metadata/pkg/pilot/pilotutil"
	"openpitrix.io/metadata/pkg/version"
)

func Main() {
	app := cli.NewApp()
	app.Name = "pilot"
	app.Usage = "pilot provides pilot service."
	app.Version = version.GetVersionString()

	app.UsageText = `pilot [global options] command [options] [args...]

EXAMPLE:
   pilot gen-config
   pilot info
   pilot list
   pilot ping
   pilot exec
   pilot getv key
   pilot confd-info
   pilot confd-start
   pilot serve
   pilot send-task
   pilot tour`

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Value:  "pilot-config.json",
			Usage:  "pilot config file",
			EnvVar: "OPENPITRIX_PILOT_CONFIG",
		},

		cli.StringFlag{
			Name:   "pilot-server-name",
			Value:  "pilot.openpitrix.io",
			Usage:  "pilot server name",
			EnvVar: "OPENPITRIX_PILOT_SERVER_NAME",
		},
		cli.StringFlag{
			Name:   "pilot-server-crt-file",
			Value:  "tls-config/pilot-server.crt",
			Usage:  "pilot server tls crt file",
			EnvVar: "OPENPITRIX_PILOT_SERVER_CRT_FILE",
		},
		cli.StringFlag{
			Name:   "pilot-server-key-file",
			Value:  "tls-config/pilot-server.key",
			Usage:  "pilot server tls key file",
			EnvVar: "OPENPITRIX_PILOT_SERVER_KEY_FILE",
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
				fmt.Println(pilot.NewDefaultConfigString())
				return
			},
		},

		{
			Name:  "info",
			Usage: "show pilot service info",

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				info, err := client.GetPilotConfig(context.Background(), &pbtypes.Empty{})
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
			Usage:     "list frontgate nodes",
			ArgsUsage: "[regexp]",

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				list, err := client.GetFrontgateList(context.Background(), &pbtypes.Empty{})
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
			Usage: "ping pilot/frontgate/drone service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "endpoint-type",
					Value: "pilot",
					Usage: "set endpoint type (pilot/frontgate/metad/drone)",
				},

				cli.StringFlag{
					Name:  "frontgate-id",
					Value: "frontgate-001",
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
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				switch s := c.String("endpoint-type"); s {
				case "pilot":
					_, err = client.PingPilot(context.Background(), &pbtypes.Empty{})
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					fmt.Println("OK")
					return

				case "frontgate":
					_, err = client.PingFrontgate(context.Background(), &pbtypes.FrontgateId{
						Id: c.String("frontgate-id"),
					})
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					fmt.Println("OK")
					return

				case "metad":
					_, err = client.PingMetadataBackend(context.Background(), &pbtypes.FrontgateId{
						Id: c.String("frontgate-id"),
					})
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					fmt.Println("OK")
					return

				case "drone":
					_, err = client.PingDrone(
						context.Background(),
						&pbtypes.DroneEndpoint{
							FrontgateId: c.String("frontgate-id"),
							DroneIp:     c.String("drone-host"),
							DronePort:   int32(c.Int("drone-port")),
						},
					)
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					fmt.Println("OK")
					return

				default:
					logger.Criticalf(nil, "unknown endpoint type: %s\n", s)
					os.Exit(1)
				}

				return
			},
		},

		{
			Name:  "exec",
			Usage: "exec command on pilot/frontgate/drone service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "endpoint-type",
					Value: "frontgate",
					Usage: "set endpoint type (frontgate/drone)",
				},

				cli.StringFlag{
					Name:  "frontgate-id",
					Value: "frontgate-001",
				},
				cli.StringFlag{
					Name:  "frontgate-node-id",
					Value: "frontgate-node-001",
				},
				cli.StringFlag{
					Name:  "drone-host",
					Value: "localhost",
				},
				cli.IntFlag{
					Name:  "drone-port",
					Value: int(pbdrone.Default_Default_DroneServicePort),
				},
				cli.IntFlag{
					Name:  "timeout",
					Value: 3,
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				switch s := c.String("endpoint-type"); s {
				case "frontgate":
					_, err = client.RunCommandOnFrontgateNode(context.Background(), &pbtypes.RunCommandOnFrontgateRequest{
						Endpoint: &pbtypes.FrontgateEndpoint{
							FrontgateId:     c.String("frontgate-id"),
							FrontgateNodeId: c.String("frontgate-node-id"),
						},
						Command:        strings.Join(c.Args(), " "),
						TimeoutSeconds: int32(c.Int("timeout")),
					})
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					fmt.Println("OK")
					return

				case "drone":
					_, err = client.RunCommandOnDrone(
						context.Background(),
						&pbtypes.RunCommandOnDroneRequest{
							Endpoint: &pbtypes.DroneEndpoint{
								FrontgateId: c.String("frontgate-id"),
								DroneIp:     c.String("drone-host"),
								DronePort:   int32(c.Int("drone-port")),
							},
							Command:        strings.Join(c.Args(), " "),
							TimeoutSeconds: int32(c.Int("timeout")),
						},
					)
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					fmt.Println("OK")
					return

				default:
					logger.Criticalf(nil, "unknown endpoint type: %s\n", s)
					os.Exit(1)
				}

				return
			},
		},

		{
			Name:  "confd-status",
			Usage: "get confd service status",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "frontgate-id",
					Value: "frontgate-001",
				},
				cli.StringFlag{
					Name:  "drone-host",
					Value: "",
				},
				cli.IntFlag{
					Name:  "drone-port",
					Value: int(pbdrone.Default_Default_DroneServicePort),
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				reply, err := client.IsConfdRunning(context.Background(), &pbtypes.DroneEndpoint{
					FrontgateId: c.String("frontgate-id"),
					DroneIp:     c.String("drone-host"),
					DronePort:   int32(c.Int("drone-port")),
				})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				if reply.GetValue() {
					fmt.Printf("confd on frontgate(%s)/drone(%s:%d) is running\n",
						c.String("frontgate-id"), c.String("drone-host"), c.Int("drone-port"),
					)
				} else {
					fmt.Printf("confd on frontgate(%s)/drone(%s:%d) not running\n",
						c.String("frontgate-id"), c.String("drone-host"), c.Int("drone-port"),
					)
				}

				return
			},
		},

		{
			Name:  "confd-info",
			Usage: "get confd service config",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "frontgate-id",
					Value: "frontgate-001",
				},
				cli.StringFlag{
					Name:  "drone-host",
					Value: "",
				},
				cli.IntFlag{
					Name:  "drone-port",
					Value: int(pbdrone.Default_Default_DroneServicePort),
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				reply, err := client.GetConfdConfig(context.Background(), &pbtypes.ConfdEndpoint{
					FrontgateId: c.String("frontgate-id"),
					DroneIp:     c.String("drone-host"),
					DronePort:   int32(c.Int("drone-port")),
				})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				fmt.Println(jsonString(reply))
				return
			},
		},

		{
			Name:  "confd-start",
			Usage: "start confd service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "frontgate-id",
					Value: "frontgate-001",
				},
				cli.StringFlag{
					Name:  "drone-host",
					Value: "",
				},
				cli.IntFlag{
					Name:  "drone-port",
					Value: int(pbdrone.Default_Default_DroneServicePort),
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				_, err = client.StartConfd(context.Background(), &pbtypes.DroneEndpoint{
					FrontgateId: c.String("frontgate-id"),
					DroneIp:     c.String("drone-host"),
					DronePort:   int32(c.Int("drone-port")),
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
			Name:  "confd-stop",
			Usage: "stop confd service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "frontgate-id",
					Value: "frontgate-001",
				},
				cli.StringFlag{
					Name:  "drone-host",
					Value: "",
				},
				cli.IntFlag{
					Name:  "drone-port",
					Value: int(pbdrone.Default_Default_DroneServicePort),
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				_, err = client.StopConfd(context.Background(), &pbtypes.DroneEndpoint{
					FrontgateId: c.String("frontgate-id"),
					DroneIp:     c.String("drone-host"),
					DronePort:   int32(c.Int("drone-port")),
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
			Name:  "get-cmd-status",
			Usage: "get cmd status",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "task-id",
					Value: "",
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				status, err := client.GetSubtaskStatus(context.Background(), &pbtypes.SubTaskId{
					TaskId: c.String("task-id"),
				})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				fmt.Println(status.Status)
				return
			},
		},
		{
			Name:  "send-task",
			Usage: "send task to pilot service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "task-file",
					Value: "task.json",
				},
			},

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				// load task json
				task := func() *pbtypes.SubTaskMessage {
					data, err := ioutil.ReadFile(c.String("task-file"))
					if err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					p := new(pbtypes.SubTaskMessage)
					if err := json.Unmarshal(data, p); err != nil {
						logger.Criticalf(nil, "%+v", err)
						os.Exit(1)
					}
					return p
				}()

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				_, err = client.HandleSubtask(context.Background(), task)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				fmt.Println("Done")
				return
			},
		},

		{
			Name:  "serve",
			Usage: "run as pilot service",

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				pbPilotTLSConfig, err := pilotutil.LoadPilotTLSConfig(
					c.GlobalString("pilot-server-crt-file"),
					c.GlobalString("pilot-server-key-file"),
					c.GlobalString("pilot-client-crt-file"),
					c.GlobalString("pilot-client-key-file"),
					c.GlobalString("openpitrix-ca-crt-file"),
					c.GlobalString("pilot-server-name"),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				pilot.Serve(cfg, pbPilotTLSConfig)
				return
			},
		},

		{
			Name:  "get-version",
			Usage: "get service version",

			Action: func(c *cli.Context) {
				cfgpath := pathutil.MakeAbsPath(c.GlobalString("config"))
				cfg := pilotutil.MustLoadPilotConfig(cfgpath)

				client, conn, err := pilotutil.DialPilotService(
					context.Background(), cfg.Host, int(cfg.ListenPort),
				)
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}
				defer conn.Close()

				info, err := client.GetPilotVersion(context.Background(), &pbtypes.Empty{})
				if err != nil {
					logger.Criticalf(nil, "%+v", err)
					os.Exit(1)
				}

				fmt.Println(jsonString(info))
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
pilot gen-config
pilot gen-cert

pilot info
pilot list

pilot getv /
pilot getv /key
pilot getv / /key

pilot confd-start
pilot confd-stop
pilot confd-status

pilot serve

GOOS=windows pilot list
LIBCONFD_GOOS=windows pilot list
`
