package controller

import (
	"bytes"
	. "github.com/zloathleo/go-httpserver/dstruct"
	"strconv"
)

//输出历史曲线Json数据
func renderRadarChartHistoryJson(pointsArray []string, hisMap map[string][]*HDas, xAxis []int64) *bytes.Buffer {
	var builder bytes.Buffer

	//root
	builder.WriteString("[ ")

	//数据量
	length := len(hisMap[pointsArray[0]])
	//测点数量
	count := len(hisMap)

	for j := 0; j < length; j++ {
		builder.WriteString("[ ")

		for i := 0; i < count; i++ {
			pn := pointsArray[i]
			das := hisMap[pn][j]

			if das != nil {
				builder.WriteString(strconv.FormatFloat(das.V, 'f', 2, 64) + ",")
			} else {
				builder.WriteString("null,")
			}
		}
		builder.Truncate(builder.Len() - 1)

		builder.WriteString(" ],")
	}
	builder.Truncate(builder.Len() - 1)

	//root
	builder.WriteString(" ]")
	return &builder
}
