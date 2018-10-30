package controller

import (
	"bytes"
	"github.com/zloathleo/go-httpserver/memcache"
	"strconv"
)

func serviceGetDas(pointsArray []string) *bytes.Buffer {

	var builder bytes.Buffer

	//root
	builder.WriteString("{ ")
	for _, pn := range pointsArray {
		builder.WriteString("\"" + pn + "\":")
		exits, value := memcache.GetCurrentData(pn)
		if exits {
			builder.WriteString(strconv.FormatFloat(value, 'f', 2, 64) + ",")
		} else {
			builder.WriteString("null,")
		}
	}
	if builder.Len() > 4 {
		builder.Truncate(builder.Len() - 1)
	}

	//root
	builder.WriteString(" }")

	return &builder

}
