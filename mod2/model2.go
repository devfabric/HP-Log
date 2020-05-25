package model2

import (
	"github.com/devfabric/HP-Log/logger"
)

func Fun2(hpLog *logger.HPLog) {
	logger := hpLog.GetLogger("model2")
	logger.Debug("model2 ", "begin Func2")
	logger.Info("model2 ", "end Func2")
}
