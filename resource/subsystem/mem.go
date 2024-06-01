package subsystem

import (
	"errors"
	"mydocker/pkg/cgroup"
	"mydocker/resource/config"
	"strconv"
)

type MemSys struct {
	Cfg *config.Config
	CgroupName string
}

func NewMemSys(cfg *config.Config) *MemSys {
	var res MemSys
	res.Cfg = cfg

	return &res
}

func (t *MemSys) Apply() error {
	if t.CgroupName == "" {
		return errors.New("empty group name")
	}

	writen := strconv.FormatInt(t.Cfg.Memory, 10)
	return cgroup.SetSpec(cgroup.CgroupOpCfg{
		CgroupName: t.CgroupName,
		SubSysName: t.Name(),
		SpecName:   "memory.limit_in_bytes",
		Value:      writen,
	})
}

func (t *MemSys) AddPid(pid int64) error {
	if t.CgroupName == "" {
		return errors.New("empty group name")
	}

	writen := strconv.FormatInt(pid, 10)
	return cgroup.SetSpec(cgroup.CgroupOpCfg{
		CgroupName: t.CgroupName,
		SubSysName: t.Name(),
		SpecName:   "task",
		Value:      writen,
	})
}

func (t *MemSys) Remove(pid int) error {
	//TODO implement me
	panic("implement me")
}

func (t *MemSys) Name() string {
	return "memory"
}
