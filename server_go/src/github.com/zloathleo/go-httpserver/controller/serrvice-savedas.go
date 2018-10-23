package controller

import (
	"github.com/zloathleo/go-httpserver/common"
	"github.com/zloathleo/go-httpserver/common/config"
	. "github.com/zloathleo/go-httpserver/dstruct"
	"github.com/zloathleo/go-httpserver/history2"
	"github.com/zloathleo/go-httpserver/memcache"

	"github.com/gin-gonic/gin/json"
	"github.com/zloathleo/go-httpserver/common/logger"
	"time"
)

const (
	//保护逻辑最大skip次数
	maxAutoBackCheckCount = 20

	//A探头电源关指令缓存点名
	aPowerOffCountPn = "Logic_APowerOffCount"
	//A探头进退指令缓存点名
	aReturnBackCountPn = "aReturnBackCount"
)

var (
	//自动保护autoBackCheckCount > maxAutoBackCheckCount 次后 才生效
	autoBackCheckCount = 0
)

func serviceSaveDas(pushDas *PushDas) {

	rows := pushDas.Rows
	for _, item := range rows {
		item.Value = common.TranslatePointValue(item.PointName, item.Value)
	}

	//历史数据库
	history2.SaveRealtimeData2DB(pushDas)
	//保存到缓存
	memcache.SaveRealtimeData(pushDas)

	//自动退出
	if config.AppConfig.App.AutoLogic {
		//检查环境退出探头或关闭电源
		go autoBackCheckStatus()
	}
}

//自动退出 检查环境退出探头或关闭电源
func autoBackCheckStatus() {
	if autoBackCheckCount < maxAutoBackCheckCount {
		autoBackCheckCount++
		return
	} else {
		if autoBackCheckCount < maxAutoBackCheckCount+2 {
			autoBackCheckCount++
		}
	}

	//A探头退到位
	_, at := memcache.GetCurrentData("1_00005")
	////B探头退到位
	//_, bt := getCurrentData("1_00007")

	//吹扫风合法值>=5
	e1 := validPoint("4_40001")
	//冷却风合法
	e2 := validPoint("4_40002")

	//镜头A温度<=55
	e3 := validPoint("4_40003")
	//探头A内部温度<=55
	e4 := validPoint("5_40013")

	////镜头B温度<=55
	//e5 := validPoint("4_40004")
	////探头B内部温度<=55
	//e6 := validPoint("5_40083")

	var rows []*Das

	logger.Debug("auto logic start------------------------")
	if e3 != nil || e4 != nil {
		//高温
		rows = append(rows, &Das{"3_00019", 0, 1})
		logger.Infof("A high alarm ,reason is [%v,%v].", e3, e4)
	}
	if e2 != nil || e4 != nil {
		//A探头电源关闭
		exist, aPowerOffCount := memcache.GetCurrentDataInt(aPowerOffCountPn)
		logger.Infof("A Cut down count [%d],reason is[%v,%v,%v,%v].", aPowerOffCount, e1, e2, e3, e4)
		if exist {
			if aPowerOffCount > 3 {
				rows = append(rows, &Das{"2_00019", 0, 0})
				logger.Infof("A Real Cut down count [%d],reason is [%v,%v].", aPowerOffCount, e2, e4)
				aPowerOffCount = 0
			} else {
				aPowerOffCount++
			}
		} else {
			aPowerOffCount = 1
		}
		memcache.SetCurrentDataInt(aPowerOffCountPn, aPowerOffCount)

	} else {
		//A探头灯正常
		rows = append(rows, &Das{"3_00017", 0, 1})
	}
	if (e1 != nil || e2 != nil || e3 != nil || e4 != nil) && at != 1 {
		//A镜头退出
		exist, aReturnBackCount := memcache.GetCurrentDataInt(aReturnBackCountPn)
		logger.Infof("A return back count [%d],reason is[%v,%v,%v,%v].", aReturnBackCount, e1, e2, e3, e4)
		if exist {
			if aReturnBackCount > 3 {
				rows = append(rows, &Das{"2_00017", 0, 0})
				logger.Infof("A Real return back count [%d],reason is[%v,%v,%v,%v].", aReturnBackCount, e1, e2, e3, e4)
				aReturnBackCount = 0
			} else {
				aReturnBackCount++
			}
		} else {
			aReturnBackCount = 1
		}
		memcache.SetCurrentDataInt(aPowerOffCountPn, aReturnBackCount)

	}
	logger.Debug("auto logic end------------------------")

	//b探头暂停
	//if e2 != nil || e6 != nil {
	//	//B探头电源关闭
	//	rows = append(rows, &Das{"2_00020", 0, 0})
	//}
	//if e1 != nil || e2 != nil || e5 != nil || bt!= 1{
	//	//B镜头退出
	//	rows = append(rows, &Das{"2_00018", 0, 0})
	//}

	if !config.AppConfig.TempvisionDas.Disable {
		if len(rows) > 0 {
			pushDas := PushDas{rows, 0}
			pushDasBytes, err := json.Marshal(pushDas)
			if err != nil {
				logger.Warnf("das dat to send ddi marshal error %v.", err)
			} else {
				begin := time.Now().UnixNano()
				cmdPostDDI(string(pushDasBytes))
				logger.Debugf("post cmd to ddi const time %d", (time.Now().UnixNano()-begin)/1000000)
			}
		}
	}

}
