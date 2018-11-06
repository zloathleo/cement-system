package history2

import (
	"github.com/zloathleo/go-httpserver/common/logger"
	. "github.com/zloathleo/go-httpserver/dstruct"
	"github.com/zloathleo/go-httpserver/utils"
	"time"
)

/**
格式化当天历史请求
只能查看from当天
dur 为枚举值[600...86400]
间隔interval 为枚举值[0,2...60]
180 	3分钟   interval=5     36点+1  即获取所有
600 	10分钟   interval=0     600点+1  即获取所有
1800	30分钟   interval=60     30点+1  radar使用
3600 	60分钟   interval=5     720点+1
14400	4小时	 interval=15    960点+1
28800 	8小时	 interval=30    960点+1
86400 	24小时   interval=60    1440点+1
604800 	168小时  interval=600    1008点+1
*/
func FormatChartHisReqParam(to int64, dur int) *HisReqParam {
	//起始时间
	var from int64
	//时间戳间隔
	var interval int

	//时长
	switch dur {
	case 180:
		{
			//不改
			from = to - 180
			interval = 1
			break
		}
	case 600:
		{
			//不改
			from = to - 600
			interval = 1
			break
		}
	case 1800:
		{
			//30分钟
			from = to - 1800
			interval = 60
			break
		}
	case 3600:
		{
			//60分钟
			from = to - 3600
			interval = 5

			//对齐时间
			from = utils.GetTimeInt5Sec(time.Unix(from, 0)).Unix()
			to = from + 3600
			break
		}
	case 14400:
		{
			//4小时
			from = to - 14400
			interval = 15

			//对齐时间
			from = utils.GetTimeInt15Sec(time.Unix(from, 0)).Unix()
			to = from + 14400

			break
		}
	case 28800:
		{
			//8小时
			from = to - 28800
			interval = 30

			//对齐时间
			from = utils.GetTimeInt30Sec(time.Unix(from, 0)).Unix()
			to = from + 28800
			break
		}
	case 86400:
		{
			//24小时
			from = to - 86400
			interval = 60

			//对齐时间
			from = utils.GetTimeIntMin(time.Unix(from, 0)).Unix()
			to = from + 86400
			break
		}
	case 604800:
		{
			//24小时*7
			from = to - 604800
			interval = 600

			//对齐时间
			from = utils.GetTimeInt10Min(time.Unix(from, 0)).Unix()
			to = from + 604800
			break
		}
	default:
		{
			return nil
		}
	}

	fromTimeFm := time.Unix(from, 0)
	toTime := time.Unix(to, 0)

	logger.Debug("原始 请求 时间: ", utils.GetTimeString(fromTimeFm), utils.GetTimeString(toTime))
	var times []*HisReqParamTime
	//需要分文件查询
	fromStepTime := from //int64
	toStepTime := utils.GetTimeDayEnd(fromTimeFm)
	logger.Debug("原始 请求 时间 begin : ", utils.GetIntTimeString(from))

	for !toStepTime.After(toTime) {
		itemCount := int(toStepTime.Unix()-from) / interval
		toStepTimeInt := from + int64(interval*itemCount)

		//第二天 的下一个 时间戳
		nextFromStepTimeInt := toStepTimeInt + int64(interval)

		//add
		hisReqParamTime := &HisReqParamTime{From: fromStepTime, To: toStepTimeInt, Count: itemCount, Interval: interval}
		logger.Debug("请求 日期分割  时间 step : ", utils.GetIntTimeString(hisReqParamTime.From), utils.GetIntTimeString(hisReqParamTime.To))
		times = append(times, hisReqParamTime)

		fromStepTime = nextFromStepTimeInt
		toStepTime = utils.GetTimeDayEnd(time.Unix(fromStepTime, 0))

	}

	hisReqParamTime := &HisReqParamTime{From: fromStepTime, To: to, Count: int(to-fromStepTime)/interval + 1, Interval: interval}
	logger.Debug("请求 日期分割  时间 step last : ", utils.GetIntTimeString(fromStepTime), utils.GetIntTimeString(to))
	times = append(times, hisReqParamTime)

	logger.Debug("原始 请求 时间 end : ", utils.GetIntTimeString(to))
	return &HisReqParam{From: time.Unix(from, 0), To: time.Unix(to, 0), Times: times, Dur: dur, Interval: interval, Count: dur/interval + 1}
}
