package v1

import (
	"github.com/uptrace/bunrouter"
	"net/http"
)

// latestHandler
// @Summary Последняя запись
// @Tags Валюты
// @Description Получить актуальные крусы BTC к фиатным валютам
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.CurrencyDec
// @Failure 500
// @Router /v1/api/currencies [GET]
func (h *Handler) latestHandler(w http.ResponseWriter, req bunrouter.Request) error {
	result, err := h.services.Latest()
	if err != nil {
		return h.responseJSON(w, req, http.StatusBadRequest, err)
	}

	return h.responseJSON(w, req, http.StatusOK, result)
}
