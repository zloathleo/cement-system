package systemalarm

import (
	"github.com/zloathleo/go-httpserver/common/logger"
)

func AddSystemAlarm(err error) {
	logger.Warnf("system alarm [%v]", err)
}
