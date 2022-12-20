package app

import (
	"context"
	"dev/test-x-tech/internal/controller/v1"
	"dev/test-x-tech/internal/repository"
	"dev/test-x-tech/internal/service"
	"dev/test-x-tech/pkg/config"
	"dev/test-x-tech/pkg/postgres"
	"dev/test-x-tech/pkg/server"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg config.Config) error {
	db, err := postgres.NewPostgresql(cfg.PSQL)
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
		logrus.Info("соединения с базами данных закрыты")
	}()

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := v1.NewHandler(services)

	//init server
	srv := new(server.Server)
	defer func() {
		_ = srv.Shutdown(context.Background())
		logrus.Info("сервер остановлен")
	}()
	go func() {
		if err = srv.Run(cfg.Server.Port, handlers.InitRouter()); err != nil {
			logrus.Fatalf("возникла ошибка при работе сервера: %s", err.Error())
		}
	}()

	if err = services.TakeCurrency(24 * time.Hour); err != nil {
		logrus.Fatalf("возникла ошибка при получении курсов валют - %s", err.Error())
	}

	if err = services.TakeBtcUsd(10 * time.Second); err != nil {
		logrus.Info()
	}

	//shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("App Shutting Down")

	return err
}
