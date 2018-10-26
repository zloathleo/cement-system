package history

import (
	"bytes"
	"github.com/zloathleo/go-httpserver/utils"
	"math"
	"github.com/zloathleo/go-httpserver/common/logger"
)

/**
* id[uint64(8Byte)] + 实时数据长度[uint16(2Byte)] + 统计数据长度[uint16(2Byte)] + 实时数据[ 时间偏移uint16(2Byte) + float32(4Byte) * n ] + 数据统计[statistics]
* 转换block写一个点历史
 */
func blockToBytes(begin uint64,id uint64, dass[]*Das) []byte{

	//id 8bytes
	buff :=  new(bytes.Buffer)
	buff.Write(utils.UInt64ToBytes(id))

	//实时数据长度 2bytes
	count := len(dass)
	dataLength := 6 * count
	buff.Write(utils.UInt16ToBytes(uint16(dataLength)))

	//统计数据长度 2bytes
	//max min avg
	buff.Write(utils.UInt16ToBytes(uint16(4 + 4 + 4)))

	var max = float32(math.SmallestNonzeroFloat32)
	var min = float32(math.MaxFloat32)
	var total float32 = 0

	//历史块 6 * n
	for _,das := range dass {
		//时间偏移 秒 600 以内
		timeOffset := das.TimeStamp - begin
		buff.Write(utils.UInt16ToBytes(uint16(timeOffset)))

		value := das.Value
		buff.Write(utils.Float32ToBytes(value))

		if value > max{
			max = value
		}
		if value < min{
			min = value
		}
		total += value
	}
	avg := total / float32(count)

	logger.Println("max:",max)
	logger.Println("min:",min)
	logger.Println("avg:",avg)

	buff.Write(utils.Float32ToBytes(max))
	buff.Write(utils.Float32ToBytes(min))
	buff.Write(utils.Float32ToBytes(avg))

	return buff.Bytes()
}

//id[uint64(8Byte)] + 实时数据长度[uint16(2Byte)] + 统计数据长度[uint16(2Byte)] + 实时数据[ 时间偏移uint16(2Byte) + float32(4Byte) * n ] + 数据统计[statistics]
func blockFromBytes(datas []byte){
	id :=utils.UInt64FromBytes(datas[0:8])
	dataLength :=int(utils.UInt16FromBytes(datas[8:10]))

	statisticsLength := utils.UInt16FromBytes(datas[10:12])
	logger.Println(id)
	logger.Println(dataLength)
	logger.Println(statisticsLength)

	count := dataLength / 6
	index := 12
	for i:=0; i<count;i++  {
		timeOffset :=utils.UInt16FromBytes(datas[index:index + 2])
		value :=utils.Float32FromBytes(datas[index + 2:index + 6])

		index = index + 6
		logger.Println("timeOffset:",timeOffset)
		logger.Println("value:",value)
	}
	max := utils.Float32FromBytes(datas[index:index + 4])
	min := utils.Float32FromBytes(datas[index+4:index + 8])
	avg := utils.Float32FromBytes(datas[index+8:index + 12])

	logger.Println("max:",max)
	logger.Println("min:",min)
	logger.Println("avg:",avg)

}
