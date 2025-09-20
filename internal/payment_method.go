package internal

// CreatePaymentMethodResponse returns a sample response for creating a payment method
func CreatePaymentMethodResponse() []byte {
	return []byte(`
{
  "createdAt": "2024-08-29T10:03:06.537Z",
  "updatedAt": "2024-08-29T10:03:06.537Z",
  "id": "68ceed244e5f2f5a69f72657",
  "metadata": {},
  "operatorDetails": {
    "operatorId": null,
    "merchantKey": "string"
  },
  "active": true,
  "type": "CARD",
  "walletTypeProductName": "string",
  "mobileMoneyDetails": {
    "countryCode": "string"
  },
  "neeroPersonDetails": {
    "personId": null,
    "accountId": "string",
    "country": "string",
    "paymentRequestId": null
  },
  "neeroMerchantDetails": {
    "merchantKey": "string",
    "storeId": "string",
    "balanceId": "string",
    "operatorId": null,
    "country": "string"
  },
  "paypalDetails": {
    "email": "string",
    "countryCode": "string"
  },
  "shortInfo": "string",
  "walletId": null
}`)
}
