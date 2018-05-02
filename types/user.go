package types

// User represents a Telegram user or bot.
type User struct {
	ID           int64   `json:"id"`            // Unique identifier for this user or bot
	IsBot        bool    `json:"is_bot"`        // True, if this user is a bot
	FirstName    string  `json:"first_name"`    // User‘s or bot’s first name
	LastName     *string `json:"last_name"`     // Optional. User‘s or bot’s last name
	Username     *string `json:"username"`      // Optional. User‘s or bot’s username
	LanguageCode *string `json:"language_code"` // Optional. IETF language tag of the user's language
}
