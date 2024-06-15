package container

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func Initer() error {
	err := preMount()
	if err != nil {
		return err
	}
	log.Info("@Initer mount over")

	// 从readPipe获取命令
	args, err := getArgsInPipe()
	if err != nil {
		log.Errorf("@Initer get args from pipe failed: %s", err.Error())
		return err
	}
	log.Infof("@Initer get args from pipe: %#v", args)

	path, err := exec.LookPath(args[0])
	if err != nil {
		log.Errorf("@Initer lookpath of %s failed: %s", args[0], err.Error())
		return err
	}
	log.Infof("@Initer get path: %s", args[0])

	// 用用户给的命令来替换当前的进程
	log.Infof("@Initer args: %#v", args)

	// syscall.Exec会直接替换当前golang程序，该程序之后的代码将不会执行
	err = syscall.Exec(path, args, os.Environ())
	if err != nil {
		log.Errorf("@ContainerIniter syscall.exec command failed: %s", err.Error())
		return err
	}

	return nil
}

func getArgsInPipe() ([]string, error) {
	rp := os.NewFile(uintptr(3), "pipe")
	b, err := io.ReadAll(rp)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(b), " "), nil
}

func preMount() error {
	// 使得容器内部的挂载独立于主机
	// 由于 systemd 默认将根挂载为共享 (--shared)，挂载事件的隔离并未自动实现。
	// 因此，代码需要显式地将根挂载点标记为 MS_PRIVATE，以确保挂载事件的隔离性，并通过 MS_REC 递归地应用这一属性。
	err := syscall.Mount("", "/", "", uintptr(syscall.MS_PRIVATE|syscall.MS_REC), "")
	if err != nil {
		log.Errorf("@ContainerIniter PRIVATE mount root dir failed: %s", err.Error())
		return err
	}

	// 挂载proc
	err = syscall.Mount("proc", "/proc", "proc",
		uintptr(syscall.MS_NODEV|syscall.MS_NOEXEC|syscall.MS_NOSUID), "")
	if err != nil {
		log.Errorf("@ContainerIniter mount proc failed: %s", err.Error())
		return err
	}

	return nil
}
