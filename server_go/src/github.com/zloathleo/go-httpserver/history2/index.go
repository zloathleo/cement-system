package history2

import (
	"database/sql"
	"fmt"
	"github.com/zloathleo/go-httpserver/common/logger"
	"github.com/zloathleo/go-httpserver/utils"
	"os"
	"sync"
	"time"
)

var (
	//当天的db 文件
	todayDb *sql.DB

	todayDbName string
	//当天db的读写锁
	todayDbLock = new(sync.RWMutex)
)

func Init() {
	initStorage()
	initCorn()
}

func Exit() {
	logger.Warnln("history storage is exiting.")
	todayDbLock.Lock() // 写加锁
	todayDb.Close()
	todayDbLock.Unlock() // 写解锁

	for fileName, connect := range allDbMap {
		if connect != nil && fileName != todayDbName {
			connect.dbLock.Lock()
			connect.db.Close()
			connect.dbLock.Unlock()
		}
	}
	logger.Warnln("history storage is exited.")
}

func initStorage() {
	dir := "data"
	begin := time.Now().UnixNano()
	exist := utils.IsFileExist(dir)
	if exist {
		logger.Warnln("data file directory is exist.")
	} else {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			logger.Fatalf("db file directory can't create. %v", err)
		} else {
			logger.Warnln("db file directory created.")
		}
	}
	connect := createAnyConnect(time.Now())
	todayDb = connect.db
	logger.Infof("init db const time %d ms.", (time.Now().UnixNano()-begin)/1000000)
}

//查找合适的历史表格
func findHisTable(timeStamp time.Time) string {
	//1539680150
	return fmt.Sprintf("his24_%d", timeStamp.Hour())
}


