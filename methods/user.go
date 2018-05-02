package methods

import "github.com/hellowearemito/go-telegram-structs/types"

// GetMe : A simple method for testing your bot's auth token. Requires no parameters. Returns basic information about the bot in form of a User object.
type GetMe types.User

// GetUserProfilePhotos
type GetUserProfilePhotos struct {
	UserID int64 `json:"user_id"` // Unique identifier of the target user
	Offset *int64 `json:"offset"` // Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit *int64 `json:"limit"` // Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
}

