package v1

import (
	"dev/test-x-tech/internal/entity"
	"encoding/json"
	"github.com/uptrace/bunrouter"
	"net/http"
)

// lastBTCUSDHandler
// @Summary Последняя запись
// @Tags Валюты
// @Description Получить последний актуальный курс BTC_USDT
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.BTCUSDTResponse
// @Failure 500
// @Router /v1/api/btcusdt [GET]
func (h *Handler) lastBTCUSDHandler(w http.ResponseWriter, req bunrouter.Request) error {
	result, err := h.services.LastBtcUsd()
	if err != nil {
		return h.responseJSON(w, req, http.StatusBadRequest, err)
	}

	return h.responseJSON(w, req, http.StatusOK, result)
}

// historyBTCUSDHandler
// @Summary история Изменений
// @Tags Валюты
// @Description Получить историю изменений, поддерживает фильтрацию по и дате и времени.
// @Description Ввести дату и время в поле Date в формате yyy-mm-dd hh:mm:ss
// @Accept  json
// @Produce  json
// @Param input body entity.BTCUSDBody true "info"
// @Success 200 {array} entity.BTCUSDTResponse
// @Failure 500
// @Router /v1/api/btcusdt [POST]
func (h *Handler) historyBTCUSDHandler(w http.ResponseWriter, req bunrouter.Request) error {
	body := req.Body
	defer req.Body.Close()

	var b entity.BTCUSDBody
	if err := json.NewDecoder(body).Decode(&b); err != nil {
		return h.responseJSON(w, req, http.StatusNotImplemented, err)
	}

	result, err := h.services.HistoryBtcUsd(b.Date, b.Limit, b.Page)
	if err != nil {
		return h.responseJSON(w, req, http.StatusBadRequest, err)
	}

	return h.responseJSON(w, req, http.StatusOK, result)
}

//todo заменить структуру на фильтр как в структур парт
