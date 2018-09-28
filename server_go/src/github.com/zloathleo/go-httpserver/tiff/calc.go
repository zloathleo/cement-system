package tiff

const (
	HotTemperature = 1400
	//128 * 96  4 4
)

//计算重心
func calcGravityCenter(tileValues []int, w int, h int) []int {
	//重心坐标,值
	xg := 0
	yg := 0
	valueg := 0
	numg := 0
	for j := 0; j < h; j++ {
		//ix  tile的坐标
		for i := 0; i < w; i++ {
			percent := tileValues[j*w+i]
			if percent > HotTemperature {
				xg += i
				yg += j
				valueg += percent
				numg++
			}
		}
	}
	if numg == 0 {
		return nil
	} else {
		//重心计算
		cx := xg / numg
		cy := yg / numg
		cv := valueg / numg
		return []int{cx, cy, cv}
	}
}

func ReadTile(filePath string, pixel int, _calc string) ([]int, int, int) {
	var tileData, tw, th = fetchTileData(filePath, pixel)
	length := tw * th

	var tileValue = make([]int, length)
	for j := 0; j < th; j++ {
		//ix  tile的坐标
		for i := 0; i < tw; i++ {
			index := j*tw + i
			tdata := tileData[index]
			//选择如何平均
			if "w" == _calc {
				tileValue[index] = getWeightPercentValue(tdata)
			} else if "m" == _calc {
				tileValue[index] = getMaxValue(tdata)
			} else {
				tileValue[index] = getPercentValue(tdata)
			}
		}
	}
	return tileValue, tw, th
}

//////////////////////////////////////
/////////计算平均
//////////////////////////////////////
//算数平均
func getPercentValue(tdata []int) int {
	sum := 0
	trueCount := 0
	for _, i := range tdata {
		if i > 100 {
			sum += i
			trueCount++
		}
	}
	if trueCount == 0 {
		return 0
	}
	return sum / trueCount
}

//加权平均
func getWeightPercentValue(tdata []int) int {
	//weightValue := 1400
	sumweight := 0
	for _, i := range tdata {
		sumweight += i * i / HotTemperature
	}
	return sumweight / len(tdata)
}

//取最大值
func getMaxValue(tdata []int) int {
	max := 0
	for _, i := range tdata {
		if max < i {
			max = i
		}
	}
	return max
}
