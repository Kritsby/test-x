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
		RUB: btcUsd.Mul(lastFiat.USD).Round(3),
		AUD: lastFiat.AUD.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		AZN: lastFiat.AZN.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		GBP: lastFiat.GBP.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		AMD: lastFiat.AMD.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		BYN: lastFiat.BYN.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		BGN: lastFiat.BGN.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		BRL: lastFiat.BRL.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		HUF: lastFiat.HUF.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		HKD: lastFiat.HKD.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		DKK: lastFiat.DKK.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		USD: btcUsd,
		EUR: lastFiat.EUR.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		INR: lastFiat.INR.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		KZT: lastFiat.KZT.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		CAD: lastFiat.CAD.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		KGS: lastFiat.KGS.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		CNY: lastFiat.CNY.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		MDL: lastFiat.MDL.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		NOK: lastFiat.NOK.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		PLN: lastFiat.PLN.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		RON: lastFiat.RON.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		XDR: lastFiat.XDR.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		SGD: lastFiat.SGD.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		TJS: lastFiat.TJS.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		TRY: lastFiat.TRY.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		TMT: lastFiat.TMT.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		UZS: lastFiat.UZS.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		UAH: lastFiat.UAH.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		CZK: lastFiat.CZK.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		SEK: lastFiat.SEK.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		CHF: lastFiat.CHF.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		ZAR: lastFiat.ZAR.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		KRW: lastFiat.KRW.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
		JPY: lastFiat.JPY.Mul(btcUsd.Div(lastFiat.USD)).Round(3),
	}

	return result, nil
}
