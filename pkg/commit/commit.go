package commit

import (
	"github.com/pkg/errors"
	"os/exec"
)

type Committer interface {
	Commit() error
}

type TarCommitter struct {
	ImageTarName  string // 打包镜像生成的文件
	PivotRootPath string // 容器所见目录在宿主机的位置，要打包的路径
}

func NewTarCommitter(tarName string, rootPath string) Committer {
	return &TarCommitter{
		ImageTarName:  tarName,
		PivotRootPath: rootPath,
	}
}

func (t *TarCommitter) Commit() error {
	cmd := exec.Command("tar","-czf",t.ImageTarName,"-C",t.PivotRootPath,".")
	_,err := cmd.CombinedOutput()
	if err!=nil{
		return errors.Wrap(err,"tar root path error")
	}

	return nil
}
