package subsystem

import (
	"mydocker/resource/config"
)

type CpuSys struct{
	Cfg *config.Config
}

func NewCpuSys(cfg *config.Config) *CpuSys{
	var res CpuSys
	res.Cfg = cfg
	return &res
}

func (t *CpuSys) Apply() error {
	//TODO implement me
	panic("implement me")
}

func (t *CpuSys) AddPid(pid int) error {
	//TODO implement me
	panic("implement me")
}

func (t *CpuSys) Remove(pid int) error {
	//TODO implement me
	panic("implement me")
}

func (t *CpuSys) Name() string {
	return "cpu"
}



type CpuSetSys struct{
	Cfg *config.Config
}

func NewCpuSetSys(cfg *config.Config) *CpuSetSys{
	var res CpuSetSys
	res.Cfg = cfg
	return &res
}

func (t *CpuSetSys) Apply() error {
	//TODO implement me
	panic("implement me")
}

func (t *CpuSetSys) AddPid(pid int) error {
	//TODO implement me
	panic("implement me")
}

func (t *CpuSetSys) Remove(pid int) error {
	//TODO implement me
	panic("implement me")
}

func (t *CpuSetSys) Name() string {
	return "cpu"
}
