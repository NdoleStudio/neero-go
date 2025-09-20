package neero

// TransactionIntentPaymentType represents the type of payment for a transaction intent.
type TransactionIntentPaymentType string

const (
	// TransactionIntentPaymentTypeMerchantCollection represents a merchant collection payment type.
	TransactionIntentPaymentTypeMerchantCollection TransactionIntentPaymentType = "MERCHANT_COLLECTION"
	// TransactionIntentPaymentTypeOrangeMoneyTransfer represents an orange money transfer payment type.
	TransactionIntentPaymentTypeOrangeMoneyTransfer TransactionIntentPaymentType = "ORANGE_MONEY_TRANSFER"
	// TransactionIntentPaymentTypeMTNMobileMoneyTransfer represents an MTN mobile money transfer payment type.
	TransactionIntentPaymentTypeMTNMobileMoneyTransfer TransactionIntentPaymentType = "MTN_MONEY_TRANSFER"
)

// TransactionIntentCreateRequest represents a request to create a transaction intent.
type TransactionIntentCreateRequest struct {
	Amount                     int                          `json:"amount"`
	CurrencyCode               string                       `json:"currencyCode"`
	PaymentType                TransactionIntentPaymentType `json:"paymentType"`
	SourcePaymentMethodID      *string                      `json:"sourcePaymentMethodId,omitempty"`
	DestinationPaymentMethodID *string                      `json:"destinationPaymentMethodId,omitempty"`
	Confirm                    bool                         `json:"confirm"`
	MetaData                   map[string]any               `json:"metadata"`
}

// TransactionIntentTransaction represents the transaction intent details.
type TransactionIntentTransaction struct {
	CreatedAt       string         `json:"createdAt"`
	UpdatedAt       string         `json:"updatedAt"`
	ID              string         `json:"id"`
	Metadata        map[string]any `json:"metadata"`
	OperatorDetails struct {
		OperatorID  any    `json:"operatorId"`
		MerchantKey string `json:"merchantKey"`
	} `json:"operatorDetails"`
	StatusUpdates []struct {
		Status    string `json:"status"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"statusUpdates"`
	SourcePaymentMethodID      string `json:"sourcePaymentMethodId"`
	SourcePaymentMethodDetails struct {
		ID                    string `json:"id"`
		Type                  string `json:"type"`
		WalletTypeProductName string `json:"walletTypeProductName"`
		ShortInfo             string `json:"shortInfo"`
		WalletID              any    `json:"walletId"`
		MerchantDetails       struct {
			MerchantKey string `json:"merchantKey"`
			StoreID     string `json:"storeId"`
			BalanceID   string `json:"balanceId"`
			OperatorID  any    `json:"operatorId"`
			Country     string `json:"country"`
		} `json:"merchantDetails"`
	} `json:"sourcePaymentMethodDetails"`
	DestinationPaymentMethodID      string `json:"destinationPaymentMethodId"`
	DestinationPaymentMethodDetails struct {
		ID                    string `json:"id"`
		Type                  string `json:"type"`
		WalletTypeProductName string `json:"walletTypeProductName"`
		ShortInfo             string `json:"shortInfo"`
		WalletID              any    `json:"walletId"`
		MerchantDetails       struct {
			MerchantKey string `json:"merchantKey"`
			StoreID     string `json:"storeId"`
			BalanceID   string `json:"balanceId"`
			OperatorID  any    `json:"operatorId"`
			Country     string `json:"country"`
		} `json:"merchantDetails"`
	} `json:"destinationPaymentMethodDetails"`
	Type               string `json:"type"`
	Amount             any    `json:"amount"`
	Currency           string `json:"currency"`
	PaymentType        string `json:"paymentType"`
	ExpirationDateTime string `json:"expirationDateTime"`
	PaymentToken       string `json:"paymentToken"`
	SuccessURL         string `json:"successUrl"`
	FailureURL         string `json:"failureUrl"`
	CancelURL          string `json:"cancelUrl"`
	Customer           struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	} `json:"customer"`
	ConfirmationSessionID string `json:"confirmationSessionId"`
	NextAction            struct {
		Type                  string `json:"type"`
		CreatedDateTime       string `json:"createdDateTime"`
		NeeroPersonWithdrawal struct {
			WithdrawalRequestID string `json:"withdrawalRequestId"`
		} `json:"neeroPersonWithdrawal"`
		URLRedirect struct {
			URLToRedirect string `json:"urlToRedirect"`
		} `json:"urlRedirect"`
	} `json:"nextAction"`
	CollectCustomerDetails        bool   `json:"collectCustomerDetails"`
	Status                        string `json:"status"`
	CreateFlowTransactionRequests []struct {
		Amount                     int    `json:"amount"`
		CurrencyCode               string `json:"currencyCode"`
		PaymentType                string `json:"paymentType"`
		SourcePaymentMethodID      string `json:"sourcePaymentMethodId"`
		DestinationPaymentMethodID string `json:"destinationPaymentMethodId"`
		TransactionType            string `json:"transactionType"`
	} `json:"createFlowTransactionRequests"`
	Fees struct {
		SourceFee struct {
			Owner struct {
				PersonID           any            `json:"personId"`
				WalletID           any            `json:"walletId"`
				MerchantKey        string         `json:"merchantKey"`
				StoreID            string         `json:"storeId"`
				BalanceID          string         `json:"balanceId"`
				ServiceNumber      string         `json:"serviceNumber"`
				CountryCode        string         `json:"countryCode"`
				Type               string         `json:"type"`
				GatewayCode        string         `json:"gatewayCode"`
				AdditionalData     map[string]any `json:"additionalData"`
				MaxPositiveBalance any            `json:"maxPositiveBalance"`
				Balance            any            `json:"balance"`
				Plan               string         `json:"plan"`
				PlanOrder          any            `json:"planOrder"`
			} `json:"owner"`
			Amount                       any `json:"amount"`
			Margin                       any `json:"margin"`
			OnAmountVoucherAmount        any `json:"onAmountVoucherAmount"`
			FreeLimitAmount              any `json:"freeLimitAmount"`
			OnFeeAmountVoucherAmount     any `json:"onFeeAmountVoucherAmount"`
			VoucherAmount                any `json:"voucherAmount"`
			ProductFeeAmount             any `json:"productFeeAmount"`
			WalletFeeAmount              any `json:"walletFeeAmount"`
			GovernmentFeeAmount          any `json:"governmentFeeAmount"`
			DenominationAdjustmentAmount any `json:"denominationAjustementAmount"`
			FeeVATAmount                 any `json:"feeVATAmount"`
			Vat                          any `json:"vat"`
			Rate                         any `json:"rate"`
			MarginRate                   any `json:"marginRate"`
			Voucher                      struct {
				Code  string `json:"code"`
				Value any    `json:"value"`
				Type  string `json:"type"`
			} `json:"voucher"`
			GovernmentFee struct {
				FeeCategory string `json:"feeCategory"`
				Fee         any    `json:"fee"`
				FeeType     string `json:"feeType"`
			} `json:"governmentFee"`
			ProductFee struct {
				FeeCategory string `json:"feeCategory"`
				Fee         any    `json:"fee"`
				FeeType     string `json:"feeType"`
			} `json:"productFee"`
			WalletTypefee struct {
				FeeCategory string `json:"feeCategory"`
				Fee         any    `json:"fee"`
				FeeType     string `json:"feeType"`
			} `json:"walletTypefee"`
			OtherPlansFees []string `json:"otherPlansFees"`
			Target         string   `json:"target"`
			Currency       string   `json:"currency"`
			MargedAmount   any      `json:"margedAmount"`
			TotalFeeAmount any      `json:"totalFeeAmount"`
			TotalAmount    any      `json:"totalAmount"`
			BillableAmount any      `json:"billableAmount"`
			Conversion     struct {
				MargedAmount        any    `json:"margedAmount"`
				TotalFeeAmount      any    `json:"totalFeeAmount"`
				TotalAmount         any    `json:"totalAmount"`
				BillableAmount      any    `json:"billableAmount"`
				ProductFeeAmount    any    `json:"productFeeAmount"`
				WalletFeeAmount     any    `json:"walletFeeAmount"`
				GovernmentFeeAmount any    `json:"governmentFeeAmount"`
				Rate                any    `json:"rate"`
				Currency            string `json:"currency"`
			} `json:"conversion"`
		} `json:"sourceFee"`
		DestinationFee struct {
			Owner struct {
				PersonID           any            `json:"personId"`
				WalletID           any            `json:"walletId"`
				MerchantKey        string         `json:"merchantKey"`
				StoreID            string         `json:"storeId"`
				BalanceID          string         `json:"balanceId"`
				ServiceNumber      string         `json:"serviceNumber"`
				CountryCode        string         `json:"countryCode"`
				Type               string         `json:"type"`
				GatewayCode        string         `json:"gatewayCode"`
				AdditionalData     map[string]any `json:"additionalData"`
				MaxPositiveBalance any            `json:"maxPositiveBalance"`
				Balance            any            `json:"balance"`
				Plan               string         `json:"plan"`
				PlanOrder          any            `json:"planOrder"`
			} `json:"owner"`
			Amount                       any `json:"amount"`
			Margin                       any `json:"margin"`
			OnAmountVoucherAmount        any `json:"onAmountVoucherAmount"`
			FreeLimitAmount              any `json:"freeLimitAmount"`
			OnFeeAmountVoucherAmount     any `json:"onFeeAmountVoucherAmount"`
			VoucherAmount                any `json:"voucherAmount"`
			ProductFeeAmount             any `json:"productFeeAmount"`
			WalletFeeAmount              any `json:"walletFeeAmount"`
			GovernmentFeeAmount          any `json:"governmentFeeAmount"`
			DenominationAjustementAmount any `json:"denominationAjustementAmount"`
			FeeVATAmount                 any `json:"feeVATAmount"`
			Vat                          any `json:"vat"`
			Rate                         any `json:"rate"`
			MarginRate                   any `json:"marginRate"`
			Voucher                      struct {
				Code  string `json:"code"`
				Value any    `json:"value"`
				Type  string `json:"type"`
			} `json:"voucher"`
			GovernmentFee struct {
				FeeCategory string `json:"feeCategory"`
				Fee         any    `json:"fee"`
				FeeType     string `json:"feeType"`
			} `json:"governmentFee"`
			ProductFee struct {
				FeeCategory string `json:"feeCategory"`
				Fee         any    `json:"fee"`
				FeeType     string `json:"feeType"`
			} `json:"productFee"`
			WalletTypefee struct {
				FeeCategory string `json:"feeCategory"`
				Fee         any    `json:"fee"`
				FeeType     string `json:"feeType"`
			} `json:"walletTypefee"`
			OtherPlansFees []string `json:"otherPlansFees"`
			Target         string   `json:"target"`
			Currency       string   `json:"currency"`
			MargedAmount   any      `json:"margedAmount"`
			TotalFeeAmount any      `json:"totalFeeAmount"`
			TotalAmount    any      `json:"totalAmount"`
			BillableAmount any      `json:"billableAmount"`
			Conversion     struct {
				MargedAmount        any    `json:"margedAmount"`
				TotalFeeAmount      any    `json:"totalFeeAmount"`
				TotalAmount         any    `json:"totalAmount"`
				BillableAmount      any    `json:"billableAmount"`
				ProductFeeAmount    any    `json:"productFeeAmount"`
				WalletFeeAmount     any    `json:"walletFeeAmount"`
				GovernmentFeeAmount any    `json:"governmentFeeAmount"`
				Rate                any    `json:"rate"`
				Currency            string `json:"currency"`
			} `json:"conversion"`
		} `json:"destinationFee"`
	} `json:"fees"`
}
