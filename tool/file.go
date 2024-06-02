package tool

import (
	"os"
	"path/filepath"
)

// EnsureDirExists 检查文件是否存在，如果不存在则创建文件
func EnsureDirExists(path string) error {
	path = filepath.Dir(path)
	_, statErr := os.Stat(path)
	if os.IsNotExist(statErr) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			return err
		}

		return nil
	}
	// 返回可能的其他错误
	return statErr
}
