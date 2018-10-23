package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func IsFileExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

//获得运行目录
func GetAppPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return filepath.Dir(path), nil
}
