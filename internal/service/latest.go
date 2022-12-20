package service

import (
	"dev/test-x-tech/internal/entity"
	"dev/test-x-tech/internal/repository"
)

type LatestService struct {
	repo repository.Latest
}

func NewLatestService(repo repository.Latest) *LatestService {
	return &LatestService{repo: repo}
}

func (l *LatestService) Latest() (entity.CurrencyDec, error) {
	lastFiat, err := l.repo.LatestFiat()
	if err != nil {
		return entity.CurrencyDec{}, err
	}

	btcUsd, err := l.repo.LatestBtcUsd()
	if err != nil {
		return entity.CurrencyDec{}, err
	}

	result := entity.CurrencyDec{
		RUB: lastFiat.AUD.Mul(btcUsd.Div(lastFiat.USD)),
		AUD: lastFiat.AUD.Mul(btcUsd.Div(lastFiat.USD)),
		AZN: lastFiat.AZN.Mul(btcUsd.Div(lastFiat.USD)),
		GBP: lastFiat.GBP.Mul(btcUsd.Div(lastFiat.USD)),
		AMD: lastFiat.AMD.Mul(btcUsd.Div(lastFiat.USD)),
		BYN: lastFiat.BYN.Mul(btcUsd.Div(lastFiat.USD)),
		BGN: lastFiat.BGN.Mul(btcUsd.Div(lastFiat.USD)),
		BRL: lastFiat.BRL.Mul(btcUsd.Div(lastFiat.USD)),
		HUF: lastFiat.HUF.Mul(btcUsd.Div(lastFiat.USD)),
		HKD: lastFiat.HKD.Mul(btcUsd.Div(lastFiat.USD)),
		DKK: lastFiat.DKK.Mul(btcUsd.Div(lastFiat.USD)),
		USD: btcUsd,
		EUR: lastFiat.EUR.Mul(btcUsd.Div(lastFiat.USD)),
		INR: lastFiat.INR.Mul(btcUsd.Div(lastFiat.USD)),
		KZT: lastFiat.KZT.Mul(btcUsd.Div(lastFiat.USD)),
		CAD: lastFiat.CAD.Mul(btcUsd.Div(lastFiat.USD)),
		KGS: lastFiat.KGS.Mul(btcUsd.Div(lastFiat.USD)),
		CNY: lastFiat.CNY.Mul(btcUsd.Div(lastFiat.USD)),
		MDL: lastFiat.MDL.Mul(btcUsd.Div(lastFiat.USD)),
		NOK: lastFiat.NOK.Mul(btcUsd.Div(lastFiat.USD)),
		PLN: lastFiat.PLN.Mul(btcUsd.Div(lastFiat.USD)),
		RON: lastFiat.RON.Mul(btcUsd.Div(lastFiat.USD)),
		XDR: lastFiat.XDR.Mul(btcUsd.Div(lastFiat.USD)),
		SGD: lastFiat.SGD.Mul(btcUsd.Div(lastFiat.USD)),
		TJS: lastFiat.TJS.Mul(btcUsd.Div(lastFiat.USD)),
		TRY: lastFiat.TRY.Mul(btcUsd.Div(lastFiat.USD)),
		TMT: lastFiat.TMT.Mul(btcUsd.Div(lastFiat.USD)),
		UZS: lastFiat.UZS.Mul(btcUsd.Div(lastFiat.USD)),
		UAH: lastFiat.UAH.Mul(btcUsd.Div(lastFiat.USD)),
		CZK: lastFiat.CZK.Mul(btcUsd.Div(lastFiat.USD)),
		SEK: lastFiat.SEK.Mul(btcUsd.Div(lastFiat.USD)),
		CHF: lastFiat.CHF.Mul(btcUsd.Div(lastFiat.USD)),
		ZAR: lastFiat.ZAR.Mul(btcUsd.Div(lastFiat.USD)),
		KRW: lastFiat.KRW.Mul(btcUsd.Div(lastFiat.USD)),
		JPY: lastFiat.JPY.Mul(btcUsd.Div(lastFiat.USD)),
	}

	return result, nil
}
