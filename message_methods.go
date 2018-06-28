package telegram

// SendMessage : Use this method to send text messages. On success, the sent Message is returned.
type SendMessage struct {
	ChatID                string       `json:"chat_id"`                  // String or Integer. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text                  string       `json:"text"`                     // Text of the message to be sent
	ParseMode             *string      `json:"parse_mode"`               // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview *bool        `json:"disable_web_page_preview"` // Disables link previews for links in this message
	DisableNotification   *bool        `json:"disable_notification"`     // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID      *int64       `json:"reply_to_message_id"`      // If the message is a reply, ID of the original message
	ReplyMarkup           *ReplyMarkup `json:"reply_markup"`             // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// ForwardMessage : Use this method to forward messages of any kind. On success, the sent Message is returned.
type ForwardMessage struct {
	ChatID              string `json:"chat_id"`              // String or Integer. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	FromChatID          string `json:"from_chat_id"`         // String or Integer. Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification *bool  `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	MessageID           int64  `json:"message_id"`           // Message identifier in the chat specified in from_chat_id
}

// SendPhoto : Use this method to send photos. On success, the sent Message is returned.
type SendPhoto struct {
	ChatID              string       `json:"chat_id"`              // String or Integer. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo               InputFile    `json:"photo"`                // InputFile or String. Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data.
	Caption             *string      `json:"caption"`              // Photo caption (may also be used when resending photos by file_id), 0-200 characters
	ParseMode           *string      `json:"parse_mode"`           // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendAudio : Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future. For sending voice messages, use the sendVoice method instead.
type SendAudio struct {
	ChatID              string       `json:"chat_id"`              // String or Integer. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Audio               InputFile    `json:"audio"`                // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data.
	Caption             *string      `json:"caption"`              // Audio caption, 0-200 characters
	ParseMode           *string      `json:"parse_mode"`           // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Duration            *string      `json:"duration"`             // Duration of the audio in seconds
	Performer           *string      `json:"performer"`            // Performer
	Title               *string      `json:"title"`                // Track name
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendDocument : Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocument struct {
	ChatID              string       `json:"chat_id"`              // String or Integer. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Document            InputFile    `json:"document"`             // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	Caption             *string      `json:"caption"`              // Document caption (may also be used when resending documents by file_id), 0-200 characters
	ParseMode           *string      `json:"parse_mode"`           // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendVideo : Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideo struct {
	ChatID              string       `json:"chat_id"`              // String or Integer. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Video               InputFile    `json:"video"`                // InputFile or String. Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data.
	Duration            *int64       `json:"duration"`             // Duration of sent video in seconds
	Width               *int64       `json:"width"`                // Video width
	Height              *int64       `json:"height"`               // Video height
	Caption             *string      `json:"caption"`              // Video caption (may also be used when resending videos by file_id), 0-200 characters
	ParseMode           *string      `json:"parse_mode"`           // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	SupportStreaming    *bool        `json:"support_streaming"`    // Pass True, if the uploaded video is suitable for streaming
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendVoice : Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoice struct {
	ChatID              string       `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Voice               InputFile    `json:"voice"`                // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	Caption             *string      `json:"caption"`              // Voice message caption, 0-200 characters
	ParseMode           *string      `json:"parse_mode"`           // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Duration            *int64       `json:"duration"`             // Duration of the voice message in seconds
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendVideoNote : As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
type SendVideoNote struct {
	ChatID              string       `json:"chat_id"`              // Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	VideoNote           InputFile    `json:"video_note"`           // InputFile or String. Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported
	Duration            *int64       `json:"duration"`             // Duration of sent video in seconds
	Length              *int64       `json:"length"`               // Video width and height
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply. Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendMediaGroup : Use this method to send a group of photos or videos as an album. On success, an array of the sent Messages is returned.
type SendMediaGroup struct {
	ChatID              string       `json:"chat_id"`              // String or Integer. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Media               []InputMedia `json:"media"`                // A JSON-serialized array describing photos and videos to be sent, must include 2–10 items
	DisableNotification *bool        `json:"disable_notification"` // Sends the messages silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the messages are a reply, ID of the original message
}

// SendLocation : Use this method to send point on the map. On success, the sent Message is returned.
type SendLocation struct {
	ChatID              string       `json:"chat_id"`              // String or Integer. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude            float64      `json:"latitude"`             // Latitude of the location
	Longitude           float64      `json:"longitude"`            // Longitude of the location
	LivePeriod          *int64       `json:"live_period"`          // Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// EditMessageLiveLocation : Use this method to edit live location messages sent by the bot or via the bot (for inline bots). A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageLiveLocation struct {
	ChatID          *string               `json:"chat_id"`           // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID       *int64                `json:"message_id"`        // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageID *string               `json:"inline_message_id"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	Latitude        float64               `json:"latitude"`          // Latitude of new location
	Longitude       float64               `json:"longitude"`         // Longitude of new location
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup"`      // A JSON-serialized object for a new inline keyboard.
}

// StopMessageLiveLocation : Use this method to stop updating a live location message sent by the bot or via the bot (for inline bots) before live_period expires. On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
type StopMessageLiveLocation struct {
	ChatID          *string               `json:"chat_id"`           // String or Integer. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID       *int64                `json:"message_id"`        // Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageID *string               `json:"inline_message_id"` // Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     *InlineKeyboardMarkup `json:"reply_markup"`      // A JSON-serialized object for a new inline keyboard.
}

// SendVenue : Use this method to send information about a venue. On success, the sent Message is returned.
type SendVenue struct {
	ChatID              string       `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude            float64      `json:"latitude"`             // Latitude of the venue
	Longitude           float64      `json:"longitude"`            // Longitude of the venue
	Title               string       `json:"title"`                // Name of the venue
	Address             string       `json:"address"`              // Address of the venue
	FoursquareID        *string      `json:"foursquare_id"`        // Foursquare identifier of the venue
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

// SendContact : Use this method to send phone contacts. On success, the sent Message is returned.
type SendContact struct {
	ChatID              string       `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	PhoneNumber         string       `json:"phone_number"`         // Contact's phone number
	FirstName           string       `json:"first_name"`           // Contact's first name
	LastName            *string      `json:"last_name"`            // Contact's last name
	DisableNotification *bool        `json:"disable_notification"` // Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageID    *int64       `json:"reply_to_message_id"`  // If the message is a reply, ID of the original message
	ReplyMarkup         *ReplyMarkup `json:"reply_markup"`         // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove keyboard or to force a reply from the user.
}

const (
	// Typing is an action of chat.
	Typing = "typing"
	// UploadPhoto is an action of chat.
	UploadPhoto = "upload_photo"
	// RecordVideo is an action of chat.
	RecordVideo = "record_video"
	// UploadVideo is an action of chat.
	UploadVideo = "upload_video"
	// RecordAudio is an action of chat.
	RecordAudio = "record_audio"
	// UploadAudio is an action of chat.
	UploadAudio = "upload_audio"
	// UploadDocument is an action of chat.
	UploadDocument = "upload_document"
	// FindLocation is an action of chat.
	FindLocation = "find_location"
	// RecordVideoNote is an action of chat.
	RecordVideoNote = "record_video_note"
	// UploadVideoNote is an action of chat.
	UploadVideoNote = "upload_video_note"
)

// SendChatAction : Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
type SendChatAction struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Action string `json:"action"`  // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
}

// GetFile : Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
type GetFile struct {
	FileID string `json:"file_id"` // File identifier to get info about
}
