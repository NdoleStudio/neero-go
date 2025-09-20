package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/neero-go"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestPaymentMethodService_Create(t *testing.T) {
	// Arrange
	request := &neero.CreatePaymentMethodRequestMobileMoneyDetails{
		PhoneNumber:         "+237691111111",
		CountryIso:          "CM",
		MobileMoneyProvider: neero.MobileMoneyProviderOrangeMoney,
	}

	// Act
	paymentMethod, response, err := client.PaymentMethod.Create(context.Background(), request)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.NotEmpty(t, paymentMethod.ID)
	spew.Dump(paymentMethod.ID)
}

func TestPaymentMethodService_ResolveDetails(t *testing.T) {
	// Arrange
	request := &neero.CreatePaymentMethodRequestMobileMoneyDetails{
		PhoneNumber:         "+237691111111",
		CountryIso:          "CM",
		MobileMoneyProvider: neero.MobileMoneyProviderOrangeMoney,
	}

	// Act
	paymentMethod, response, err := client.PaymentMethod.ResolveDetails(context.Background(), request)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.NotEmpty(t, paymentMethod.Name)
	spew.Dump(paymentMethod.Name)
}
