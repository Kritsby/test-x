package repository

import (
	"dev/test-x-tech/internal/entity"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
)

type BtcUsder interface {
	InsertBtcUsd(usd decimal.Decimal, guid string) error
	LastBtcUsd() (entity.BTCUSDTResult, error)
	HistoryBtcUsd(dateTime string, limit, page int) (entity.BTCUSDTResponse, error)
}

type Currencer interface {
	InsertCurrency(val *entity.CurrencyDec, guid string) error
	HistoryCurrency(date string, limit int, page int) (entity.CurrencyResponse, error)
	LastCurrency() (entity.CurrencyDec, error)
}

type Latest interface {
	LatestFiat() (entity.CurrencyDec, error)
	LatestBtcUsd() (decimal.Decimal, error)
}

type Repository struct {
	BtcUsder
	Currencer
	Latest
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		BtcUsder:  NewBtcUsd(db),
		Currencer: NewCurrencyPostgres(db),
		Latest:    NewLatest(db),
	}
}
