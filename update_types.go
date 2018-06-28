package telegram

// Update represents an incoming update. At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateID           int64               `json:"update_id"`            // The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message            *Message            `json:"message"`              // Optional. New incoming message of any kind — text, photo, sticker, etc.
	EditedMessage      *Message            `json:"edited_message"`       // Optional. New version of a message that is known to the bot and was edited
	ChannelPost        *Message            `json:"channel_post"`         // Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	EditedChannelPost  *Message            `json:"edited_channel_post"`  // Optional. New version of a channel post that is known to the bot and was edited
	InlineQuery        *Query              `json:"inline_query"`         // Optional. New incoming inline query
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"` // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery      *CallbackQuery      `json:"callback_query"`       // Optional. New incoming callback query
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`       // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`   // Optional. New incoming pre-checkout query. Contains full information about checkout
}

// WebhookInfo represents the webhook's information.
type WebhookInfo struct {
	URL                  string   `json:"url"`                    // Webhook URL, may be empty if webhook is not set up
	HasCustomCertificate bool     `json:"has_custom_certificate"` // True, if a custom certificate was provided for webhook certificate checks
	PendingUpdateCount   int64    `json:"pending_update_count"`   // Number of updates awaiting delivery
	LastErrorDate        *int64   `json:"last_error_date"`        // Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage     *string  `json:"last_error_message"`     // Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	MaxConnections       *int64   `json:"max_connections"`        // Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	AllowedUpdates       []string `json:"allowed_updates"`        // Optional. A list of update types the bot is subscribed to. Defaults to all update types
}
