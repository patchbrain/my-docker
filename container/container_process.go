package container

import (
	log "github.com/sirupsen/logrus"
	"mydocker/config"
	"mydocker/pkg/mount"
	"os"
	"os/exec"
	"syscall"
)

const ROOTFS = "/root/busybox"

// NewParentProcess 返回一个父进程，自己调用了自己，但是调用时在最前面塞了一个参数init
// 也就是要执行mydocker init args... 命令
// 打开tty就是把当前进程的输入输出定向到标准输入输出上
func NewParentProcess(tty bool, volStr string) (*exec.Cmd, *os.File, mount.Mounter) {
	rp, wp, err := os.Pipe()
	if err != nil {
		log.Errorf("create pipe failed: %s", err.Error())
		return nil, nil, nil
	}

	resCmd := exec.Command("/proc/self/exe", "init")
	resCmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: uintptr(syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNS),
	}

	if tty {
		resCmd.Stdout = os.Stdout
		resCmd.Stdin = os.Stdin
		resCmd.Stderr = os.Stderr
	}

	// 准备unionFS
	var vol mount.Volume
	if volStr != "" {
		vol, err = mount.GetVolume(volStr)
		if err != nil {
			log.Infof("volume failed: %s", err.Error())
			os.Exit(-1)
		}
	}

	ol := mount.NewOverlay(config.MntPath, "/root/busybox.tar", vol)
	err = ol.Mount()
	if err != nil {
		log.Infof("gen unionFS env failed: %s", err.Error())
		return nil, nil, nil
	}

	// 将readPipe传给子进程
	resCmd.ExtraFiles = []*os.File{rp}
	resCmd.Dir = ol.(*mount.Overlay).MergePath // 确定rootfs路径

	return resCmd, wp, ol
}
