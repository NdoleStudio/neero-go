package neero

import (
	"context"
	"encoding/json"
	"net/http"
)

// balanceService is the API client for the `/payment-gateway/api/v1/balances/` endpoint
type balanceService service

// Get balance By Payment Method Id
//
// API Docs: https://app.theneo.io/neero/neero-payment-gateway/balances/get-balance-by-payment-method-id
func (service *balanceService) Get(ctx context.Context, paymentMethodID string) (*BalanceResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/payment-gateway/api/v1/balances/payment-method/"+paymentMethodID, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	balance := new(BalanceResponse)
	if err = json.Unmarshal(*response.Body, balance); err != nil {
		return nil, response, err
	}

	return balance, response, nil
}
