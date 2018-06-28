package telegram

// SendSticker : Use this method to send .webp stickers. On success, the sent Message is returned.
type SendSticker struct {
	ChatID              string       `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Sticker             interface{}  `json:"sticker"`              // InputFile or String. Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data.
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// GetStickerSet : Use this method to get a sticker set. On success, a StickerSet object is returned.
type GetStickerSet struct {
	Name string `json:"name"` // Name of the sticker set
}

// UploadStickerFile : Use this method to upload a .png file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
type UploadStickerFile struct {
	UserID     int64     `json:"user_id"`     // User identifier of sticker file owner
	PngSticker InputFile `json:"png_sticker"` // Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px.
}

// CreateNewStickerSet : Use this method to create new sticker set owned by a user. The bot will be able to edit the created sticker set. Returns True on success.
type CreateNewStickerSet struct {
	UserID        int64         `json:"user_id"`        // User identifier of created sticker set owner
	Name          string        `json:"name"`           // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
	Title         string        `json:"title"`          // Sticker set title, 1-64 characters
	PngSticker    InputFile     `json:"png_sticker"`    // InputFile or String. Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	Emojis        string        `json:"emojis"`         // One or more emoji corresponding to the sticker
	ContainsMasks *bool         `json:"contains_masks"` // Pass True, if a set of mask stickers should be created
	MaskPosition  *MaskPosition `json:"mask_position"`  // A JSON-serialized object for position where the mask should be placed on faces
}

// AddStickerToSet : Use this method to add a new sticker to a set created by the bot. Returns True on success.
type AddStickerToSet struct {
	UserID       int64         `json:"user_id"`     // User identifier of sticker set owner
	Name         string        `json:"name"`        // Sticker set name
	PngSticker   InputFile     `json:"png_sticker"` // InputFile or String. Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	Emojis       string        `json:"emojis"`      // One or more emoji corresponding to the sticker
	MaskPosition *MaskPosition // A JSON-serialized object for position where the mask should be placed on faces
}

// SetStickerPositionInSet : Use this method to move a sticker in a set created by the bot to a specific position . Returns True on success.
type SetStickerPositionInSet struct {
	Sticker  string `json:"sticker"`  // File identifier of the sticker
	Position int64  `json:"position"` // New sticker position in the set, zero-based
}

// DeleteStickerFromSet : Use this method to delete a sticker from a set created by the bot. Returns True on success.
type DeleteStickerFromSet struct {
	Sticker string `json:"sticker"` // File identifier of the sticker
}
