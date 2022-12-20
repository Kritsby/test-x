package repository

import (
	"context"
	"dev/test-x-tech/internal/entity"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CurrencyPostgres struct {
	db *pgxpool.Pool
}

func NewCurrencyPostgres(db *pgxpool.Pool) *CurrencyPostgres {
	return &CurrencyPostgres{db: db}
}

func (b *CurrencyPostgres) LastCurrency() (entity.CurrencyDec, error) {
	var responce entity.CurrencyDec

	tx, err := b.db.Begin(context.Background())
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

func (b *CurrencyPostgres) HistoryCurrency(date string, limit int, page int) (entity.CurrencyResponse, error) {
	var responce entity.CurrencyResponse

	tx, err := b.db.Begin(context.Background())
	if err != nil {
		return entity.CurrencyResponse{}, err
	}
	defer tx.Rollback(context.Background())

	query := `
	SELECT
	    date_get, "AUD", "AZN", "GBP", "AMD",
	    "BYN", "BGN", "BRL", "HUF", "HKD", "DKK",
	    "USD", "EUR", "INR", "KZT", "CAD", "KGS",
	    "CNY", "MDL", "NOK", "PLN", "RON", "XDR",
	    "SGD", "TJS", "TRY", "TMT", "UZS", "UAH", 
	    "CZK", "SEK", "CHF", "ZAR", "KRW", "JPY"
	FROM
	    currency
	WHERE CAST(date_get AS TEXT)  like '%` + date + `%'
	ORDER BY date_get DESC
	LIMIT $1
	OFFSET $2;`

	rows, err := tx.Query(context.Background(), query, limit, page)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return responce, err
		}
	}

	var history entity.CurrencyDec

	var i int

	for rows.Next() {
		err = rows.Scan(
			&history.Date, &history.AUD, &history.AZN, &history.GBP, &history.AMD,
			&history.BYN, &history.BGN, &history.BRL, &history.HUF, &history.HKD,
			&history.DKK, &history.USD, &history.EUR, &history.INR, &history.KZT,
			&history.CAD, &history.KGS, &history.CNY, &history.MDL, &history.NOK,
			&history.PLN, &history.RON, &history.XDR, &history.SGD, &history.TJS,
			&history.TRY, &history.TMT, &history.UZS, &history.UAH, &history.CZK,
			&history.SEK, &history.CHF, &history.ZAR, &history.KRW, &history.JPY)
		if err != nil {
			fmt.Println(err)
			return responce, err
		}

		i += 1

		responce.History = append(responce.History, history)
	}

	var curInt entity.CurrencyInt

	curInt.AUD = history.AUD

	responce.Total = i

	if err = tx.Commit(context.Background()); err != nil {
		return responce, err
	}

	return responce, nil
}

func (b *CurrencyPostgres) InsertCurrency(val *entity.CurrencyDec, guid string) error {
	tx, err := b.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), queryInsert,
		guid,
		val.AUD, val.AZN, val.AMD, val.GBP,
		val.BYN, val.BGN, val.BRL, val.HUF,
		val.HKD, val.DKK, val.USD, val.EUR,
		val.INR, val.KZT, val.CAD, val.KGS,
		val.CNY, val.MDL, val.NOK, val.PLN,
		val.RON, val.XDR, val.SGD, val.TJS,
		val.TRY, val.TMT, val.UZS, val.UAH,
		val.CZK, val.SEK, val.CHF, val.ZAR,
		val.KRW, val.JPY)
	if err != nil {
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func withoutAnyFilters(tx pgx.Tx, limit int, page int) (entity.CurrencyResponse, error) {
	var responce entity.CurrencyResponse

	query := queryWithoutFilters

	rows, err := tx.Query(context.Background(), query, limit, page)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return responce, err
		}
	}

	var history entity.CurrencyDec

	var i int

	for rows.Next() {
		err = rows.Scan(
			&history.Date, &history.AUD, &history.AZN, &history.GBP,
			&history.AMD, &history.BYN, &history.BGN, &history.BRL,
			&history.HUF, &history.HKD, &history.DKK, &history.USD,
			&history.EUR, &history.INR, &history.KZT, &history.CAD,
			&history.KGS, &history.CNY, &history.MDL, &history.NOK,
			&history.PLN, &history.RON, &history.XDR, &history.SGD,
			&history.TJS, &history.TRY, &history.TMT, &history.UZS,
			&history.UAH, &history.CZK, &history.SEK, &history.CHF,
			&history.ZAR, &history.KRW, &history.JPY)
		if err != nil {
			return responce, err
		}

		i += 1

		responce.History = append(responce.History, history)
	}

	responce.Total = i

	if err = tx.Commit(context.Background()); err != nil {
		return responce, err
	}

	return responce, nil
}

func withDateFilter(tx pgx.Tx, date string, limit int, page int) (entity.CurrencyResponse, error) {
	var responce entity.CurrencyResponse

	return responce, nil
}

const (
	queryLast = `
	SELECT
	    date_get,
	    "AUD",
	    "AZN",
	    "GBP",
	    "AMD",
	    "BYN",
	    "BGN",
	    "BRL",
	    "HUF",
	    "HKD",
	    "DKK",
	    "USD",
	    "EUR",
	    "INR",
	    "KZT",
	    "CAD",
	    "KGS",
	    "CNY",
	    "MDL", 
	    "NOK",
	    "PLN",
	    "RON",
	    "XDR",
	    "SGD", 
	    "TJS", 
	    "TRY",
	    "TMT",
	    "UZS", 
	    "UAH", 
	    "CZK", 
	    "SEK", 
	    "CHF",
	    "ZAR",
	    "KRW",
	    "JPY"
	FROM
	    currency
	WHERE
	    date_get = current_date`

	queryInsert = `
	INSERT INTO currency(GUID,
	                     DATE_GET,
	                     "AUD",
	                     "AZN",
	                     "AMD",
	                     "GBP",
	                     "BYN",
	                     "BGN", 
	                     "BRL",
	                     "HUF",
	                     "HKD",
	                     "DKK",
	                     "USD",
	                     "EUR",
	                     "INR",
	                     "KZT",
	                     "CAD",
	                     "KGS",
	                     "CNY",
	                     "MDL",
	                     "NOK",
	                     "PLN",
	                     "RON",
	                     "XDR",
	                     "SGD",
	                     "TJS",
	                     "TRY",
	                     "TMT",
	                     "UZS",
	                     "UAH",
	                     "CZK",
	                     "SEK",
	                     "CHF",
	                     "ZAR",
	                     "KRW",
	                     "JPY")
	VALUES($1, CURRENT_DATE, $2, $3, $4, $5, $6, $7, $8, $9, $10,
	       $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
	       $21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
	       $31, $32, $33, $34, $35) ON CONFLICT (date_get) DO NOTHING;`

	queryWithoutFilters = `
	SELECT
	    date_get,
	    "AUD",
	    "AZN",
	    "GBP",
	    "AMD",
	    "BYN",
	    "BGN",
	    "BRL",
	    "HUF",
	    "HKD",
	    "DKK",
	    "USD",
	    "EUR",
	    "INR",
	    "KZT",
	    "CAD",
	    "KGS",
	    "CNY",
	    "MDL", 
	    "NOK",
	    "PLN",
	    "RON",
	    "XDR",
	    "SGD", 
	    "TJS", 
	    "TRY",
	    "TMT",
	    "UZS", 
	    "UAH", 
	    "CZK", 
	    "SEK", 
	    "CHF",
	    "ZAR",
	    "KRW",
	    "JPY"
	FROM
	    currency
	ORDER BY date_get DESC 
	LIMIT $1
	OFFSET $2;`
)
