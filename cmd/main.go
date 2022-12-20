package main

import (
	_ "dev/test-x-tech/docs"
	"dev/test-x-tech/internal/app"
	"dev/test-x-tech/pkg/config"
	"github.com/sirupsen/logrus"
)

var cfg config.Config

func init() {
	err := cfg.InitCfg()
	if err != nil {
		panic(err)
	}
}

// @title ITSM API
// @version 1.0
// @description API Server for ITSM
// @BasePath /
func main() {
	err := app.Run(cfg)
	if err != nil {
		logrus.Error(err.Error())
	}
}
