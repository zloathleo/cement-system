package tiff

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

const (
	rootpath = "D:/tiff-sn/"
	//rootpath = "D:/tiff-hx/20180801/"
)

var fileFound = errors.New("filefound")

func WalkTiffOne(index int, pixel int, _calc string) ([]int, int, int) {

	walkIndex := 0
	//选择的文件数据
	var selectedFileData []int
	var w, h int
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".tif" {
			if index == walkIndex {
				log.Println(path)
				selectedFileData, w, h = ReadTile(path, pixel, _calc)
				return fileFound
			}
			walkIndex++
		}
		return nil
	})
	if err == fileFound {
		return selectedFileData, w, h
	} else {
		return selectedFileData, w, h
	}
}

//统计一段周期高温区域概率
func WalkTiffPeriod(period int, pixel int, _calc string) ([]int, int, int) {
	w := 0
	h := 0
	if pixel == 2 {
		w = 256
		h = 192
	} else if pixel == 4 {
		w = 128
		h = 96
	} else {
		pixel = 8
		w = 64
		h = 48
	}

	//热区tile
	var hots = make([]int, w*h)

	filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".tif" {
			tileValues, w, h := ReadTile(path, pixel, _calc)

			for j := 0; j < h; j++ {
				for i := 0; i < w; i++ {
					tileIndex := j*w + i
					percent := tileValues[tileIndex]
					if percent > HotTemperature {
						//热区+1
						cv := hots[tileIndex]
						hots[tileIndex] = cv + 1

					}
				}
			}

			log.Println(hots)
		}
		return nil
	})
	return hots, w, h

}
