package v1

import (
	"dev/test-x-tech/internal/service"
	"fmt"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter() *bunrouter.Router {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
	)

	swagHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)
	bswag := bunrouter.HTTPHandlerFunc(swagHandler)
	router.GET("/swagger/:*", bswag)

	router.WithGroup("/v1/api", func(g *bunrouter.Group) {
		g.WithGroup("/btcusdt", func(g *bunrouter.Group) {
			g.GET("", h.lastBTCUSDHandler)
			g.POST("", h.historyBTCUSDHandler)
		})
		g.WithGroup("/currencies", func(g *bunrouter.Group) {
			g.GET("", h.lastCurrencyHandler)
			g.POST("", h.historyCurrencyHandler)
		})
		g.GET("/latest", h.latestHandler)
	})

	return router
}

func (h *Handler) responseJSON(w http.ResponseWriter, req bunrouter.Request, code int, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if code != http.StatusOK {
		logrus.Warningf(fmt.Sprintf("route: %s, http code: %d, error: %v", req.Route(), code, value))
		return bunrouter.JSON(w, bunrouter.H{
			"route":  req.Route(),
			"params": req.Params().Map(),
			"error":  value,
		})
	}

	return bunrouter.JSON(w, bunrouter.H{
		"route":  req.Route(),
		"params": req.Params().Map(),
		"data":   value,
	})
}
