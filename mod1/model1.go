package model1

import (
	"github.com/devfabric/HP-Log/logger"
)

func Fun1(hpLog *logger.HPLog) {
	logger := hpLog.GetLogger("model1")
	logger.Debug("model1 ", "begin Func1")
	logger.Info("model1 ", "end Func1")
}
