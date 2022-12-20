package repository

import (
	"context"
	"dev/test-x-tech/internal/entity"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
)

type LatestPostgres struct {
	db *pgxpool.Pool
}

func NewLatest(db *pgxpool.Pool) *LatestPostgres {
	return &LatestPostgres{db: db}
}

func (l *LatestPostgres) LatestFiat() (entity.CurrencyDec, error) {
	var responce entity.CurrencyDec

	tx, err := l.db.Begin(context.Background())
	if err != nil {
		return responce, err
	}
	defer tx.Rollback(context.Background())

	query := queryLast

	err = tx.QueryRow(context.Background(), query).Scan(
		&responce.Date, &responce.AUD, &responce.AZN, &responce.GBP,
		&responce.AMD, &responce.BYN, &responce.BGN, &responce.BRL,
		&responce.HUF, &responce.HKD, &responce.DKK, &responce.USD,
		&responce.EUR, &responce.INR, &responce.KZT, &responce.CAD,
		&responce.KGS, &responce.CNY, &responce.MDL, &responce.NOK,
		&responce.PLN, &responce.RON, &responce.XDR, &responce.SGD,
		&responce.TJS, &responce.TRY, &responce.TMT, &responce.UZS,
		&responce.UAH, &responce.CZK, &responce.SEK, &responce.CHF,
		&responce.ZAR, &responce.KRW, &responce.JPY)
	if err != nil {
		fmt.Println(err)
		return entity.CurrencyDec{}, err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return responce, err
	}

	return responce, nil
}

func (l *LatestPostgres) LatestBtcUsd() (decimal.Decimal, error) {
	query := `SELECT value FROM btc_usd ORDER BY time_get DESC LIMIT 1`

	var value decimal.Decimal
	err := l.db.QueryRow(context.Background(), query).Scan(&value)
	if err != nil {
		return decimal.Zero, fmt.Errorf("не удалось получить значение курса. ошибка - %w", err)
	}
	return value, nil
}
