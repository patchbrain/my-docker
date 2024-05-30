package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"mydocker/container"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgroup limit: mydocker run -it [command]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
	},
	Action: func(c *cli.Context) error {
		log.Infof("into run command, args: %#v",c.Args())
		if len(c.Args()) == 0 {
			return fmt.Errorf("missing command")
		}
		cmd := c.Args() // 获取要在容器中执行的命令
		tty := c.Bool("it")    // 获取是否有-it这个参数
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "init container. don't call it outside",
	Action: func(c *cli.Context) error {
		log.Infof("into init command...")
		err := container.Initer()
		return err
	},
}
