package v1

import (
	"bytes"
	"dev/test-x-tech/internal/entity"
	"dev/test-x-tech/internal/service"
	mock_service "dev/test-x-tech/internal/service/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bunrouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_currency(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCurrencer)

	tests := []struct {
		name                 string
		method               string
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "OK",
			method: "GET",
			mockBehavior: func(s *mock_service.MockCurrencer) {
				s.EXPECT().LastCurrency().Return(entity.CurrencyDec{}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"data":{"AUD":"0","AZN":"0","GBP":"0","AMD":"0","BYN":"0","BGN":"0","BRL":"0","HUF":"0","HKD":"0","DKK":"0","USD":"0","EUR":"0","INR":"0","KZT":"0","CAD":"0","KGS":"0","CNY":"0","MDL":"0","NOK":"0","PLN":"0","RON":"0","XDR":"0","SGD":"0","TJS":"0","TRY":"0","TMT":"0","UZS":"0","UAH":"0","CZK":"0","SEK":"0","CHF":"0","ZAR":"0","KRW":"0","JPY":"0"},"params":null,"route":"/latest"}
`,
		},
		{
			name:      "OK History",
			method:    "POST",
			inputBody: `{"Date": "", "Limit": 0, "Page": 0}`,
			mockBehavior: func(s *mock_service.MockCurrencer) {
				s.EXPECT().HistoryCurrency("", 0, 0, nil).Return(entity.CurrencyResponseInt{}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"data":{"total":0,"history":null},"params":null,"route":"/history"}
`,
		},
		{
			name:   "Service Error",
			method: "GET",
			mockBehavior: func(s *mock_service.MockCurrencer) {
				s.EXPECT().LastCurrency().Return(entity.CurrencyDec{}, errors.New("something went wrong"))
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{"error":{},"params":null,"route":"/latest"}
`,
		},
		{
			name:      "Service Error History",
			method:    "POST",
			inputBody: `{"Date": "", "Limit": 0, "Page": 0}`,
			mockBehavior: func(s *mock_service.MockCurrencer) {
				s.EXPECT().HistoryCurrency("", 0, 0, nil).Return(entity.CurrencyResponseInt{}, errors.New("something went wrong"))
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{"error":{},"params":null,"route":"/history"}
`,
		},
		{
			name:               "No body",
			method:             "POST",
			mockBehavior:       func(s *mock_service.MockCurrencer) {},
			expectedStatusCode: 501,
			expectedResponseBody: `{"error":{},"params":null,"route":"/history"}
`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockCurrencer(c)
			testCase.mockBehavior(repo)

			services := &service.Service{
				Currencer: repo,
			}
			handlers := NewHandler(services)

			r := bunrouter.New()
			var w *httptest.ResponseRecorder
			var req *http.Request
			switch testCase.method {
			case "GET":
				r.GET("/latest", handlers.lastCurrencyHandler)

				w = httptest.NewRecorder()
				req = httptest.NewRequest("GET", "/latest",
					nil)
			case "POST":
				r.POST("/history", handlers.historyCurrencyHandler)

				w = httptest.NewRecorder()
				req = httptest.NewRequest("POST", "/history",
					bytes.NewBufferString(testCase.inputBody))
			}

			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
