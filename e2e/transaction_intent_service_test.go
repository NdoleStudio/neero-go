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

func TestTransactionIntentService_CashIn(t *testing.T) {
	// Arrange
	source := os.Getenv("NEERO_SOURCE_PAYMENT_ID")
	destination := os.Getenv("NEERO_DESTINATION_PAYMENT_ID")
	request := &neero.TransactionIntentCreateRequest{
		Amount:                     10000,
		CurrencyCode:               "XAF",
		PaymentType:                neero.TransactionIntentPaymentTypeMerchantCollection,
		SourcePaymentMethodID:      &source,
		DestinationPaymentMethodID: &destination,
		Confirm:                    true,
		MetaData:                   map[string]any{"order_id": "1234"},
	}

	// Act
	transaction, response, err := client.TransactionIntent.CashIn(context.Background(), request)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.NotEmpty(t, transaction.ID)
	spew.Dump(transaction.ID)
}

func TestTransactionIntentService_Get(t *testing.T) {
	// Arrange
	transactionID := "68cf03354e5f2f5a69f72773"

	// Act
	transaction, response, err := client.TransactionIntent.Get(context.Background(), transactionID)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, transactionID, transaction.ID)
	assert.Equal(t, "SUCCESSFUL", transaction.Status)
	spew.Dump(transaction.ID)
}

func TestTransactionIntentService_CashOut(t *testing.T) {
	// Arrange
	source := os.Getenv("NEERO_DESTINATION_PAYMENT_ID")
	destination := "68cf045e0aac937c7e5140b1"
	request := &neero.TransactionIntentCreateRequest{
		Amount:                     1000,
		CurrencyCode:               "XAF",
		PaymentType:                neero.TransactionIntentPaymentTypeOrangeMoneyTransfer,
		SourcePaymentMethodID:      &source,
		DestinationPaymentMethodID: &destination,
		Confirm:                    true,
		MetaData:                   map[string]any{"order_id": "1234"},
	}

	// Act
	transaction, response, err := client.TransactionIntent.CashOut(context.Background(), request)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.NotEmpty(t, transaction.ID)
	spew.Dump(transaction.ID)
}
