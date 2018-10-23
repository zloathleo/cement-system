package controller

import (
	"bytes"
	"github.com/zloathleo/go-httpserver/common/logger"
	. "github.com/zloathleo/go-httpserver/dstruct"
	"github.com/zloathleo/go-httpserver/history2"
	"github.com/zloathleo/go-httpserver/utils"
)

const (
	mixHistChart = 600 //10分钟
)

/**
格式化当天历史请求
只能查看from当天
dur 为枚举值[600...86400]
间隔interval 为枚举值[0,2...60]
600 	10分钟   interval=0     600点+1  即获取所有
1800	30分钟   interval=2     900点+1
3600 	60分钟   interval=5     720点+1
14400	4小时	 interval=15    960点+1
28800 	8小时	 interval=30    960点+1
86400 	24小时   interval=60    1440点+1
604800 	168小时  interval=600    1008点+1
*/
func fetchHistoryChartData(pointsArray []string, to int64, dur int) *bytes.Buffer {
	if pointsArray == nil || len(pointsArray) == 0 {
		return nil
	}
	hisReqParam := history2.FormatChartHisReqParam(to, dur)
	if hisReqParam == nil {
		logger.Warnf("request his param is err to[%d], dur[%d]", to, dur)
		return nil
	}

	logger.Infof("full req time %v || %v:", hisReqParam.From, hisReqParam.To)

	fromHour := utils.GetTimeIntHour(hisReqParam.From).Hour()
	toHour := utils.GetTimeIntHour(hisReqParam.To).Hour()

	blockForm := hisReqParam.From
	blockTo := hisReqParam.To
	hisCount := hisReqParam.Count
	interval := hisReqParam.Interval

	//先初始化数据结果
	//时间戳列表
	timestamptArray := make([]int64, hisCount, hisCount)
	for i := 0; i < hisCount; i++ {
		timestamptArray[i] = blockForm.Unix() + int64(i*interval)
	}

	//历史数据
	var hisMap = make(map[string][]*HDas)
	for _, point := range pointsArray {
		hisMap[point] = make([]*HDas, hisCount, hisCount)
	}

	hisReqContext := &HisReqContext{Count: hisCount, Index: 0}

	i := fromHour
	//从第一个小时开始
	for ;i < toHour; i++ {
		blockTo = utils.GetTimeEndOfHour(blockForm)

		//当前table 需要查找的数量
		tableDataCount := int(utils.GetTimeNextIntHour(blockForm).Unix()-blockForm.Unix()) / interval
		history2.SelectSingleTableHistoryData(pointsArray, blockForm, i, blockForm.Unix(), blockTo.Unix(), interval, tableDataCount, hisMap, hisReqContext)

		hisReqContext.Index = hisReqContext.Index + tableDataCount
		logger.Infof("req time tableDataCount %v || %v || %d", blockForm, blockTo, tableDataCount)
		blockForm = utils.GetTimeNextIntHour(blockTo)
	}

	tableDataCount := int(hisReqParam.To.Unix()-blockForm.Unix())/interval + 1
	logger.Infof("req time tableDataCount %v || %v || %d", blockForm, hisReqParam.To, tableDataCount)
	history2.SelectSingleTableHistoryData(pointsArray, blockForm, toHour, blockForm.Unix(), hisReqParam.To.Unix(), hisReqParam.Interval, tableDataCount, hisMap, hisReqContext)
	hisReqContext.Index = hisReqContext.Index + tableDataCount

	return history2.RenderChartHistoryJson(hisMap, timestamptArray)

}
