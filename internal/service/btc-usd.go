package service

import (
	"dev/test-x-tech/internal/entity"
	"dev/test-x-tech/internal/repository"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"sync"
	"time"
)

const layout = "02:01:2006 15:04:05"

type BtcUsdService struct {
	repo repository.BtcUsder

	stopChan    chan struct{}
	closeChan   chan struct{}
	isStarted   bool
	muStarted   sync.Mutex
	inProcess   bool
	muProcessed sync.Mutex
}

func NewBtcUsdService(repo repository.BtcUsder) *BtcUsdService {
	return &BtcUsdService{
		repo: repo,

		stopChan:    make(chan struct{}),
		closeChan:   make(chan struct{}),
		isStarted:   false,
		muStarted:   sync.Mutex{},
		inProcess:   false,
		muProcessed: sync.Mutex{}}
}

func (s *BtcUsdService) TakeBtcUsd(periodicity time.Duration) error {
	s.muStarted.Lock()

	switch {
	case s.isStarted:
		s.muStarted.Unlock()
		return fmt.Errorf("процесс парсинга уже работает")
	default:
		s.isStarted = true
		logrus.Infof("процесс парсинга запущен, периодичность: %v", periodicity)
	}
	s.muStarted.Unlock()

	go func() {
		defer close(s.closeChan)

		for s.isStarted {
			btcUsd, err := parseBTCUSD()
			if err != nil {
				logrus.Infof("ошибка при попытке парсинга BTC-USDT: %v\n", err.Error())
				ticker := time.Tick(time.Second)

				for i := 9; i >= 0; i-- {
					<-ticker
					fmt.Printf("\rПовторная попытка через %v", i)
				}
				fmt.Printf("\nПробуем снова\n")
				continue
			}

			guid := uuid.New().String()

			err = s.repo.InsertBtcUsd(btcUsd, guid)
			if err != nil {
				logrus.Infof("ошибка при попытке вставить значение курса: %s", err.Error())
			}

			select {
			case <-s.stopChan:
				return
			case <-time.After(periodicity):
			}
		}
	}()

	return nil
}

func parseBTCUSD() (decimal.Decimal, error) {
	resp, _ := http.Get("https://api.kucoin.com/api/v1/market/stats?symbol=BTC-USDT")
	defer resp.Body.Close()

	a := new(entity.BTSUSDJSON)

	_ = json.NewDecoder(resp.Body).Decode(a)

	val := strings.Replace(a.Data.Buy, ",", ".", 1)
	price, err := decimal.NewFromString(val)
	if err != nil {
		return decimal.Zero, err
	}

	return price, nil
}

func (s *BtcUsdService) LastBtcUsd() (*entity.BTCUSDTResult, error) {
	last, err := s.repo.LastBtcUsd()
	if err != nil {
		return &entity.BTCUSDTResult{}, err
	}

	result := new(entity.BTCUSDTResult)

	date := fmt.Sprintf("%s", result.Timestamp)
	dateTime, err := time.Parse(layout, date)
	result.Timestamp = dateTime
	result.Value = last.Value

	return result, nil
}

func (s *BtcUsdService) HistoryBtcUsd(dateTime string, limit, page int) (*entity.BTCUSDTResponse, error) {
	history, err := s.repo.HistoryBtcUsd(dateTime, limit, page)
	if err != nil {
		return &entity.BTCUSDTResponse{}, err
	}

	result := new(entity.BTCUSDTResponse)
	var historySlice []entity.BTCUSDTResult

	for _, h := range history.History {
		date := fmt.Sprintf("%s", h.Timestamp)
		dateTime, err := time.Parse(layout, date)
		if err != nil {
			return nil, err
		}
		historySlice = append(historySlice, entity.BTCUSDTResult{
			Timestamp: dateTime,
			Value:     h.Value,
		})
	}

	result.Total = history.Total
	result.History = historySlice

	return result, nil
}
