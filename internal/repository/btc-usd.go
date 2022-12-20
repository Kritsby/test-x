package repository

import (
	"context"
	"dev/test-x-tech/internal/entity"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
)

type BtcUsdPostgres struct {
	db *pgxpool.Pool
}

func NewBtcUsd(db *pgxpool.Pool) *BtcUsdPostgres {
	return &BtcUsdPostgres{db: db}
}

func (b *BtcUsdPostgres) InsertBtcUsd(usd decimal.Decimal, guid string) error {
	tx, err := b.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var oldValue decimal.Decimal

	query := `
	SELECT
	    value 
	FROM
	    btc_usd
	ORDER BY time_get DESC
	LIMIT 1`

	err = tx.QueryRow(context.Background(), query).Scan(&oldValue)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
	}

	queryInsert := `
		INSERT INTO
		    btc_usd(guid, time_get, value)
		VALUES($1, current_timestamp, $2)`

	if !oldValue.Equal(usd) {
		_, err = tx.Exec(context.Background(), queryInsert, guid, usd)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func (b *BtcUsdPostgres) LastBtcUsd() (entity.BTCUSDTResult, error) {
	query := `SELECT time_get, value FROM btc_usd ORDER BY time_get DESC LIMIT 1`

	var value entity.BTCUSDTResult
	err := b.db.QueryRow(context.Background(), query).Scan(&value.Timestamp, &value.Value)
	if err != nil {
		return entity.BTCUSDTResult{}, fmt.Errorf("не удалось получить значение курса. ошибка - %w", err)
	}

	return value, nil
}

func (b *BtcUsdPostgres) HistoryBtcUsd(dateTime string, limit, page int) (entity.BTCUSDTResponse, error) {
	var result entity.BTCUSDTResponse
	tx, err := b.db.Begin(context.Background())
	if err != nil {
		return entity.BTCUSDTResponse{}, err
	}
	defer tx.Rollback(context.Background())

	query := `
	SELECT
	    time_get, value
	FROM btc_usd
	WHERE CAST(time_get AS TEXT) LIKE '%` + dateTime + `%'
	ORDER BY time_get DESC
	LIMIT $1
	OFFSET $2;`

	rows, err := tx.Query(context.Background(), query, limit, page)
	if err != nil {
		return entity.BTCUSDTResponse{}, err
	}

	var i int
	var r entity.BTCUSDTResult

	for rows.Next() {
		err = rows.Scan(&r.Timestamp, &r.Value)
		if err != nil {
			return entity.BTCUSDTResponse{}, err
		}

		i += 1
		result.History = append(result.History, r)
	}

	result.Total = i

	if err = rows.Err(); err != nil && err != pgx.ErrNoRows {
		return entity.BTCUSDTResponse{}, fmt.Errorf("не удалось получить список итемов, ошибка получения серверов: %w", err)
	}

	if err = tx.Commit(context.Background()); err != nil {
		return entity.BTCUSDTResponse{}, err
	}

	return result, nil
}
