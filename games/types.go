package games

import "github.com/hellowearemito/go-telegram-structs/types"

// Game : This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	Title        string                `json:"title"`         // Title of the game
	Description  string                `json:"description"`   // Description of the game
	Photo        []types.PhotoSize     `json:"photo"`         // Photo that will be displayed in the game message in chats.
	Text         string                `json:"text"`          // Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	TextEntities []types.MessageEntity `json:"text_entities"` // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Animation    *Animation            `json:"animation"`     // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
}

// Animation : You can provide an animation for your game so that it looks stylish in chats (check out Lumberjack for an example). This object represents an animation file to be displayed in the message containing a game.
type Animation struct {
	FileID   string           `json:"file_id"`   // Unique file identifier
	Thumb    *types.PhotoSize `json:"thumb"`     // Optional. Animation thumbnail as defined by sender
	FileName *string          `json:"file_name"` // Optional. Original animation filename as defined by sender
	MimeType *string          `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	FileSize *int64           `json:"file_size"` // Optional. File size
}

// CallbackGame : A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct{}

// GameHighScore : This object represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int64      `json:"position"` // Position in high score table for the game
	User     types.User `json:"user"`     // User
	Score    int64      `json:"score"`    // Score
}
