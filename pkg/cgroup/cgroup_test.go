package cgroup

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSetSpec(t *testing.T) {
	tests := []struct {
		Name   string
		Config CgroupOpCfg
	}{
		{Name: "mem set",
			Config: CgroupOpCfg{
				CgroupName: "test_cg",
				SubSysName: "memory",
				SpecName:   "memory.limit_in_bytes",
				Value:      "10000000",
				AutoCreate: true,
			}},
		{Name: "pid set",
			Config: CgroupOpCfg{
				CgroupName: "test_cg",
				SubSysName: "memory",
				SpecName:   "tasks",
				Value:      "9090909",
				AutoCreate: true,
			}},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := SetSpec(tt.Config)
			if err != nil {
				logrus.Errorf("%s: %s", tt.Name, err.Error())
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct{
		Name string
		Cfg CgroupOpCfg
	}{
		{
			Name: "delete",
			Cfg:  CgroupOpCfg{
				CgroupName: "test_cg",
				SubSysName: "memory",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			err := Delete(tt.Cfg)
			if err != nil {
				logrus.Errorf("%s: %s", tt.Name, err.Error())
			}
		})
	}
}