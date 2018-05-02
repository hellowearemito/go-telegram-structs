package payments

import "github.com/hellowearemito/go-telegram-structs/types"

// LabeledPrice : This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label  string `json:"label"`  // Portion label
	Amount int64  `json:"amount"` // Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

// Invoice : This object contains basic information about an invoice.
type Invoice struct {
	Title          string `json:"title"`           // Product name
	Description    string `json:"description"`     // Product description
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter that can be used to generate this invoice
	Currency       string `json:"currency"`        // Three-letter ISO 4217 currency code
	TotalAmount    int64  `json:"total_amount"`    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

// ShippingAddress : This object represents a shipping address.
type ShippingAddress struct {
	CountryCode string `json:"country_code"`  // ISO 3166-1 alpha-2 country code
	State       string `json:"state"`         // State, if applicable
	City        string `json:"city"`          // City
	StreetLine1 string `json:"street_line_1"` // First line for the address
	StreetLine2 string `json:"street_line_2"` // Second line for the address
	PostCode    string `json:"post_code"`     // Address post code
}

// OrderInfo : This object represents information about an order.
type OrderInfo struct {
	Name            *string          `json:"name"`             // Optional. User name
	PhoneNumber     *string          `json:"phone_number"`     // Optional. User's phone number
	Email           *string          `json:"email"`            // Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address"` // Optional. User shipping address
}

// ShippingOption : This object represents one shipping option.
type ShippingOption struct {
	ID     string         `json:"id"`     // Shipping option identifier
	Title  *string        `json:"title"`  // Option title
	Prices []LabeledPrice `json:"prices"` // List of price portions
}

// SuccessfulPayment : This object contains basic information about a successful payment.
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`                   // Three-letter ISO 4217 currency code
	TotalAmount             int64      `json:"total_amount"`               // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string     `json:"invoice_payload"`            // Bot specified invoice payload
	ShippingOptionID        *string    `json:"shipping_option_id"`         // Optional. Identifier of the shipping option chosen by the user
	OrderInfo               *OrderInfo `json:"order_info"`                 // Optional. Order info provided by the user
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"` // Telegram payment identifier
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"` // Provider payment identifier
}

// ShippingQuery : This object contains information about an incoming shipping query.
type ShippingQuery struct {
	ID              string          `json:"id"`               // Unique query identifier
	From            types.User      `json:"from"`             // User who sent the query
	InvoicePayload  string          `json:"invoice_payload"`  // Bot specified invoice payload
	ShippingAddress ShippingAddress `json:"shipping_address"` // User specified shipping address
}

// PreCheckoutQuery : This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	ID               string     `json:"id"`                 // Unique query identifier
	From             types.User `json:"from"`               // User who sent the query
	Currency         string     `json:"currency"`           // Three-letter ISO 4217 currency code
	TotalAmount      int64      `json:"total_amount"`       // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload   string     `json:"invoice_payload"`    // Bot specified invoice payload
	ShippingOptionID *string    `json:"shipping_option_id"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo        *OrderInfo `json:"order_info"`         // Optional. Order info provided by the user
}
