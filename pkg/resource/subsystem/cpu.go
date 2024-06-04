package subsystem

import (
	"errors"
	"mydocker/pkg/cgroup"
	"mydocker/pkg/resource/config"
	"strconv"
	"strings"
)

type CpuSys struct {
	Cfg        *config.Config
	CgroupName string
}

func NewCpuSys(cfg *config.Config) *CpuSys {
	var res CpuSys
	res.Cfg = cfg
	res.CgroupName = cfg.CgroupName

	return &res
}

func (t *CpuSys) Apply() error {
	if t.CgroupName == "" {
		return errors.New("empty group name")
	}

	if t.Cfg.Cpu == nil{
		return nil
	}
	
	quota := "cpu.cfs_quota_us"
	period := "cpu.cfs_period_us"
	writtenP := "1000000"
	writtenQ := strconv.FormatInt(*t.Cfg.Cpu*1000000/100, 10)
	err := cgroup.SetSpec(cgroup.CgroupOpCfg{
		CgroupName: t.CgroupName,
		SubSysName: t.Name(),
		SpecName:   period,
		Value:      writtenP,
		AutoCreate: true,
	})

	err = cgroup.SetSpec(cgroup.CgroupOpCfg{
		CgroupName: t.CgroupName,
		SubSysName: t.Name(),
		SpecName:   quota,
		Value:      writtenQ,
		AutoCreate: true,
	})

	if err != nil {
		return err
	}

	return nil
}

func (t *CpuSys) AddPid(pid int) error {
	if t.CgroupName == "" {
		return errors.New("empty group name")
	}

	written := strconv.Itoa(pid)
	return cgroup.SetSpec(cgroup.CgroupOpCfg{
		CgroupName: t.CgroupName,
		SubSysName: t.Name(),
		SpecName:   "tasks",
		Value:      written,
		AutoCreate: true,
	})
}

func (t *CpuSys) Destroy() error {
	return cgroup.Delete(cgroup.CgroupOpCfg{
		CgroupName: t.Cfg.CgroupName,
		SubSysName: t.Name(),
	})
}

func (t *CpuSys) Name() string {
	return "cpu"
}

type CpuSetSys struct {
	Cfg        *config.Config
	CgroupName string
}

func NewCpuSetSys(cfg *config.Config) *CpuSetSys {
	var res CpuSetSys
	res.Cfg = cfg
	res.CgroupName = cfg.CgroupName

	return &res
}

func (t *CpuSetSys) Apply() error {
	if t.CgroupName == "" {
		return errors.New("empty group name")
	}

	if t.Cfg.CpuSet == nil{
		return nil
	}
	
	writen := strings.Join(t.Cfg.CpuSet, ",")

	return cgroup.SetSpec(cgroup.CgroupOpCfg{
		CgroupName: t.CgroupName,
		SubSysName: t.Name(),
		SpecName:   "cpuset.cpus",
		Value:      writen,
		AutoCreate: true,
	})
}

func (t *CpuSetSys) AddPid(pid int) error {
	if t.CgroupName == "" {
		return errors.New("empty group name")
	}

	writen := strconv.Itoa(pid)
	return cgroup.SetSpec(cgroup.CgroupOpCfg{
		CgroupName: t.CgroupName,
		SubSysName: t.Name(),
		SpecName:   "tasks",
		Value:      writen,
		AutoCreate: true,
	})
}

func (t *CpuSetSys) Destroy() error {
	return cgroup.Delete(cgroup.CgroupOpCfg{
		CgroupName: t.Cfg.CgroupName,
		SubSysName: t.Name(),
	})
}

func (t *CpuSetSys) Name() string {
	return "cpuset"
}
