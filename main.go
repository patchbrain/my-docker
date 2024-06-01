package main

import (
	"github.com/urfave/cli"
	"os"
	"path"
	"runtime"
	"strings"
)
import log "github.com/sirupsen/logrus"

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = `simple docker`

	app.Commands = []cli.Command{runCommand,initCommand}

	app.Before = func(c *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		})
		log.SetReportCaller(true)

		log.SetOutput(os.Stdout)
		return nil
	}

	if err:=app.Run(os.Args); err!=nil{
		log.Fatalf("app-mydocker run failed: %s",err.Error())
	}
}