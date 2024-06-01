package resource

import (
	log "github.com/sirupsen/logrus"
	"mydocker/resource/subsystem"
	"sync"
)

type ResourceManager struct {
	SubSysSet []subsystem.System
}

var defautRecMgr *ResourceManager
var once sync.Once

func MgrIns() *ResourceManager {
	once.Do(func() {
		defautRecMgr = &ResourceManager{
			SubSysSet: make([]subsystem.System, 0),
		}
	})
	return defautRecMgr
}

func (t *ResourceManager) Apply() {
	if len(t.SubSysSet) == 0{
		log.Errorf("no subsystems, please register subsystem first")
		return
	}

	for _, system := range t.SubSysSet {
		err := system.Apply()
		if err != nil {
			log.Errorf("apply sub system failed: %s", err.Error())
			continue
		}
	}

	return
}

func (t *ResourceManager) Register(sub... subsystem.System) {
	log.Infof("register subsystems: %#v", sub)
	t.SubSysSet = append(t.SubSysSet, sub...)
	return
}
