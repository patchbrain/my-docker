package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"strings"
)

type Config struct {
	Memory     *int64
	Cpu        *int64
	CpuSet     []string
	CgroupName string
}

func NewConfig(c *cli.Context) Config {
	var res Config
	mem := c.Int64("mem")
	if mem != 0{
		res.Memory = &mem
	}

	cpu := c.Int64("cpu")
	if cpu != 0{
		res.Cpu = &cpu
	}

	if c.String("cpuset") != ""{
		res.CpuSet = strings.Fields(c.String("cpuset"))
	}

	log.Infof("@NewConfig gen a config: %#v", res)
	return res
}
