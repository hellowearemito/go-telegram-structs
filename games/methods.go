package games

import "github.com/hellowearemito/go-telegram-structs/types"

// SendGame : Use this method to send a game. On success, the sent Message is returned.
type SendGame struct {
	ChatID              int64                       `json:"chat_id"`              // Unique identifier for the target chat
	GameShortName       string                      `json:"game_short_name"`      // Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather.
	DisableNotification *bool                       `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64                      `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *types.InlineKeyboardMarkup `json:"reply_markup"`         // A JSON-serialized object for an inline keyboard. If empty, one ‘Play game_title’ button will be shown. If not empty, the first button must launch the game.
}

// SetGameScore : Use this method to set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScore struct {
	UserID             int64   `json:"user_id"`              // User identifier
	Score              int64   `json:"score"`                // New score, must be non-negative
	Force              *bool   `json:"force"`                // Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEdigMessage *bool   `json:"disable_edig_message"` // Pass True, if the game message should not be automatically edited to include the current scoreboard
	ChatID             *int64  `json:"chat_id"`              // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageID          *int64  `json:"message_id"`           // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageID    *string `json:"inline_message_id"`    // Required if chat_id and message_id are not specified. Identifier of the inline message
}

// GetGameHighScores : Use this method to get data for high score tables. Will return the score of the specified user and several of his neighbors in a game. On success, returns an Array of GameHighScore objects.
type GetGameHighScores struct {
	UserID          int64   `json:"user_id"`           // Target user id
	ChatID          *int64  `json:"chat_id"`           // Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageID       *int64  `json:"message_id"`        // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageID *string `json:"inline_message_id"` // Required if chat_id and message_id are not specified. Identifier of the inline message
}
