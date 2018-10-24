// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package drone

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"

	"openpitrix.io/libconfd"
	"openpitrix.io/logger"
	"openpitrix.io/metadata/pkg/drone/yunify_confdfunc"
	"openpitrix.io/metadata/pkg/internal/funcutil"
	"openpitrix.io/metadata/pkg/pb/drone"
	"openpitrix.io/metadata/pkg/pb/types"
	"openpitrix.io/metadata/pkg/version"
)

var (
	_                      pbdrone.DroneServiceServer = (*Server)(nil)
	DowloadPathPattern                                = "/opt/openpitrix/bin/%s"
	DowloadFilePathPattern                            = "/opt/openpitrix/bin/%s/%s"
	PilotVersionFilePath                              = "/opt/openpitrix/conf/pilot-version"
)

func (p *Server) createPilotVersionFile(pilotVersion string) error {
	f, err := os.Create(PilotVersionFilePath)
	if err != nil {
		return err
	}

	_, err = f.WriteString(pilotVersion)
	if err != nil {
		return err
	}

	return nil
}

func (p *Server) DistributeDrone(ctx context.Context, in *pbtypes.DroneBinary) (*pbtypes.Empty, error) {
	pilotVersion := in.GetVersion().GetValue()
	droneBinary := in.GetDrone().GetValue()

	err := os.MkdirAll(fmt.Sprintf(DowloadPathPattern, pilotVersion), os.ModeDir|os.ModePerm)
	if err != nil {
		return &pbtypes.Empty{}, err
	}

	filePath := fmt.Sprintf(DowloadFilePathPattern, pilotVersion, "drone")

	logger.Infof(nil, "Write drone to [%s]", filePath)
	f, err := os.Create(filePath)
	if err != nil {
		return &pbtypes.Empty{}, err
	}

	_, err = f.Write(droneBinary)
	if err != nil {
		return &pbtypes.Empty{}, err
	}

	err = os.Chmod(filePath, os.ModePerm)
	if err != nil {
		return &pbtypes.Empty{}, err
	}

	err = p.createPilotVersionFile(pilotVersion)
	if err != nil {
		return &pbtypes.Empty{}, err
	}

	os.Exit(0)

	return &pbtypes.Empty{}, nil
}

func (p *Server) GetPilotVersion(context.Context, *pbtypes.Empty) (*pbtypes.Version, error) {
	err := fmt.Errorf("TODO")
	logger.Warnf(nil, "%+v", err)
	return nil, err
}
func (p *Server) GetFrontgateVersion(context.Context, *pbtypes.Empty) (*pbtypes.Version, error) {
	err := fmt.Errorf("TODO")
	logger.Warnf(nil, "%+v", err)
	return nil, err
}
func (p *Server) GetDroneVersion(context.Context, *pbtypes.DroneEndpoint) (*pbtypes.Version, error) {
	reply := &pbtypes.Version{
		ShortVersion:   version.ShortVersion,
		GitSha1Version: version.GitSha1Version,
		BuildDate:      version.BuildDate,
	}
	return reply, nil
}

func (p *Server) PingMetadataBackend(context.Context, *pbtypes.FrontgateEndpoint) (*pbtypes.Empty, error) {
	err := fmt.Errorf("TODO")
	logger.Warnf(nil, "%+v", err)
	return nil, err
}

func (p *Server) GetDroneConfig(context.Context, *pbtypes.Empty) (*pbtypes.DroneConfig, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	return p.cfg.Get(), nil
}
func (p *Server) SetDroneConfig(ctx context.Context, cfg *pbtypes.DroneConfig) (*pbtypes.Empty, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	if reflect.DeepEqual(cfg, p.cfg.Get()) {
		return &pbtypes.Empty{}, nil
	}

	if err := p.cfg.Set(cfg); err != nil {
		logger.Warnf(nil, "%+v", err)
		return &pbtypes.Empty{}, err
	}

	if err := p.cfg.Save(); err != nil {
		logger.Warnf(nil, "%+v", err)
		return &pbtypes.Empty{}, err
	}

	return &pbtypes.Empty{}, nil
}

func (p *Server) GetConfdConfig(ctx context.Context, arg *pbtypes.Empty) (*pbtypes.ConfdConfig, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	return p.confd.GetConfig(), nil
}

func (p *Server) SetConfdConfig(ctx context.Context, arg *pbtypes.ConfdConfig) (*pbtypes.Empty, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	cfg := p.cfg.Get()
	fnHookKeyAdjuster := func(absKey string) (realKey string) {
		if absKey == "/self" {
			return "/" + cfg.ConfdSelfHost
		}
		if strings.HasPrefix(absKey, "/self/") {
			return "/" + cfg.ConfdSelfHost + absKey[len("/self/")-1:]
		}
		return absKey
	}

	if err := p.confd.SetConfig(arg, fnHookKeyAdjuster); err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}
	if err := p.confd.Save(); err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	return &pbtypes.Empty{}, nil
}

func (p *Server) GetFrontgateConfig(context.Context, *pbtypes.Empty) (*pbtypes.FrontgateConfig, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	cfg, err := p.fg.GetConfig()
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}
	return cfg, nil
}

func (p *Server) SetFrontgateConfig(ctx context.Context, cfg *pbtypes.FrontgateConfig) (*pbtypes.Empty, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	err := p.fg.SetConfig(cfg)
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}
	return &pbtypes.Empty{}, nil
}

func (p *Server) IsConfdRunning(ctx context.Context, arg *pbtypes.Empty) (*pbtypes.Bool, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	return &pbtypes.Bool{Value: p.confd.IsRunning()}, nil
}

func (p *Server) StartConfd(ctx context.Context, arg *pbtypes.Empty) (*pbtypes.Empty, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	cfg := p.cfg.Get()
	logger.Infof(nil, "StartConfd: %v", cfg)

	err := p.confd.Start(func(opt *libconfd.Config) {
		opt.FuncMap = yunify_confdfunc.MakeCustomFuncMap()
		opt.HookAbsKeyAdjuster = func(absKey string) (realKey string) {
			if absKey == "/"+cfg.ConfdSelfHost || strings.HasPrefix(absKey, "/"+cfg.ConfdSelfHost+"/") {
				return absKey
			}

			if absKey == "/self" {
				return "/" + cfg.ConfdSelfHost
			}
			if strings.HasPrefix(absKey, "/self/") {
				return "/" + cfg.ConfdSelfHost + absKey[len("/self/")-1:]
			} else {
				return "/" + cfg.ConfdSelfHost + absKey
			}
		}
		opt.HookOnCheckCmdDone = func(trName, cmd string, err error) {
			if err != nil {
				logger.Warnf(nil, "%+v", err)
				return
			}
			if trName == "/etc/confd/conf.d/cmd.info.toml" {
				go func() {
					logger.Infof(nil, "LoadLastCmdStatus: %v", cfg.CmdInfoLogPath)

					status, isEmpty, err := LoadLastCmdStatus(cfg.CmdInfoLogPath)
					if isEmpty {
						return
					}
					if err != nil {
						logger.Warnf(nil, "%+v", err)
						p.fg.ReportSubTaskStatus(&pbtypes.SubTaskStatus{
							TaskId: status.SubtaskId,
							Status: pbtypes.TaskStatus_failed.String(),
						})
					} else {
						p.fg.ReportSubTaskStatus(&pbtypes.SubTaskStatus{
							TaskId: status.SubtaskId,
							Status: status.Status,
						})
					}
				}()
			}
		}
		opt.HookOnReloadCmdDone = func(trName, cmd string, err error) {
			if err != nil {
				logger.Warnf(nil, "%+v", err)
				return
			}
			if trName == "/etc/confd/conf.d/cmd.info.toml" {
				go func() {
					status, isEmpty, err := LoadLastCmdStatus(cfg.CmdInfoLogPath)
					if isEmpty {
						return
					}
					if err != nil {
						logger.Warnf(nil, "%+v", err)
						p.fg.ReportSubTaskStatus(&pbtypes.SubTaskStatus{
							TaskId: status.SubtaskId,
							Status: pbtypes.TaskStatus_failed.String(),
						})
					} else {
						p.fg.ReportSubTaskStatus(&pbtypes.SubTaskStatus{
							TaskId: status.SubtaskId,
							Status: status.Status,
						})
					}
				}()
			}
		}
		opt.HookOnUpdateDone = func(trName string, err error) {
			if err != nil {
				logger.Warnf(nil, "%+v", err)
				return
			}
		}
	})

	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	return &pbtypes.Empty{}, nil
}

func (p *Server) StopConfd(ctx context.Context, arg *pbtypes.Empty) (*pbtypes.Empty, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	if err := p.confd.Stop(); err != nil {
		logger.Errorf(nil, "StopConfd: %v", err)
		return nil, err
	}

	return &pbtypes.Empty{}, nil
}

func (p *Server) GetTemplateFiles(ctx context.Context, arg *pbtypes.Empty) (*pbtypes.StringList, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	if !p.confd.IsRunning() {
		err := fmt.Errorf("drone: confd is not started")
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	cfg := p.confd.GetConfig()
	confdir := cfg.GetProcessorConfig().GetConfdir()
	if confdir == "" {
		err := fmt.Errorf("drone: invaid confdir: %q", confdir)
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	_, paths, err := libconfd.ListTemplateResource(filepath.Join(confdir, "conf.d"))
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	reply := &pbtypes.StringList{}
	for _, s := range paths {
		reply.ValueList = append(reply.ValueList, filepath.Base(s))
	}

	return reply, nil
}

func (p *Server) GetValues(ctx context.Context, arg *pbtypes.StringList) (*pbtypes.StringMap, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	if !p.confd.IsRunning() {
		err := fmt.Errorf("drone: confd is not started")
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	client := p.confd.GetBackendClient()
	if client == nil {
		return nil, fmt.Errorf("drone: confd is not started")
	}

	keys := arg.GetValueList()
	m, err := client.GetValues(keys)
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	reply := &pbtypes.StringMap{
		ValueMap: make(map[string]string),
	}

	for _, k := range keys {
		reply.ValueMap[k] = m[k]
	}

	return reply, nil
}

func (p *Server) PingPilot(ctx context.Context, arg *pbtypes.FrontgateEndpoint) (*pbtypes.Empty, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	c, err := p.fg.getClient()
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	_, err = c.PingPilot(&pbtypes.Empty{})
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	return &pbtypes.Empty{}, nil
}

func (p *Server) PingFrontgate(ctx context.Context, arg *pbtypes.FrontgateEndpoint) (*pbtypes.Empty, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	c, err := p.fg.getClient()
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	_, err = c.PingFrontgate(&pbtypes.Empty{})
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	return &pbtypes.Empty{}, nil
}

func (p *Server) PingDrone(ctx context.Context, arg *pbtypes.Empty) (*pbtypes.Empty, error) {
	logger.Infof(nil, funcutil.CallerName(1))
	return &pbtypes.Empty{}, nil
}

func (p *Server) RunCommand(ctx context.Context, arg *pbtypes.RunCommandOnDroneRequest) (*pbtypes.String, error) {
	logger.Infof(nil, funcutil.CallerName(1))

	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/C", arg.GetCommand())
	} else {
		c = exec.Command("/bin/sh", "-c", arg.GetCommand())
	}

	var b bytes.Buffer
	c.Stdout = &b
	c.Stderr = &b

	if err := c.Start(); err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	var timeout = time.Second * 3
	if x := arg.GetTimeoutSeconds(); x > 0 {
		timeout = time.Duration(x) * time.Second
	}

	timer := time.AfterFunc(timeout, func() {
		c.Process.Kill()
	})

	err := c.Wait()
	timer.Stop()

	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	return &pbtypes.String{Value: string(b.Bytes())}, nil
}
