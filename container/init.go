package container

import (
	log "github.com/sirupsen/logrus"
	"os"
	"syscall"
)

func Initer(cmd string, args []string) error {
	// 使得容器内部的挂载独立于主机
	// 由于 systemd 默认将根挂载为共享 (--shared)，挂载事件的隔离并未自动实现。
	// 因此，代码需要显式地将根挂载点标记为 MS_PRIVATE，以确保挂载事件的隔离性，并通过 MS_REC 递归地应用这一属性。
	err := syscall.Mount("","/","",uintptr(syscall.MS_PRIVATE|syscall.MS_REC),"")
	if err!=nil{
		log.Errorf("@ContainerIniter PRIVATE mount root dir failed: %s",err.Error())
		return err
	}

	// 挂载proc
	err = syscall.Mount("proc", "/proc", "proc",
		uintptr(syscall.MS_NODEV|syscall.MS_NOEXEC|syscall.MS_NOSUID), "")
	if err!=nil{
		log.Errorf("@ContainerIniter mount proc failed: %s",err.Error())
		return err
	}

	// 用用户给的命令来替换当前的进程
	args = append(args, cmd)
	log.Infof("args: %#v",args)
	err = syscall.Exec(cmd,args,os.Environ())
	if err!=nil{
		log.Errorf("@ContainerIniter syscall.exec command failed: %s",err.Error())
		return err
	}

	return nil
}
