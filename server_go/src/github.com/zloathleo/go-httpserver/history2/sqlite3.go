package history2

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zloathleo/go-httpserver/common/logger"
	"github.com/zloathleo/go-httpserver/systemalarm"
	"github.com/zloathleo/go-httpserver/utils"
	"os"
	"sync"
	"time"
)

const (
	driverName = "sqlite3"
	his24table = `CREATE TABLE his24_%d (
    pn    VARCHAR (32) NOT NULL,
    t     INT (32)     NOT NULL,
    value DOUBLE       NOT NULL,
    PRIMARY KEY (
        pn,
        t
    )
	);`
)

var (
	//fileName -- > db
	allDbMap = make(map[string]*SQLite3FileConnect, 32)
)

//根据时间获得当天的DB
func createAnyConnect(t time.Time) *SQLite3FileConnect {
	fileName := utils.GetDayString(t)
	dbConnect := allDbMap[fileName]
	if dbConnect != nil {
		return dbConnect
	} else {
		connect := newSQLite3FileConnect(fileName, true)
		allDbMap[fileName] = connect
		return connect
	}
}

func createTomorrowConnect() *SQLite3FileConnect {
	tomorrow := utils.GetTimeNextDay0Hour(time.Now())
	return createAnyConnect(tomorrow)
}

//获得文件已经存在的连接,不创建新文件
func getAnyConnect(t time.Time) *SQLite3FileConnect {
	fileName := utils.GetDayString(t)
	dbConnect := allDbMap[fileName]
	if dbConnect != nil {
		return dbConnect
	} else {
		connect := newSQLite3FileConnect(fileName, false)
		allDbMap[fileName] = connect
		return connect
	}
}

func getTodayConnect() *SQLite3FileConnect {
	return getAnyConnect(time.Now())
}

func getDBPath(fileName string) string {
	return "data/data-" + fileName + ".db"
}

type SQLite3FileConnect struct {
	name           string //fileName not path
	dbFilePathName string
	db             *sql.DB
	dbLock         *sync.RWMutex
}

func newSQLite3FileConnect(fileName string, create bool) *SQLite3FileConnect {
	dbFilePathName := getDBPath(fileName)

	logger.Infof("newSQLite3FileConnect %s ", dbFilePathName)
	connect := &SQLite3FileConnect{name: fileName, dbFilePathName: dbFilePathName}
	if utils.IsFileExist(dbFilePathName) {
		logger.Infof("newSQLite3FileConnect %s is exist.", dbFilePathName)
		var err error
		connect.db, err = connect.connectGivenFileDb()
		if err != nil || connect.db == nil {
			logger.Errorf("connect given file db %v err.", dbFilePathName)
		} else {
			connect.dbLock = new(sync.RWMutex)
			return connect
		}
	} else if create {
		_, err := os.Create(dbFilePathName)
		if err == nil {
			logger.Infof("newSQLite3FileConnect create %s file success.", dbFilePathName)
			connect.db, err = connect.initGivenFileDb()
			if err != nil || connect.db == nil {
				logger.Errorf("init given file db %v err.", dbFilePathName)
			} else {
				connect.dbLock = new(sync.RWMutex)
				return connect
			}
		} else {
			systemalarm.AddSystemAlarm(err)
			logger.Errorf("newSQLite3FileConnect create %s file is err %v.", dbFilePathName, err)
		}
	}
	return nil
}

/**
初始化指定db文件
*/
func (conn *SQLite3FileConnect) initGivenFileDb() (*sql.DB, error) {
	cdb, err := conn.connectGivenFileDb()

	if cdb != nil && err == nil {
		tx, _ := cdb.Begin()
		for i := 0; i < 24; i++ {
			createTableSql := fmt.Sprintf(his24table, i)
			_, createTableErr := cdb.Exec(createTableSql)
			if createTableErr != nil {
				systemalarm.AddSystemAlarm(createTableErr)
			}
		}
		tx.Commit()
	}
	return cdb, err
}

/**
连接指定db文件
*/
func (conn *SQLite3FileConnect) connectGivenFileDb() (*sql.DB, error) {
	cdb, err := sql.Open(driverName, conn.dbFilePathName)
	if err != nil {
		return nil, err
	}
	if err = cdb.Ping(); err != nil {
		return nil, err
	}
	return cdb, nil
}
