package subsystem

import (
	"errors"
	"mydocker/pkg/cgroup"
	"mydocker/resource/config"
	"strconv"
)

type MemSys struct {
	Cfg        *config.Config
	CgroupName string
}

func NewMemSys(cfg *config.Config) *MemSys {
	var res MemSys
	res.Cfg = cfg
	res.CgroupName = cfg.CgroupName

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
		AutoCreate: true,
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
		SpecName:   "tasks",
		Value:      writen,
		AutoCreate: true,
	})
}

func (t *MemSys) Destroy() error {
	return cgroup.Delete(cgroup.CgroupOpCfg{
		CgroupName: t.Cfg.CgroupName,
		SubSysName: t.Name(),
	})
}

func (t *MemSys) Name() string {
	return "memory"
}
