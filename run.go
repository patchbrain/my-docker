package main

import (
	log "github.com/sirupsen/logrus"
	"mydocker/container"
	"mydocker/resource/config"
	"os"
	"strings"
)

func Run(tty bool, cmd []string, cfg config.Config) {
	// 组装一个通过init包装cmd的命令
	parent, wp := container.NewParentProcess(tty)
	if err := parent.Start(); err != nil {
		log.Errorf("@Run parent cmd failed: %s", err.Error())
	}

	// 将命令从writePipe中发送
	err := writeArgs2Pipe(wp, cmd)
	if err != nil {
		log.Errorf("@Run write args 2 pipe failed: %s, cmd: %#v", err.Error(), cmd)
	}

	if err := parent.Wait(); err != nil {
		log.Errorf("@Run parent wait failed: %s", err.Error())
	}
	os.Exit(-1)
}

func writeArgs2Pipe(writePipe *os.File, cmd []string) error {
	log.Infof("@writeArgs2Pipe start, cmd: %#v",cmd)
	cmdStr := strings.Join(cmd, " ")
	_, err := writePipe.WriteString(cmdStr)
	if err != nil {
		return err
	}
	log.Infof("@writeArgs2Pipe write 2 pipe end, cmd: %#v",cmd)
	writePipe.Close()

	return nil
}
