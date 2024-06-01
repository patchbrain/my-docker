package cgroup

import (
	"errors"
	"os"
	"path/filepath"
)

var SupTable = map[string]struct{}{
	"memory.limit_in_bytes": {},
	"cpu.cfs_quota_us":      {},
	"cpu.cfs_period_us":     {},
	"cpuset.cpus":           {},
}

func GetRootPath(subSysName string) string {
	return "/sys/fs/cgroup"
}

type CgroupOpCfg struct {
	CgroupName string
	SubSysName string
	SpecName   string
	Value      string
}

func SetSpec(cfg CgroupOpCfg) error {
	if _, ok := SupTable[cfg.SpecName]; !ok {
		return errors.New("not supported subsystem")
	}

	root := GetRootPath(cfg.SubSysName)
	path := filepath.Join(root, cfg.CgroupName, cfg.SpecName)

	set(path, cfg.Value)

	return nil
}

func Delete() {

}

func set(path string, value string) error {
	if err := os.WriteFile(path, []byte(value), 0644); err != nil {
		return err
	}

	return nil
}
