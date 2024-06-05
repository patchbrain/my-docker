package mount

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

type Overlay struct {
	LowerTarPath string
	LowerPath    string
	UpperPath    string
	WorkPath     string
	MergePath    string
	Volume       Volume
}

func NewOverlay(rootMnt, tarPath string, vol Volume) Mounter {
	return &Overlay{
		LowerTarPath: tarPath,
		LowerPath:    filepath.Join(rootMnt, "lower"),
		UpperPath:    filepath.Join(rootMnt, "upper"),
		WorkPath:     filepath.Join(rootMnt, "work"),
		MergePath:    filepath.Join(rootMnt, "merge"),
		Volume:       vol,
	}
}

func (t *Overlay) Mount() error {
	// 准备overlay各层的文件夹
	err := t.MkRelaDir()
	if err != nil {
		return errors.Wrap(err, "make relative dirs error")
	}

	// 执行mount命令
	err = t.execMount()
	if err != nil {
		return errors.Wrap(err, "run mount cmd error")
	}

	return nil
}

func (t *Overlay) UnMount() error {
	err := t.execUmount()
	if err != nil {
		return errors.Wrap(err, "run umount cmd error")
	}

	err = t.removeAll()
	if err != nil {
		return errors.Wrap(err, "remove relative dirs error")
	}

	return nil
}

func (t *Overlay) MkRelaDir() error {
	var err error
	err = t.makeLowerTier()
	if err != nil {
		return errors.Wrap(err, "make lower error error")
	}

	// 准备upper环境
	err = os.MkdirAll(t.UpperPath, 0777)
	if err != nil {
		return errors.Wrap(err, "make upper tier error")
	}

	// 准备work环境
	err = os.MkdirAll(t.WorkPath, 0777)
	if err != nil {
		return errors.Wrap(err, "make work tier error")
	}

	// 准备挂载点目录
	err = os.MkdirAll(t.MergePath, 0777)
	if err != nil {
		return errors.Wrap(err, "make merge tier error")
	}

	// 如果有volume的话，创建挂载点的目录
	if t.Volume.Src != "" {
		err = os.MkdirAll(t.Volume.Dst, 0777)
		if err != nil {
			return errors.Wrap(err, "make volume dst dir error")
		}
	}

	return nil
}

// 执行具体的mount命令
func (t *Overlay) execMount() error {
	var err error
	cmd1 := exec.Command("mount", []string{
		"-t",
		"overlay",
		"overlay",
		"-o",
		fmt.Sprintf(`lowerdir=%s,upperdir=%s,workdir=%s`, t.LowerPath, t.UpperPath, t.WorkPath),
		t.MergePath,
	}...)
	if err = cmd1.Run(); err != nil {
		return errors.Wrap(err, "mount merge error")
	}

	if t.Volume.Src != "" {
		cmd2 := exec.Command("mount", []string{
			"-o",
			"bind",
			t.Volume.Src,
			filepath.Join(t.MergePath, t.Volume.Dst),
		}...)
		if err = cmd2.Run(); err != nil {
			return errors.Wrap(err, "mount bind error")
		}
	}

	return nil
}

func (t *Overlay) makeLowerTier() error {
	targetDir := "/mnt/mydocker/lower"
	err := os.MkdirAll(targetDir, 0777)
	if err != nil {
		return errors.Wrap(err, "mkdir lower error")
	}

	cmd := exec.Command("tar", []string{
		"-xvf",
		t.LowerTarPath,
		"-C",
		targetDir,
	}...)
	if err = cmd.Run(); err != nil {
		return errors.Wrap(err, "run tar error")
	}

	return nil
}

// 清除所有相关目录
func (t *Overlay) removeAll() error {
	for _, path := range []string{t.WorkPath, t.UpperPath, t.LowerPath, t.MergePath} {
		log.Infof("remove path: %s...", path)
		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	}

	return nil
}

// 执行umount，接触挂载
func (t *Overlay) execUmount() error {
	cmd1 := exec.Command("umount", t.MergePath)
	if err := cmd1.Run(); err != nil {
		return err
	}

	if t.Volume.Src != "" {
		cmd2 := exec.Command("umount", filepath.Join(t.MergePath, t.Volume.Dst))
		if err := cmd2.Run(); err != nil {
			return err
		}
	}

	return nil
}
