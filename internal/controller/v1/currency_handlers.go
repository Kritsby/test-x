package v1

import (
	"dev/test-x-tech/internal/entity"
	"encoding/json"
	"github.com/uptrace/bunrouter"
	"net/http"
)

// lastCurrencyHandler
// @Summary Последняя запись
// @Tags Валюты
// @Description Получить последние актиальные курсы валют
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.CurrencyDec
// @Failure 500
// @Router /v1/api/latest [GET]
func (h *Handler) lastCurrencyHandler(w http.ResponseWriter, req bunrouter.Request) error {
	last, err := h.services.LastCurrency()
	if err != nil {
		return h.responseJSON(w, req, http.StatusBadRequest, err)
	}

	return h.responseJSON(w, req, http.StatusOK, last)
}

// historyCurrencyHandler
// @Summary история Изменений
// @Tags Валюты
// @Description Получить историю изменений поддерживает фильтрацию по валютам и дате
// @Description чтобы использовать нужно указать через запятую в двойных кавычках и заглавными буквами
// @Description названия валют в поле Name, а так же дату в поле Date в формате yyy-mm-dd
// @Accept  json
// @Produce  json
// @Param input body entity.CurrencyBody true "info"
// @Success 200 {array} entity.CurrencyResponse
// @Failure 500
// @Router /v1/api/currencies [POST]
func (h *Handler) historyCurrencyHandler(w http.ResponseWriter, req bunrouter.Request) error {
	body := req.Body
	defer req.Body.Close()

	var currency entity.CurrencyBody
	if err := json.NewDecoder(body).Decode(&currency); err != nil {
		return h.responseJSON(w, req, http.StatusNotImplemented, err)
	}

	var currencySlice []string

	for _, v := range currency.Name {
		currencySlice = append(currencySlice, v)
	}

	history, err := h.services.HistoryCurrency(currency.Date, currency.Limit, currency.Page, currencySlice)
	if err != nil {
		return h.responseJSON(w, req, http.StatusBadRequest, err)
	}

	return h.responseJSON(w, req, http.StatusOK, history)
}
