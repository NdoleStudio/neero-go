package neero

// PaymentMethodType represents the type of payment method.
type PaymentMethodType string

const (
	// PaymentMethodTypeMobileMoney represents the mobile money payment method type.
	PaymentMethodTypeMobileMoney = PaymentMethodType("MOBILE_MONEY")

	// PaymentMethodTypeNeroMerchant represents the neero merchant payment method type.
	PaymentMethodTypeNeroMerchant = PaymentMethodType("NEERO_MERCHANT")
)

// MobileMoneyProvider represents the mobile money provider.
type MobileMoneyProvider string

const (
	// MobileMoneyProviderMTNMoney represents the MTN mobile money provider.
	MobileMoneyProviderMTNMoney = MobileMoneyProvider("MTN_MONEY")

	// MobileMoneyProviderOrangeMoney represents the orange money provider.
	MobileMoneyProviderOrangeMoney = MobileMoneyProvider("ORANGE_MONEY")
)

// PaymentMethodAttribute represents the attribute of the payment method.
type PaymentMethodAttribute string

const (
	// paymentMethodAttributeMobileMoney represents the mobile money payment method attribute.
	paymentMethodAttributeMobileMoney = PaymentMethodAttribute("mobileMoneyDetails")
)

// CreatePaymentMethodRequest represents a request to create a payment method.
type CreatePaymentMethodRequest interface {
	// Type returns the payment method type.
	Type() PaymentMethodType

	// Attribute returns the payment method attribute.
	Attribute() PaymentMethodAttribute
}

// CreatePaymentMethodResponse represents the response from creating a payment method.
type CreatePaymentMethodResponse struct {
	CreatedAt             string                                           `json:"createdAt"`
	UpdatedAt             string                                           `json:"updatedAt"`
	ID                    string                                           `json:"id"`
	Metadata              any                                              `json:"metadata"`
	OperatorDetails       *CreatePaymentMethodResponseOperatorDetails      `json:"operatorDetails"`
	Active                bool                                             `json:"active"`
	Type                  string                                           `json:"type"`
	WalletTypeProductName string                                           `json:"walletTypeProductName"`
	MobileMoneyDetails    *CreatePaymentMethodResponseMobileMoneyDetails   `json:"mobileMoneyDetails"`
	NeeroPersonDetails    *CreatePaymentMethodResponseNeeroPersonDetails   `json:"neeroPersonDetails"`
	NeeroMerchantDetails  *CreatePaymentMethodResponseNeeroMerchantDetails `json:"neeroMerchantDetails"`
	PaypalDetails         *CreatePaymentMethodResponsePaypalDetails        `json:"paypalDetails"`
	ShortInfo             string                                           `json:"shortInfo"`
	WalletID              any                                              `json:"walletId"`
}

// ResolvePaymentMethodDetailsResponse represents the response from resolving payment method details.
type ResolvePaymentMethodDetailsResponse struct {
	Name string `json:"name"`
}

// CreatePaymentMethodResponseOperatorDetails represents the operator details in the payment method response.
type CreatePaymentMethodResponseOperatorDetails struct {
	OperatorID  any    `json:"operatorId"`
	MerchantKey string `json:"merchantKey"`
}

// CreatePaymentMethodResponseMobileMoneyDetails represents the mobile money details in the payment method response.
type CreatePaymentMethodResponseMobileMoneyDetails struct {
	CountryCode string `json:"countryCode"`
}

// CreatePaymentMethodResponseNeeroPersonDetails represents the neero person details in the payment method response.
type CreatePaymentMethodResponseNeeroPersonDetails struct {
	PersonID         any    `json:"personId"`
	AccountID        string `json:"accountId"`
	Country          string `json:"country"`
	PaymentRequestID any    `json:"paymentRequestId"`
}

// CreatePaymentMethodResponseNeeroMerchantDetails represents the neero merchant details in the payment method response.
type CreatePaymentMethodResponseNeeroMerchantDetails struct {
	MerchantKey string `json:"merchantKey"`
	StoreID     string `json:"storeId"`
	BalanceID   string `json:"balanceId"`
	OperatorID  any    `json:"operatorId"`
	Country     string `json:"country"`
}

// CreatePaymentMethodResponsePaypalDetails represents the PayPal details in the payment method response.
type CreatePaymentMethodResponsePaypalDetails struct {
	Email       string `json:"email"`
	CountryCode string `json:"countryCode"`
}

// CreatePaymentMethodRequestMobileMoneyDetails represents the details for creating a mobile money payment method.
type CreatePaymentMethodRequestMobileMoneyDetails struct {
	PhoneNumber         string              `json:"phoneNumber"`
	CountryIso          string              `json:"countryIso"`
	MobileMoneyProvider MobileMoneyProvider `json:"mobileMoneyProvider"`
}

// Type returns the payment method type.
func (c *CreatePaymentMethodRequestMobileMoneyDetails) Type() PaymentMethodType {
	return PaymentMethodTypeMobileMoney
}

// Attribute returns the payment method attribute.
func (c *CreatePaymentMethodRequestMobileMoneyDetails) Attribute() PaymentMethodAttribute {
	return paymentMethodAttributeMobileMoney
}

// CreatePaymentMethodRequestNeroMerchantDetails represents the details for creating a neero merchant payment method.
type CreatePaymentMethodRequestNeroMerchantDetails struct {
	MerchantKey string `json:"merchantKey"`
	StoreID     string `json:"storeId"`
	BalanceID   string `json:"balanceId"`
	OperatorID  string `json:"operatorId"`
}

// Type returns the payment method type.
func (c *CreatePaymentMethodRequestNeroMerchantDetails) Type() PaymentMethodType {
	return PaymentMethodTypeNeroMerchant
}

// Attribute returns the payment method attribute.
func (c *CreatePaymentMethodRequestNeroMerchantDetails) Attribute() PaymentMethodAttribute {
	return "neeroMerchantDetails"
}
