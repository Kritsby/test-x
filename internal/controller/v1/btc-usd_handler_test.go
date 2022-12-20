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

func TestHandler_btc_usd(t *testing.T) {
	type mockBehavior func(s *mock_service.MockBtcUsder)

	tests := []struct {
		name                 string
		method               string
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "OK Latest",
			method: "GET",
			mockBehavior: func(s *mock_service.MockBtcUsder) {
				s.EXPECT().LastBtcUsd().Return(&entity.BTCUSDTResult{}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"data":{"timestamp":"","value":"0"},"params":null,"route":"/latest"}
`,
		},
		{
			name:      "OK History",
			method:    "POST",
			inputBody: `{"Date": "", "Limit": 0, "Page": 0}`,
			mockBehavior: func(s *mock_service.MockBtcUsder) {
				s.EXPECT().HistoryBtcUsd("", 0, 0).Return(&entity.BTCUSDTResponse{}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"data":{"total":0,"history":null},"params":null,"route":"/history"}
`,
		},
		{
			name:   "Service Error Latest",
			method: "GET",
			mockBehavior: func(s *mock_service.MockBtcUsder) {
				s.EXPECT().LastBtcUsd().Return(&entity.BTCUSDTResult{}, errors.New("something went wrong"))
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{"error":{},"params":null,"route":"/latest"}
`,
		},
		{
			name:      "Service Error History",
			method:    "POST",
			inputBody: `{"Date": "", "Limit": 0, "Page": 0}`,
			mockBehavior: func(s *mock_service.MockBtcUsder) {
				s.EXPECT().HistoryBtcUsd("", 0, 0).Return(&entity.BTCUSDTResponse{}, errors.New("something went wrong"))
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{"error":{},"params":null,"route":"/history"}
`,
		},
		{
			name:               "No body",
			method:             "POST",
			mockBehavior:       func(s *mock_service.MockBtcUsder) {},
			expectedStatusCode: 501,
			expectedResponseBody: `{"error":{},"params":null,"route":"/history"}
`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockBtcUsder(c)
			testCase.mockBehavior(repo)

			services := &service.Service{
				BtcUsder: repo,
			}
			handlers := NewHandler(services)

			r := bunrouter.New()
			var w *httptest.ResponseRecorder
			var req *http.Request
			switch testCase.method {
			case "GET":
				r.GET("/latest", handlers.lastBTCUSDHandler)

				w = httptest.NewRecorder()
				req = httptest.NewRequest("GET", "/latest",
					nil)
			case "POST":
				r.POST("/history", handlers.historyBTCUSDHandler)

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
