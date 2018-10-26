package utils

import "os"

func IsFileExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
