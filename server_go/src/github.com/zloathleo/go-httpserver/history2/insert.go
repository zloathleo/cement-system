package history2

import (
	"fmt"
	"github.com/zloathleo/go-httpserver/common/logger"
	. "github.com/zloathleo/go-httpserver/dstruct"
	"github.com/zloathleo/go-httpserver/memcache"
	"github.com/zloathleo/go-httpserver/systemalarm"
	"github.com/zloathleo/go-httpserver/utils"
	"strings"
	"time"
)

/**
插入数据到当前文件和当前表格
*/
func SaveRealtimeData2DB(pushDas *PushDas) {
	rows := pushDas.Rows

	//统一使用服务器时间
	timeStamp := time.Now()
	now := timeStamp.Unix()

	tableName := findHisTable(timeStamp)
	begin := time.Now().UnixNano()
	todayDbLock.Lock()
	insertRealtimeDas(tableName, now, rows)
	todayDbLock.Unlock()
	logger.Debugf("SaveRealtimeData2DB %d rows const: %d ms", len(rows), (time.Now().UnixNano()-begin)/int64(1000000))
}

func batchInsert(tableName string, timeStamp int64, unsavedRows []*Das) error {
	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0, len(unsavedRows)*3)
	for _, item := range unsavedRows {
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, item.PointName)
		valueArgs = append(valueArgs, timeStamp)
		valueArgs = append(valueArgs, item.Value)
	}
	stmt := fmt.Sprintf("INSERT INTO "+tableName+" (pn, t, value) VALUES %s", strings.Join(valueStrings, ","))
	_, err := todayDb.Exec(stmt, valueArgs...)
	return err
}

func insertRealtimeDas(tableName string, timeStamp int64, rows []*Das) {
	err := batchInsert(tableName, timeStamp, rows)
	if err != nil {
		systemalarm.AddSystemAlarm(err)
	}
	//tx, _ := todayDb.Begin()
	//stmt, _ := todayDb.Prepare("INSERT INTO " + tableName + " (pn, t, value) VALUES (?, ?, ?)")
	//
	//defer func() {
	//	if stmt != nil {
	//		stmt.Close()
	//	}
	//}()
	////logger.Println("db save timeStamp data:", timeStamp)
	//for _, item := range rows {
	//	//这里使用同一时间戳
	//	if !math.IsNaN(item.Value) {
	//		stmt.Exec(item.PointName, timeStamp, item.Value)
	//	}
	//}
	//tx.Commit()

}

//预存数据进新table
func preInsertCache(tableName string, timeStamp time.Time) {
	rows := memcache.PointCacheMap
	if len(rows) == 0 {
		return
	}
	//整点时间
	intHourTime := utils.GetTimeIntHour(timeStamp).Unix()
	logger.Infof("insert cahce into new table [%s],[%v] at time [%v]", tableName, timeStamp, intHourTime)

	tx, _ := todayDb.Begin()
	stmt, _ := todayDb.Prepare("INSERT INTO " + tableName + " (pn, t, value) VALUES (?, ?, ?)")
	defer stmt.Close()
	//logger.Println("db save timeStamp data:", timeStamp)
	for key, value := range rows {
		//这里使用同一时间戳
		stmt.Exec(key, intHourTime, value)
	}
	stmt.Exec("---key---", intHourTime, 999999)
	tx.Commit()
}
