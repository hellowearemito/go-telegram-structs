package telegram

// SendInvoice : Use this method to send invoices. On success, the sent Message is returned.
type SendInvoice struct {
	ChatID                    int64                 `json:"chat_id"`                       // Unique identifier for the target private chat
	Title                     string                `json:"title"`                         // Product name, 1-32 characters
	Description               string                `json:"description"`                   // Product description, 1-255 characters
	Payload                   string                `json:"payload"`                       // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string                `json:"provider_token"`                // Payments provider token, obtained via Botfather
	StartParameter            string                `json:"start_parameter"`               // Unique deep-linking parameter that can be used to generate this invoice when used as a start parameter
	Currency                  string                `json:"currency"`                      // Three-letter ISO 4217 currency code, see more on currencies
	Prices                    []LabeledPrice        `json:"prices"`                        // Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	ProviderData              *string               `json:"provider_data"`                 // JSON-encoded data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoURL                  *string               `json:"photo_url"`                     // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 *int64                `json:"photo_size"`                    // Photo size
	PhotoWidth                *int64                `json:"photo_width"`                   // Photo width
	PhotoHeight               *int64                `json:"photo_height"`                  // Photo height
	NeedName                  *bool                 `json:"need_name"`                     // Pass True, if you require the user's full name to complete the order
	NeedPhoneNumber           *bool                 `json:"need_phone_number"`             // Pass True, if you require the user's phone number to complete the order
	NeedEmail                 *bool                 `json:"need_email"`                    // Pass True, if you require the user's email address to complete the order
	NeedShippingAddress       *bool                 `json:"need_shipping_address"`         // Pass True, if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider *bool                 `json:"send_phone_number_to_provider"` // Pass True, if user's phone number should be sent to provider
	SendEmailToProvider       *bool                 `json:"send_email_to_provider"`        // Pass True, if user's email address should be sent to provider
	IsFlexible                *bool                 `json:"is_flexible"`                   // Pass True, if the final price depends on the shipping method
	DisableNotification       *bool                 `json:"disable_notification"`          // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID          *int64                `json:"reply_to_message_id"`           // 	If the message is a reply, ID of the original message
	ReplyMarkup               *InlineKeyboardMarkup `json:"reply_markup"`                  // A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
}

// AnswerShippingQuery : if you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
type AnswerShippingQuery struct {
	ShippingQueryID string           `json:"shipping_query_id"` // Unique identifier for the query to be answered
	Ok              bool             `json:"ok"`                // Specify True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	ShippingOptions []ShippingOption `json:"shipping_options"`  // Required if ok is True. A JSON-serialized array of available shipping options.
	ErrorMessage    *string          `json:"error_message"`     // Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
}

// AnswerPreCheckoutQuery : Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryID string  `json:"pre_checkout_query_id"` // Unique identifier for the query to be answered
	Ok                 bool    `json:"ok"`                    // Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	ErrorMessage       *string `json:"error_message"`         // Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
}
