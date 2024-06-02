package cgroup

import (
	"errors"
	"github.com/sirupsen/logrus"
	"mydocker/tool"
	"os"
	"path/filepath"
)

var SupTable = map[string]struct{}{
	"memory.limit_in_bytes": {},
	"cpu.cfs_quota_us":      {},
	"cpu.cfs_period_us":     {},
	"cpuset.cpus":           {},
	"tasks":                 {},
}

func GetRootPath(subSysName string) string {
	return filepath.Join("/sys/fs/cgroup", subSysName)
}

type CgroupOpCfg struct {
	CgroupName string
	SubSysName string
	SpecName   string
	Value      string
	AutoCreate bool
}

func SetSpec(cfg CgroupOpCfg) error {
	if _, ok := SupTable[cfg.SpecName]; !ok {
		return errors.New("not supported subsystem")
	}

	root := GetRootPath(cfg.SubSysName)
	path := filepath.Join(root, cfg.CgroupName, cfg.SpecName)

	if cfg.AutoCreate {
		err := tool.EnsureDirExists(path)
		if err != nil {
			return err
		}
	}

	logrus.Infof("path: %s", path)
	err := set(path, cfg.Value)
	if err != nil {
		return err
	}

	return nil
}

func Delete(cfg CgroupOpCfg) error {
	root := GetRootPath(cfg.SubSysName)
	path := filepath.Join(root, cfg.CgroupName)
	logrus.Infof("delete path: %s",path)
	return os.RemoveAll(path)
}

func set(path string, value string) error {
	if err := os.WriteFile(path, []byte(value), 0644); err != nil {
		return err
	}

	return nil
}
