/**
每晚生成db file
每小时初始化表格数据
*/
package history2

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"github.com/zloathleo/go-httpserver/common/logger"
	"github.com/zloathleo/go-httpserver/memcache"
	"github.com/zloathleo/go-httpserver/systemalarm"
	"github.com/zloathleo/go-httpserver/utils"
	"time"
)

const (
	cornSpecCreateDBFile     = "0 59 23 * * ?"  //每天23:59整执行一次 -- 创建db文件
	cornSpecChangeCurrentDB  = "0 0 0 * * ?"    //每天0:0:0.500 执行一次 -- 切换db并插入新表缓存数据
	cornSpecInitNewTableInit = "0 0 1-23 * * ?" //每小时0:0:0.500 执行一次 0:0:0.500不执行 -- 插入新表缓存数据
)



func initCorn() {
	c := cron.New()

	//每天23:59整执行一次 -- 创建db文件
	c.AddFunc(cornSpecCreateDBFile, func() {
		createTomorrowConnect()
	})

	//每天0:0:0.500 执行一次 -- 切换db并插入新表缓存数据
	c.AddFunc(cornSpecChangeCurrentDB, func() {
		time.Sleep(500 * time.Millisecond)
		todayDbLock.Lock() // 写加锁
		exchangeTodayDB()
		todayDbLock.Unlock() // 写解锁

	})

	//每小时0:0:0.500 执行一次 0:0:0.500不执行 -- 插入新表缓存数据
	c.AddFunc(cornSpecInitNewTableInit, func() {
		time.Sleep(500 * time.Millisecond)
		todayDbLock.Lock() // 写加锁
		initNewTable()
		todayDbLock.Unlock() // 写解锁

	})

	c.Start()
}

/**
切换DB文件
*/
func exchangeTodayDB() {
	connect := getTodayConnect()

	if connect == nil {
		logger.Errorf("today db file %s is not exist", utils.GetTodayString())
		systemErr := errors.New(fmt.Sprintf("today db file %s is not exist", utils.GetTodayString()))
		systemalarm.AddSystemAlarm(systemErr)
	} else {
		todayDb = connect.db
		logger.Infof("exchange today db %s ok", utils.GetTodayString())
		initNewTable()
	}
}

/**
初始化当前表格,缓存中数据写入当前新表
启动
每小时0:0:0.500 执行
每天凌晨0:0:0.500执行
*/
func initNewTable() {
	logger.Info("begin insert memcache into newtable")
	if len(memcache.PointCacheMap) == 0 {
		return
	}
	//刷新点列表
	memcache.RefreshPointList()
	now := time.Now()
	tableName := findHisTable(now)
	preInsertCache(tableName, now)
	logger.Infof("end init new table %s at time %v ", tableName, now)
}
