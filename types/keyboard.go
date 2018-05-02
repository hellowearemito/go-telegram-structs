package types

import "github.com/hellowearemito/go-telegram-structs/games"

// ReplyMarkup represents reply markup: InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply.
type ReplyMarkup interface{}

// ReplyKeyboardMarkup represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`          // Array of button rows, each represented by an Array of KeyboardButton objects
	ResizeBoard     *bool              `json:"resize_board"`      // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard *bool              `json:"one_time_keyboard"` // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	Selective       *bool              `json:"selective"`         // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
}

// KeyboardButton represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields are mutually exclusive.
type KeyboardButton struct {
	Text            string `json:"text"`             // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestContact  *bool  `json:"request_contact"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestLocation *bool  `json:"request_location"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
}

// ReplyKeyboardRemove: Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"` // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      *bool `json:"selective"`       // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

// InlineKeyboardButton represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text                         string  `json:"text"`                             // Label text on the button
	URL                          *string `json:"url"`                              // Optional. HTTP url to be opened when button is pressed
	CallbackData                 *string `json:"callback_data"`                    // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery            *string `json:"switch_inline_query"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat"` // Optional. If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.
	CallbackGame                 games.CallbackGame `json:"callback_game"`                    // Optional. Description of the game that will be launched when the user presses the button. NOTE: This type of button must always be the first button in the first row.
	Pay *bool `json:"pay"` // Optional. Specify True, to send a Pay button. NOTE: This type of button must always be the first button in the first row.
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	ID              string  `json:"id"`                // Unique identifier for this query
	From            User    `json:"from"`              // Sender
	Message         Message `json:"message"`           // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	InlineMessageID string  `json:"inline_message_id"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string  `json:"chat_instance"`     // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            string  `json:"data"`              // Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
	GameShortName   string  `json:"game_short_name"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

// ForceReply: Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply bool  `json:"force_reply"` // Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	Selective  *bool `json:"selective"`   // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}
