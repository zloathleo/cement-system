package history2

import (
	"database/sql"
	"fmt"
	"github.com/zloathleo/go-httpserver/common/logger"
	. "github.com/zloathleo/go-httpserver/dstruct"
	"github.com/zloathleo/go-httpserver/systemalarm"
	"strconv"
	"strings"
	"time"
	"github.com/zloathleo/go-httpserver/utils"
)

func getSelectPointsString(points []string) string {
	a := make([]string, len(points))
	for i := 0; i < len(points); i++ {
		a[i] = "'" + points[i] + "'"
	}
	return strings.Join(a, ",")
}

//查找合适的历史表格
func findHisTableByIndex(index int) string {
	return fmt.Sprintf("his24_%d", index)
}

/**
查找指定文件 指定表格的历史
*/
func SelectSingleTableHistoryData(points []string, begin time.Time, tableIndex int, from int64, to int64, interval int, tableDataCount int, hisMap map[string][]*HDas, hisReqContext *HisReqContext) {
	if points == nil || len(points) == 0 {
		logger.Debug("request point is empty.")
		return
	}
	currentConnect := getAnyConnect(begin)
	if currentConnect == nil {
		logger.Infof("can find connect %s.", utils.GetDayString(begin))
		return
	}
	tableName := findHisTableByIndex(tableIndex)
	pointsString := getSelectPointsString(points)

	selectSql := "select * from " + tableName + " t where t.pn in (" + pointsString + ") and t.t >= " + strconv.FormatInt(from, 10) + " and t.t <= " + strconv.FormatInt(to, 10) + " ORDER BY t.t"
	logger.Debug(selectSql)

	var readLock = todayDbLock
	var readDb = todayDb
	//如果不是查询今天
	if currentConnect.name != todayDbName {
		readLock = currentConnect.dbLock
		readDb = currentConnect.db
	}
	readLock.RLock()
	rows, err := readDb.Query(selectSql, 1)

	defer func() {
		readLock.RUnlock()
		if rows != nil {
			rows.Close()
		}
	}()

	if err == nil && rows != nil {
		foreachRow(points, from, to, interval, tableDataCount, rows, hisMap, hisReqContext)
	} else {
		systemalarm.AddSystemAlarm(err)
		return
	}

}

func foreachRow(points []string, from int64, to int64, interval int, tableDataCount int, rows *sql.Rows, hisMap map[string][]*HDas, hisReqContext *HisReqContext) {
	var fromUInt64 = uint64(from)
	//var appendTimestampIndex = hisReqContext.Index

	var pn string
	var t uint64
	var value float64

	//需要的时间戳 pn -> 开始时间
	//pointBeginMap := make(map[string]uint64)

	//上一次的Index
	pointJumpMap := make(map[string]int)
	//上一次的值
	pointValueMap := make(map[string]*HDas)
	for _, pn := range points {
		//pointBeginMap[pn] = uint64(from)
		pointJumpMap[pn] = 0
		pointValueMap[pn] = nil
	}

	for rows.Next() {
		err := rows.Scan(&pn, &t, &value)
		//logger.Println("读取值::", time.Unix(int64(t), 0).Format("2006-01-02 15:04:05"), value)
		if err != nil {
			systemalarm.AddSystemAlarm(err)
			return
		} else {

			handleRecode(pn, t, value,
				fromUInt64, interval, tableDataCount, hisMap, hisReqContext,
				pointJumpMap, pointValueMap)

		}
	}

	for _pointName, lastDas := range pointValueMap {
		if lastDas != nil {
			//最后一个数值补齐  补齐后续1分钟
			lastEndTime := lastDas.T + 60
			if int64(lastEndTime) > to {
				lastEndTime = uint64(to)
			}
			jump := int(lastEndTime-fromUInt64) / interval
			//newIndex := hisReqContext.Index + jump
			lastJump := pointJumpMap[_pointName]
			lastValue := lastDas.V

			//if math.IsNaN(lastValue) {
			//} else {
			//从上次Index开始补齐值
			for j := lastJump + 1; j <= jump; j++ {
				jt := fromUInt64 + uint64(interval*j)
				jIdx := hisReqContext.Index + j
				hisMap[_pointName][jIdx] = &HDas{V: lastValue, T: jt}
				//logger.Println("最后补齐:", time.Unix(int64(jt), 0), lastValue)
			}
			//}
		}
	}

}

func handleRecode(pn string, t uint64, value float64,
	fromUInt64 uint64, interval int, tableDataCount int, hisMap map[string][]*HDas, hisReqContext *HisReqContext,
	pointJumpMap map[string]int, pointValueMap map[string]*HDas) {

	remainder := int(t-fromUInt64) % interval
	jump := int(t-fromUInt64) / interval
	newIndex := hisReqContext.Index + jump

	lastJump := pointJumpMap[pn]
	lastValue := pointValueMap[pn]
	if lastValue == nil {

	} else {
		//从上次Index开始补齐值
		for j := lastJump + 1; j <= jump; j++ {
			jt := fromUInt64 + uint64(interval*j)
			jIdx := hisReqContext.Index + j
			hisMap[pn][jIdx] = &HDas{V: lastValue.V, T: jt}
			//logger.Println("补齐:", time.Unix(int64(jt), 0), lastValue)
		}
	}

	//余数==0代表就在需要的时间戳伤
	if remainder == 0 {
		hisMap[pn][newIndex] = &HDas{V: value, T: t}

		//logger.Println("正式插入1:", time.Unix(int64(t), 0), value)
	} else {
	}

	pointJumpMap[pn] = jump
	pointValueMap[pn] = &HDas{V: value, T: t}
}
