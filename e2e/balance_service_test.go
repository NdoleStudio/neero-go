package e2e

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/NdoleStudio/neero-go"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestBalanceService_Get(t *testing.T) {
	// Arrange
	request := &neero.CreatePaymentMethodRequestNeroMerchantDetails{
		MerchantKey: os.Getenv("NEERO_MERCHANT_KEY"),
		StoreID:     os.Getenv("NEERO_STORE_ID"),
		BalanceID:   os.Getenv("NEERO_BALANCE_ID"),
		OperatorID:  os.Getenv("NEERO_OPERATOR_ID"),
	}

	paymentMethod, _, _ := client.PaymentMethod.Create(context.Background(), request)
	spew.Dump(paymentMethod.ID)

	// Act
	balance, response, err := client.Balance.Get(context.Background(), paymentMethod.ID)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	spew.Dump(balance)
}
