package updates

import (
	"github.com/hellowearemito/go-telegram-structs/types"
	"github.com/hellowearemito/go-telegram-structs/inline"
	"github.com/hellowearemito/go-telegram-structs/payments"
)

// Update represents an incoming update. At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateID           int64               `json:"update_id"`            // The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message            *types.Message       `json:"message"`              // Optional. New incoming message of any kind — text, photo, sticker, etc.
	EditedMessage      *types.Message       `json:"edited_message"`       // Optional. New version of a message that is known to the bot and was edited
	ChannelPost        *types.Message       `json:"channel_post"`         // Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	EditedChannelPost  *types.Message       `json:"edited_channel_post"`  // Optional. New version of a channel post that is known to the bot and was edited
	InlineQuery        *inline.Query         `json:"inline_query"`         // Optional. New incoming inline query
	ChosenInlineResult *inline.ChosenInlineResult  `json:"chosen_inline_result"` // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery      *types.CallbackQuery `json:"callback_query"`       // Optional. New incoming callback query
	ShippingQuery      *payments.ShippingQuery       `json:"shipping_query"`       // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery   *payments.PreCheckoutQuery    `json:"pre_checkout_query"`   // Optional. New incoming pre-checkout query. Contains full information about checkout
}
