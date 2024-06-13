package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	config2 "mydocker/config"
	"mydocker/container"
	"mydocker/pkg/commit"
	"mydocker/pkg/resource/config"
	"path/filepath"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgroup limit: mydocker run -it [command]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
		cli.Int64Flag{
			Name:  "mem",
			Usage: "memory limit(Byte), eg. -mem 100",
		},
		cli.Int64Flag{
			Name:  "cpu",
			Usage: "cpu time limit, eg. -cpu 100",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpu set, eg. -cpuset 2,3,4",
		},
		cli.StringFlag{
			Name:  "v",
			Usage: "volume, eg. -v host/source/path:container/dst/path",
		},
	},
	Action: func(c *cli.Context) error {
		log.Infof("into run command, args: %#v", c.Args())
		if len(c.Args()) == 0 {
			return fmt.Errorf("missing command")
		}
		cmd := c.Args()     // 获取要在容器中执行的命令
		tty := c.Bool("it") // 获取是否有-it这个参数
		volStr := c.String("v")
		rCfg := config.NewConfig(c)
		rCfg.CgroupName = "my_docker_cg"
		Run(tty, cmd, rCfg, volStr)
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

var commitCommand = cli.Command{
	Name:  "commit",
	Usage: "commit your change.",
	Action: func(c *cli.Context) error {
		log.Infof("into commit command...")
		args := c.Args()
		tarName := args.Get(0)
		log.Infof("commit command, first arg: %s",tarName)
		cmt := commit.NewTarCommitter(tarName, filepath.Join(config2.MntPath, "merge"))
		err := cmt.Commit()
		return err
	},
}
