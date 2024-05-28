package main

import (
	log "github.com/sirupsen/logrus"
	"mydocker/container"
	"os"
)

func Run(tty bool, cmd string) {
	// 组装一个通过init包装cmd的命令
	parent := container.NewParentProcess(tty, cmd)
	if err := parent.Start(); err != nil {
		log.Errorf("@Run parent cmd failed: %s", err.Error())
	}

	if err := parent.Wait(); err != nil {
		log.Errorf("@Run parent wait failed: %s", err.Error())
	}
	os.Exit(-1)
}
