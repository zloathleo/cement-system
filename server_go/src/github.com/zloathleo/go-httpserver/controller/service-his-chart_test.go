package controller

import (
	"github.com/zloathleo/go-httpserver/utils"
	"github.com/zloathleo/go-httpserver/common/logger"
	"testing"
	"time"
)

func TestFetchHistoryChartData(t *testing.T) {
	to := utils.GetTimeParseString("2018-10-20 13:10:07")
	//to ,_ := time.ParseInLocation("2006-01-02 15:04:05","2018-10-19 22:19:44",time.Local)

	//pns := []string{"1_00002", "4_40001","4_40002"}

	begin := time.Now().UnixNano()
	pns := []string{"4_40002","4_40003"}
	builder:= fetchHistoryChartData(pns, to.Unix(), 1800)
	logger.Println(builder.String())
	logger.Infof(" const time %d ms.", (time.Now().UnixNano()-begin)/1000000)

	//begin := time.Now().UnixNano()
	//toJosn(hisMap)
	//logger.Infof(" const time %d ms.", (time.Now().UnixNano()-begin)/1000000)

	//begin = time.Now().UnixNano()
	//toJosn2(hisMap)
	//logger.Infof(" const time %d ms.", (time.Now().UnixNano()-begin)/1000000)
}

//func toJosn2(hisMap map[string][]*HDas) {
//
//	var builder bytes.Buffer
//
//	builder.WriteString("{ ")
//	for pn, dasArray := range hisMap {
//		builder.WriteString("\"" + pn + "\":[ ")
//
//		for _, das := range dasArray {
//			if das != nil {
//				builder.WriteString(strconv.FormatFloat(das.V, 'f', 4, 64) + ",")
//			} else {
//				builder.WriteString("null,")
//			}
//		}
//		builder.Truncate(builder.Len()-1)
//
//		builder.WriteString(" ],")
//	}
//	builder.Truncate(builder.Len()-1)
//	builder.WriteString(" }")
//	logger.Println(builder.String())
//}

//func toJosn(hisMap map[string][]*HDas) {
//	var builder strings.Builder
//
//	builder.WriteString("{")
//	for pn, dasArray := range hisMap {
//		builder.WriteString("\"" + pn + "\":[")
//
//		for _, das := range dasArray {
//			if das != nil {
//				builder.WriteString(strconv.FormatFloat(das.V, 'f', 4, 64) + ",")
//			} else {
//				builder.WriteString("null,")
//			}
//
//		}
//
//		builder.WriteString("],")
//	}
//	builder.WriteString("}")
//	logger.Println(builder.String())
//}
