package main

import (
	"mywork/go-my-playground/3.advance_app/lib_framwork/zap/config"
	"mywork/go-my-playground/3.advance_app/lib_framwork/zap/logger"
)

func main() {
	config.InitConfig()
	logger.InitLogger()

}
