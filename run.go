package main

import (
	log "github.com/sirupsen/logrus"
	"mydocker/container"
	"mydocker/pkg/resource"
	"mydocker/pkg/resource/config"
	"mydocker/pkg/resource/subsystem"
	"os"
	"strings"
)

func Run(tty bool, cmd []string, cfg config.Config, volStr string) {
	// 组装一个通过init包装cmd的命令
	parent, wp, ol := container.NewParentProcess(tty, volStr)
	if parent == nil {
		log.Errorf("@Run gen parent cmd error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Errorf("@Run parent cmd failed: %s", err.Error())
	}

	// 进行Cgroup配置的应用
	log.Infof("start to config CGroup..")
	memSub := subsystem.NewMemSys(&cfg)
	cpuSub := subsystem.NewCpuSys(&cfg)
	cpuSetSub := subsystem.NewCpuSetSys(&cfg)
	resourceMgr := resource.MgrIns().Register(memSub, cpuSub, cpuSetSub)
	resourceMgr.Apply()

	container.SetEndFn(func() error {
		err := resourceMgr.Destroy()
		if err != nil {
			log.Errorf("destroy cgroup failed: %s", err.Error())
			return err
		}

		err = ol.UnMount()
		if err != nil {
			log.Errorf("unmount failed: %s", err.Error())
			return err
		}

		return nil
	})

	// 将命令从writePipe中发送
	err := writeArgs2Pipe(wp, cmd)
	if err != nil {
		log.Errorf("@Run write args 2 pipe failed: %s, cmd: %#v", err.Error(), cmd)
	}

	if tty{
		if err := parent.Wait(); err != nil {
			log.Errorf("@Run parent wait failed: %s", err.Error())
		}

		err = container.EndFn()
		if err!=nil{
			log.Errorf("@Run run end func failed: %s", err.Error())
		}
	}
}

func writeArgs2Pipe(writePipe *os.File, cmd []string) error {
	log.Infof("@writeArgs2Pipe start, cmd: %#v", cmd)
	cmdStr := strings.Join(cmd, " ")
	_, err := writePipe.WriteString(cmdStr)
	if err != nil {
		return err
	}
	log.Infof("@writeArgs2Pipe write 2 pipe end, cmd: %#v", cmd)
	writePipe.Close()

	return nil
}
