package neero

import (
	"context"
	"encoding/json"
	"net/http"
)

// paymentMethodService is the API client for the `/payment-methods` endpoint
type paymentMethodService service

// Create a Payment Method in the system.
//
// API Docs: https://app.theneo.io/neero/neero-payment-gateway/payment-methods/create-payment-method
func (service *paymentMethodService) Create(ctx context.Context, createRequest CreatePaymentMethodRequest) (*CreatePaymentMethodResponse, *Response, error) {
	payload := map[string]any{
		"type":                            createRequest.Type(),
		string(createRequest.Attribute()): createRequest,
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/payment-gateway/api/v1/payment-methods", payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	paymentMethod := new(CreatePaymentMethodResponse)
	if err = json.Unmarshal(*response.Body, paymentMethod); err != nil {
		return nil, response, err
	}

	return paymentMethod, response, nil
}

// ResolveDetails returns the details of Payment Method from Provider
//
// API Docs: https://app.theneo.io/neero/neero-payment-gateway/payment-methods/resolve-payment-method-details-1
func (service *paymentMethodService) ResolveDetails(ctx context.Context, createRequest CreatePaymentMethodRequest) (*ResolvePaymentMethodDetailsResponse, *Response, error) {
	payload := map[string]any{
		"type":                            createRequest.Type(),
		string(createRequest.Attribute()): createRequest,
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/payment-gateway/api/v1/payment-methods/resolve-details", payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	details := new(ResolvePaymentMethodDetailsResponse)
	if err = json.Unmarshal(*response.Body, details); err != nil {
		return nil, response, err
	}

	return details, response, nil
}
