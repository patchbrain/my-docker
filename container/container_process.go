package container

import (
	"os"
	"os/exec"
	"syscall"
)

// NewParentProcess 返回一个父进程，自己调用了自己，但是调用时在最前面塞了一个参数init
// 也就是要执行mydocker init args... 命令
// 打开tty就是把当前进程的输入输出定向到标准输入输出上
func NewParentProcess(tty bool, cmd string) *exec.Cmd {
	args := []string{"init", cmd}
	resCmd := exec.Command("/proc/self/exe", args...)
	resCmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: uintptr(syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNS),
	}

	if tty {
		resCmd.Stdout = os.Stdout
		resCmd.Stdin = os.Stdin
		resCmd.Stderr = os.Stderr
	}

	return resCmd
}
