package v1

import (
	"dev/test-x-tech/internal/entity"
	"dev/test-x-tech/internal/service"
	mock_service "dev/test-x-tech/internal/service/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bunrouter"
	"net/http/httptest"
	"testing"
)

func TestHandler_latest(t *testing.T) {
	type mockBehavior func(s *mock_service.MockLatester)

	tests := []struct {
		name                 string
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockLatester) {
				s.EXPECT().Latest().Return(entity.CurrencyDec{Date: "2022-12-10"}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{"data":{"Date":"2022-12-10","AUD":"0","AZN":"0","GBP":"0","AMD":"0","BYN":"0","BGN":"0","BRL":"0","HUF":"0","HKD":"0","DKK":"0","USD":"0","EUR":"0","INR":"0","KZT":"0","CAD":"0","KGS":"0","CNY":"0","MDL":"0","NOK":"0","PLN":"0","RON":"0","XDR":"0","SGD":"0","TJS":"0","TRY":"0","TMT":"0","UZS":"0","UAH":"0","CZK":"0","SEK":"0","CHF":"0","ZAR":"0","KRW":"0","JPY":"0"},"params":null,"route":"/latest"}
`,
		},
		{
			name: "Service Error",
			mockBehavior: func(s *mock_service.MockLatester) {
				s.EXPECT().Latest().Return(entity.CurrencyDec{}, errors.New("something went wrong"))
			},
			expectedStatusCode: 400,
			expectedResponseBody: `{"error":{},"params":null,"route":"/latest"}
`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockLatester(c)
			testCase.mockBehavior(repo)

			services := &service.Service{
				Latester: repo,
			}
			handlers := NewHandler(services)

			r := bunrouter.New()
			r.GET("/latest", handlers.latestHandler)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/latest",
				nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
