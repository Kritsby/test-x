package entity

import (
	"encoding/xml"
	"github.com/shopspring/decimal"
	"time"
)

type BTSUSDJSON struct {
	Data struct {
		Buy string `json:"buy"`
	} `json:"data"`
}
type BTCUSDTResponse struct {
	Total   int             `json:"total"`
	History []BTCUSDTResult `json:"history"`
}

type BTCUSDTResult struct {
	Timestamp time.Time       `json:"timestamp"`
	Value     decimal.Decimal `json:"value" swaggertype:"string"`
}

type BTCUSDBody struct {
	Date  string `json:"date"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

type CurrencyResponseInt struct {
	Total   int           `json:"total"`
	History []CurrencyInt `json:"history"`
}

type CurrencyResponse struct {
	Total   int           `json:"total"`
	History []CurrencyDec `json:"history"`
}

type CurrencyDec struct {
	Date interface{}     `json:"Date,omitempty" swaggertype:"string"`
	AUD  decimal.Decimal `json:"AUD,omitempty"`
	AZN  decimal.Decimal `json:"AZN,omitempty"`
	GBP  decimal.Decimal `json:"GBP,omitempty"`
	AMD  decimal.Decimal `json:"AMD,omitempty"`
	BYN  decimal.Decimal `json:"BYN,omitempty"`
	BGN  decimal.Decimal `json:"BGN,omitempty"`
	BRL  decimal.Decimal `json:"BRL,omitempty"`
	HUF  decimal.Decimal `json:"HUF,omitempty"`
	HKD  decimal.Decimal `json:"HKD,omitempty"`
	DKK  decimal.Decimal `json:"DKK,omitempty"`
	USD  decimal.Decimal `json:"USD,omitempty"`
	EUR  decimal.Decimal `json:"EUR,omitempty"`
	INR  decimal.Decimal `json:"INR,omitempty"`
	KZT  decimal.Decimal `json:"KZT,omitempty"`
	CAD  decimal.Decimal `json:"CAD,omitempty"`
	KGS  decimal.Decimal `json:"KGS,omitempty"`
	CNY  decimal.Decimal `json:"CNY,omitempty"`
	MDL  decimal.Decimal `json:"MDL,omitempty"`
	NOK  decimal.Decimal `json:"NOK,omitempty"`
	PLN  decimal.Decimal `json:"PLN,omitempty"`
	RON  decimal.Decimal `json:"RON,omitempty"`
	XDR  decimal.Decimal `json:"XDR,omitempty"`
	SGD  decimal.Decimal `json:"SGD,omitempty"`
	TJS  decimal.Decimal `json:"TJS,omitempty"`
	TRY  decimal.Decimal `json:"TRY,omitempty"`
	TMT  decimal.Decimal `json:"TMT,omitempty"`
	UZS  decimal.Decimal `json:"UZS,omitempty"`
	UAH  decimal.Decimal `json:"UAH,omitempty"`
	CZK  decimal.Decimal `json:"CZK,omitempty"`
	SEK  decimal.Decimal `json:"SEK,omitempty"`
	CHF  decimal.Decimal `json:"CHF,omitempty"`
	ZAR  decimal.Decimal `json:"ZAR,omitempty"`
	KRW  decimal.Decimal `json:"KRW,omitempty"`
	JPY  decimal.Decimal `json:"JPY,omitempty"`
	RUB  decimal.Decimal `json:",omitempty"`
}

type CurrencyInt struct {
	Date string      `json:"Date"`
	AUD  interface{} `json:"AUD,omitempty"`
	AZN  interface{} `json:"AZN,omitempty"`
	GBP  interface{} `json:"GBP,omitempty"`
	AMD  interface{} `json:"AMD,omitempty"`
	BYN  interface{} `json:"BYN,omitempty"`
	BGN  interface{} `json:"BGN,omitempty"`
	BRL  interface{} `json:"BRL,omitempty"`
	HUF  interface{} `json:"HUF,omitempty"`
	HKD  interface{} `json:"HKD,omitempty"`
	DKK  interface{} `json:"DKK,omitempty"`
	USD  interface{} `json:"USD,omitempty"`
	EUR  interface{} `json:"EUR,omitempty"`
	INR  interface{} `json:"INR,omitempty"`
	KZT  interface{} `json:"KZT,omitempty"`
	CAD  interface{} `json:"CAD,omitempty"`
	KGS  interface{} `json:"KGS,omitempty"`
	CNY  interface{} `json:"CNY,omitempty"`
	MDL  interface{} `json:"MDL,omitempty"`
	NOK  interface{} `json:"NOK,omitempty"`
	PLN  interface{} `json:"PLN,omitempty"`
	RON  interface{} `json:"RON,omitempty"`
	XDR  interface{} `json:"XDR,omitempty"`
	SGD  interface{} `json:"SGD,omitempty"`
	TJS  interface{} `json:"TJS,omitempty"`
	TRY  interface{} `json:"TRY,omitempty"`
	TMT  interface{} `json:"TMT,omitempty"`
	UZS  interface{} `json:"UZS,omitempty"`
	UAH  interface{} `json:"UAH,omitempty"`
	CZK  interface{} `json:"CZK,omitempty"`
	SEK  interface{} `json:"SEK,omitempty"`
	CHF  interface{} `json:"CHF,omitempty"`
	ZAR  interface{} `json:"ZAR,omitempty"`
	KRW  interface{} `json:"KRW,omitempty"`
	JPY  interface{} `json:"JPY,omitempty"`
}

type CurrencyBody struct {
	Date  string   `json:"date"`
	Name  []string `json:"name"`
	Page  int      `json:"page"`
	Limit int      `json:"limit"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valute  []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName  xml.Name `xml:"Valute"`
	CharCode string   `xml:"CharCode"`
	Nominal  string   `xml:"Nominal"`
	Name     string   `xml:"Name"`
	Value    string   `xml:"Value"`
}
