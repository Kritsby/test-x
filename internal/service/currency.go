package service

import (
	"dev/test-x-tech/internal/entity"
	"dev/test-x-tech/internal/repository"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"net/http"
	"strings"
	"sync"
	"time"
)

type CurrencyService struct {
	repo repository.Currencer

	stopChan    chan struct{}
	closeChan   chan struct{}
	isStarted   bool
	muStarted   sync.Mutex
	inProcess   bool
	muProcessed sync.Mutex
}

func NewCurrencyService(repo repository.Currencer) *CurrencyService {
	return &CurrencyService{
		repo: repo,

		stopChan:    make(chan struct{}),
		closeChan:   make(chan struct{}),
		isStarted:   false,
		muStarted:   sync.Mutex{},
		inProcess:   false,
		muProcessed: sync.Mutex{}}
}

func (s *CurrencyService) TakeCurrency(periodicity time.Duration) error {
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
			val, err := parseCurrency()
			if err != nil {
				logrus.Infof("ошибка при попытке парсинга фиатных валют: %v\n", err)
				ticker := time.Tick(time.Second)

				for i := 9; i >= 0; i-- {
					<-ticker
					fmt.Printf("\rПовторная попытка через %v", i)
				}
				fmt.Printf("\nПробуем снова\n")
				continue
			}

			logrus.Info("парсинг - ок")

			guid := uuid.New().String()

			var currency entity.CurrencyDec

			for _, v := range val.Valute {
				val := strings.Replace(v.Value, ",", ".", 1)
				value, err := decimal.NewFromString(val)
				if err != nil {
					return
				}

				switch v.CharCode {
				case "AUD":
					currency.AUD = value
				case "AZN":
					currency.AZN = value
				case "AMD":
					currency.AMD = value
				case "BYN":
					currency.BYN = value
				case "BGN":
					currency.BGN = value
				case "BRL":
					currency.BRL = value
				case "HUF":
					currency.HUF = value
				case "HKD":
					currency.HKD = value
				case "DKK":
					currency.DKK = value
				case "USD":
					currency.USD = value
				case "EUR":
					currency.EUR = value
				case "INR":
					currency.INR = value
				case "KZT":
					currency.KZT = value
				case "CAD":
					currency.CAD = value
				case "GBP":
					currency.GBP = value
				case "KGS":
					currency.KGS = value
				case "CNY":
					currency.CNY = value
				case "MDL":
					currency.MDL = value
				case "NOK":
					currency.NOK = value
				case "PLN":
					currency.PLN = value
				case "RON":
					currency.RON = value
				case "XDR":
					currency.XDR = value
				case "SGD":
					currency.SGD = value
				case "TJS":
					currency.TJS = value
				case "TRY":
					currency.TRY = value
				case "TMT":
					currency.TMT = value
				case "UZS":
					currency.UZS = value
				case "UAH":
					currency.UAH = value
				case "CZK":
					currency.CZK = value
				case "SEK":
					currency.SEK = value
				case "CHF":
					currency.CHF = value
				case "ZAR":
					currency.ZAR = value
				case "KRW":
					currency.KRW = value
				case "JPY":
					currency.JPY = value
				}
			}

			err = s.repo.InsertCurrency(&currency, guid)
			if err != nil {
				logrus.Infof("ошибка при попытке заполнить базу: %v", err)
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

func parseCurrency() (*entity.ValCurs, error) {
	resp, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	val := new(entity.ValCurs)

	d := xml.NewDecoder(resp.Body)
	d.CharsetReader = charset.NewReaderLabel
	err = d.Decode(&val)
	if err != nil {
		return &entity.ValCurs{}, err
	}

	return val, nil
}

func (s *CurrencyService) HistoryCurrency(date string, limit int, page int, currencySlice []string) (entity.CurrencyResponseInt, error) {
	history, err := s.repo.HistoryCurrency(date, limit, page)
	if err != nil {
		return entity.CurrencyResponseInt{}, err
	}

	var curResponse entity.CurrencyResponseInt
	var curInt entity.CurrencyInt
	var curIntSlice []entity.CurrencyInt

	for _, v := range history.History {
		date := fmt.Sprintf("%s", v.Date)
		curInt.Date = date[1:11]
		for _, c := range currencySlice {
			switch c {
			case "AUD":
				curInt.AUD = v.AUD
			case "AZN":
				curInt.AZN = v.AZN
			case "AMD":
				curInt.AMD = v.AMD
			case "BYN":
				curInt.BYN = v.BYN
			case "BGN":
				curInt.BGN = v.BGN
			case "BRL":
				curInt.BRL = v.BRL
			case "HUF":
				curInt.HUF = v.HUF
			case "HKD":
				curInt.HKD = v.HKD
			case "DKK":
				curInt.DKK = v.DKK
			case "USD":
				curInt.USD = v.USD
			case "EUR":
				curInt.EUR = v.EUR
			case "INR":
				curInt.INR = v.INR
			case "KZT":
				curInt.KZT = v.KZT
			case "CAD":
				curInt.CAD = v.CAD
			case "GBP":
				curInt.GBP = v.GBP
			case "KGS":
				curInt.KGS = v.KGS
			case "CNY":
				curInt.CNY = v.CNY
			case "MDL":
				curInt.MDL = v.MDL
			case "NOK":
				curInt.NOK = v.NOK
			case "PLN":
				curInt.PLN = v.PLN
			case "RON":
				curInt.RON = v.RON
			case "XDR":
				curInt.XDR = v.XDR
			case "SGD":
				curInt.SGD = v.SGD
			case "TJS":
				curInt.TJS = v.TJS
			case "TRY":
				curInt.TRY = v.TRY
			case "TMT":
				curInt.TMT = v.TMT
			case "UZS":
				curInt.UZS = v.UZS
			case "UAH":
				curInt.UAH = v.UAH
			case "CZK":
				curInt.CZK = v.CZK
			case "SEK":
				curInt.SEK = v.SEK
			case "CHF":
				curInt.CHF = v.CHF
			case "ZAR":
				curInt.ZAR = v.ZAR
			case "KRW":
				curInt.KRW = v.KRW
			case "JPY":
				curInt.JPY = v.JPY
			default:
				marshalStruc, err := json.Marshal(history)
				if err != nil {
					fmt.Println(err)
				}
				err = json.Unmarshal(marshalStruc, &curResponse)
				if err != nil {
					fmt.Println(err)
				}

				var c entity.CurrencyResponseInt

				for _, v := range curResponse.History {
					date = fmt.Sprintf("%s", v.Date)
					v.Date = date[:10]
					c.History = append(c.History, v)
				}

				return c, nil
			}
		}
		curIntSlice = append(curIntSlice, curInt)
	}

	curResponse.Total = history.Total
	curResponse.History = curIntSlice

	return curResponse, nil
}

func (s *CurrencyService) LastCurrency() (entity.CurrencyDec, error) {
	last, err := s.repo.LastCurrency()
	if err != nil {
		return entity.CurrencyDec{}, err
	}
	return last, nil
}
