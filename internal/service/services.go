package service

import (
	"dev/test-x-tech/internal/entity"
	"dev/test-x-tech/internal/repository"
	"time"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type BtcUsder interface {
	TakeBtcUsd(periodicity time.Duration) error
	LastBtcUsd() (*entity.BTCUSDTResult, error)
	HistoryBtcUsd(dateTime string, limit, page int) (*entity.BTCUSDTResponse, error)
}

type Currencer interface {
	TakeCurrency(periodicity time.Duration) error
	HistoryCurrency(date string, limit int, page int, CurrencySlice []string) (entity.CurrencyResponseInt, error)
	LastCurrency() (entity.CurrencyDec, error)
}

type Latester interface {
	Latest() (entity.CurrencyDec, error)
}

type Service struct {
	BtcUsder
	Currencer
	Latester
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		BtcUsder:  NewBtcUsdService(repo.BtcUsder),
		Currencer: NewCurrencyService(repo.Currencer),
		Latester:  NewLatestService(repo.Latest),
	}
}
