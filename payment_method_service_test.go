package neero

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/neero-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestPaymentMethodService_Create(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := internal.MakeTestServer(http.StatusOK, internal.CreatePaymentMethodResponse())
	client := New(WithBaseURL(server.URL))
	request := &CreatePaymentMethodRequestMobileMoneyDetails{
		PhoneNumber:         "+237678877898",
		CountryIso:          "CM",
		MobileMoneyProvider: MobileMoneyProviderMTNMoney,
	}

	// Act
	paymentMethod, response, err := client.PaymentMethod.Create(context.Background(), request)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, "68ceed244e5f2f5a69f72657", paymentMethod.ID)

	// Teardown
	server.Close()
}
