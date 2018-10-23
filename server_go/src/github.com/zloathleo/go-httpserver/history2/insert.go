package history2

import (
	"github.com/zloathleo/go-httpserver/common/logger"
	. "github.com/zloathleo/go-httpserver/dstruct"
	"github.com/zloathleo/go-httpserver/memcache"
	"github.com/zloathleo/go-httpserver/utils"
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
	todayDbLock.Lock()
	insertRealtimeDas(tableName, now, rows)
	todayDbLock.Unlock()
}

func insertRealtimeDas(tableName string, timeStamp int64, rows []*Das) {
	tx, _ := todayDb.Begin()
	stmt, _ := todayDb.Prepare("INSERT INTO " + tableName + " (pn, t, value) VALUES (?, ?, ?)")

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	//logger.Println("db save timeStamp data:", timeStamp)
	for _, item := range rows {
		//这里使用同一时间戳
		stmt.Exec(item.PointName, timeStamp, item.Value)
	}
	tx.Commit()

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
