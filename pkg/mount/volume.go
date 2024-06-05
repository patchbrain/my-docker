package mount

import (
	"errors"
	"os"
	"strings"
)

type Volume struct {
	Src string // 源挂载目录
	Dst string // 目标挂载目录
}

func GetVolume(volStr string) (Volume, error) {
	empty := Volume{}
	vol := strings.Split(volStr, ":")
	if len(vol) != 2 {
		return empty, errors.New("invalid volume param")
	}

	src := vol[0]
	_, err := os.Stat(src)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return empty, errors.New("source dir not exist")
		}

		return empty, err
	}

	return Volume{
		Src: vol[0],
		Dst: vol[1],
	}, nil
}
