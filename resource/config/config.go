package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"strings"
)

type Config struct {
	Memory     int64
	Cpu        int64
	CpuSet     []string
	CgroupName string
}

func NewConfig(c *cli.Context) Config {
	var res Config
	res.Memory = c.Int64("mem")
	res.Cpu = c.Int64("cpu")
	cpuSetStr := c.String("cpuset")
	res.CpuSet = strings.Fields(cpuSetStr)
	log.Infof("@NewConfig gen a config: %#v", res)
	return res
}
