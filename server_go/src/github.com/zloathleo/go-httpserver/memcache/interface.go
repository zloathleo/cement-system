package memcache

import (
	. "github.com/zloathleo/go-httpserver/dstruct"
	"github.com/zloathleo/go-httpserver/utils"
	"math"
)

var (
	//点列表
	PointCacheMap = make(map[string]float64)
)

/**
刷新点列表
*/
func RefreshPointList() {
	for key := range PointCacheMap {
		exist, v := GetCurrentData(key)
		//logger.Println("RefreshPointList key %s,%v:", key, v)
		if exist {
			PointCacheMap[key] = v
		} else {
			delete(PointCacheMap, key)
		}
	}
}

func SaveRealtimeData(pushDas *PushDas) {
	rows := pushDas.Rows
	for _, item := range rows {
		if !math.IsNaN(item.Value){
			saveValue(item)
			PointCacheMap[item.PointName] = item.Value
		}
	}
}

//保存值到缓存
func saveValue(item *Das) {
	//60s 超时
	cache.Set([]byte(item.PointName), utils.Float64ToBytes(item.Value), 60)
}

//获取当前值
func GetCurrentData(pn string) (bool, float64) {
	valueBytes, _ := cache.Get([]byte(pn))
	if valueBytes != nil {
		return true, utils.Float64FromBytes(valueBytes)
	} else {
		return false, 0
	}
}

//获取当前值
func GetCurrentDataInt(pn string) (bool, int) {
	valueBytes, _ := cache.Get([]byte(pn))
	if valueBytes != nil {
		return true, int(utils.UInt32FromBytes(valueBytes))
	} else {
		return false, 0
	}
}

func SetCurrentDataInt(pn string, v int) {
	//60s 超时
	cache.Set([]byte(pn), utils.UInt32ToBytes(uint32(v)), 60)
}
