package systemalarm

import (
	"fmt"
	"github.com/zloathleo/go-httpserver/common/logger"
	"github.com/zloathleo/go-httpserver/dstruct"
	"sync"
	"time"
)

var (
	alarmStack []*dstruct.Alarm
	//当天db的读写锁
	alarmsDbLock = new(sync.RWMutex)
)

func init() {
	alarmStack = make([]*dstruct.Alarm, 0)
}
func AddSystemAlarm(err error) {
	alarmsDbLock.Lock()
	defer alarmsDbLock.Unlock()
	str := fmt.Sprintf("system alarm [%v]", err)
	newAlarm := &dstruct.Alarm{Module: "system", Err: str, TimeStamp: time.Now().Unix()}

	if len(alarmStack) >= 20 {
		alarmStack = alarmStack[1:20]
	}
	alarmStack = append(alarmStack, newAlarm)
	logger.Warnf("system alarm [%v]", err)

}

func GetSystemAlarm() []*dstruct.Alarm {
	alarmsDbLock.RLock()
	defer alarmsDbLock.RUnlock()
	return alarmStack[0:len(alarmStack)]
}
