package tiff

import (
	"bufio"
	"golang.org/x/image/tiff"
	"log"
	"math"
	"os"
)

const (
	RawW    = 512
	RawH    = 384
	RawSize = 196608
)

/**
tilePixel:tile分辨率像素 x,y
*/
func fetchTileData(filePath string, tilePixel int) ([][]int, int, int) {
	var rData = readRaw(filePath)

	//tilePixel 高分辨率4 低分辨率10
	var tileWCount = 0     //tile数量W
	var tileHCount = 0     //tile数量H
	var tileCount = 0      //tile数量
	var tilePixelCount = 0 //tile内像素点数量
	if tilePixel == 2 {
		tileWCount = 256
		tileHCount = 192
	} else if tilePixel == 4 {
		tileWCount = 128
		tileHCount = 96
	} else {
		tilePixel = 8
		tileWCount = 64
		tileHCount = 48
	}
	tileCount = tileWCount * tileHCount
	tilePixelCount = tilePixel * tilePixel
	//实际处理像素
	WW := tilePixel * tileWCount
	HH := tilePixel * tileHCount

	log.Println("tilePixelCount:", tilePixelCount)

	//var tileData [tileCount][tilePixelCount]int
	var tileData = make([][]int, tileCount)
	//初始化切片
	for x := 0; x < tileCount; x++ {
		tileData[x] = make([]int, tilePixelCount)
	}

	for j := 0; j < HH; j++ {
		//ix  tile的坐标
		iy := int(math.Floor(float64(j) / float64(tilePixel)))

		for i := 0; i < WW; i++ {
			// iy tile的坐标
			ix := int(math.Floor(float64(i) / float64(tilePixel)))

			//tx ty tile内的坐标
			tx := i % tilePixel
			ty := j % tilePixel
			//log.Println("tile内的坐标:", tx, ty)

			//rData[i * H + j]
			rIndex := j*WW + i
			v := rData[rIndex]
			//log.Println("原始坐标:", rIndex)

			//tile序号
			iIndex := iy*tileWCount + ix
			//tileArray := tileData[iIndex]
			//log.Println("tile序号:", iIndex)

			//排除2这种异常值
			if v > 100 {
				//tile内序号
				tIndex := ty*tilePixel + tx
				//tileDatas := tileData[iIndex]
				//log.Println("值信息:", tx, ty, tileDatas, v)
				tileData[iIndex][tIndex] = v
				//log.Println("值信息:", tx, ty, tIndex, v)
				//log.Println("-------------------------------------------")
			}
		}
	}
	//log.Println(tileData)
	return tileData, tileWCount, tileHCount
}

//解析原始数据
func readRaw(filePath string) [RawSize]int {
	file, _ := os.Open(filePath)
	fileReader := bufio.NewReader(file)
	image, _ := tiff.Decode(fileReader)
	var tempData [RawSize]int //512  384
	for j := 0; j < RawH; j++ {
		for i := 0; i < RawW; i++ {
			color := image.At(i, j)
			r, _, _, _ := color.RGBA()
			// g b a 无左右?
			if r > 100 {
				tempData[j*RawW+i] = int(r)
			}
		}
	}
	return tempData
}
