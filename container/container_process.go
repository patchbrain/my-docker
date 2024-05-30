package container

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"syscall"
)

// NewParentProcess 返回一个父进程，自己调用了自己，但是调用时在最前面塞了一个参数init
// 也就是要执行mydocker init args... 命令
// 打开tty就是把当前进程的输入输出定向到标准输入输出上
func NewParentProcess(tty bool) (*exec.Cmd, *os.File) {
	rp,wp,err := os.Pipe()
	if err!=nil{
		log.Errorf("create pipe failed: %s",err.Error())
		return nil, nil
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

	// 将readPipe传给子进程
	resCmd.ExtraFiles = []*os.File{rp}

	return resCmd,wp
}
