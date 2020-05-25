package main

import (
	"fmt"

	"github.com/devfabric/HP-Log/config"
	"github.com/devfabric/HP-Log/logger"
	model1 "github.com/devfabric/HP-Log/mod1"
	model2 "github.com/devfabric/HP-Log/mod2"
)

func main() {
	logConfig, err := config.LoadHPLogConfig("./")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(logConfig)

	hpLog, logger := logger.InitLogger("main")

	logger.Debug("main ", "begin main")
	logger.Info("main ", "end main")

	model1.Fun1(hpLog)
	model2.Fun2(hpLog)
}
