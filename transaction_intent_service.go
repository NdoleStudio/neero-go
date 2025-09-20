package neero

import (
	"context"
	"encoding/json"
	"net/http"
)

// transactionIntentService is the API client for the `/payment-gateway/api/v1/transaction-intents` endpoint
type transactionIntentService service

// CashIn creates a transaction in which the destination is a Merchant Account you own.
//
// API Docs: https://app.theneo.io/neero/neero-payment-gateway/transaction-intents-1/create-cash-in-payment-intent
func (service *transactionIntentService) CashIn(ctx context.Context, payload *TransactionIntentCreateRequest) (*TransactionIntentTransaction, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/payment-gateway/api/v1/transaction-intents/cash-in", payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(TransactionIntentTransaction)
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}

// CashOut creates a transaction in which the source is your Merchant Account and the destination can be a wallet, bank account, credit card
//
// API Docs: https://app.theneo.io/neero/neero-payment-gateway/transaction-intents-1/create-cash-out-payment-intent
func (service *transactionIntentService) CashOut(ctx context.Context, payload *TransactionIntentCreateRequest) (*TransactionIntentTransaction, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/payment-gateway/api/v1/transaction-intents/cash-out", payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(TransactionIntentTransaction)
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}

// Get a transaction intent by ID
//
// API Docs: https://app.theneo.io/neero/neero-payment-gateway/transaction-intents-1/find-transaction-intent-by-id
func (service *transactionIntentService) Get(ctx context.Context, transactionIntentID string) (*TransactionIntentTransaction, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/payment-gateway/api/v1/transaction-intents/"+transactionIntentID, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	balance := new(TransactionIntentTransaction)
	if err = json.Unmarshal(*response.Body, balance); err != nil {
		return nil, response, err
	}

	return balance, response, nil
}
