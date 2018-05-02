package inline

import "github.com/hellowearemito/go-telegram-structs/types"

// Query represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type Query struct {
	ID       string          `json:"id"`       // Unique identifier for this query
	From     types.User      `json:"from"`     // Sender
	Location *types.Location `json:"location"` // Optional. Sender location, only for bots that request user location
	Query    string          `json:"query"`    // Text of the query (up to 512 characters)
	Offset   string          `json:"offset"`   // Offset of the results to be returned, can be controlled by the bot
}

// BaseResult represents the basic result structure.
type BaseResult struct {
	Type                string                      `json:"type"`                  // Type of the result
	ID                  string                      `json:"id"`                    // Unique identifier for this result, 1-64 Bytes
	InputMessageContent InputMessageContent         `json:"input_message_content"` // Content of the message to be sent
	ReplyMarkup         *types.InlineKeyboardMarkup `json:"reply_markup"`          // Optional. Inline keyboard attached to the message
}

// BaseResultWithCaption represents the basic result with captions fields.
type BaseResultWithCaption struct {
	BaseResult
	Caption   *string `json:"caption"`    // Optional. Caption of the photo to be sent, 0-200 characters
	ParseMode *string `json:"parse_mode"` // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
}

// ResultArticle represents a link to an article or web page.
type ResultArticle struct {
	BaseResult
	Title       string  `json:"title"`        // Title of the result
	URL         *string `json:"url"`          // Optional. URL of the result
	HideURL     *bool   `json:"hide_url"`     // Optional. Pass True, if you don't want the URL to be shown in the message
	Description *string `json:"description"`  // Optional. Short description of the result
	ThumbURL    *string `json:"thumb_url"`    // Optional. Url of the thumbnail for the result
	ThumbWidth  *int64  `json:"thumb_width"`  // Optional. Thumbnail width
	ThumbHeight *int64  `json:"thumb_height"` // Optional. Thumbnail height
}

// ResultPhoto represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type ResultPhoto struct {
	BaseResultWithCaption
	PhotoURL    string  `json:"photo_url"`    // A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
	ThumbURL    string  `json:"thumb_url"`    // URL of the thumbnail for the photo
	PhotoWidth  *int64  `json:"photo_width"`  // Optional. Width of the photo
	PhotoHeigth *int64  `json:"photo_heigth"` // Optional. Height of the photo
	Title       *string `json:"title"`        // Optional. Title for the result
	Description *string `json:"description"`  // Optional. Short description of the result
}

// ResultGif represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type ResultGif struct {
	BaseResultWithCaption
	URL      string  `json:"gif_url"`      // A valid URL for the GIF file. File size must not exceed 1MB
	Width    *int64  `json:"gif_width"`    // Optional. Width of the GIF
	Height   *int64  `json:"gif_height"`   // Optional. Height of the GIF
	Duration *int64  `json:"gif_duration"` // Optional. Duration of the GIF
	ThumbURL string  `json:"thumb_url"`    // URL of the static thumbnail for the result (jpeg or gif)
	Title    *string `json:"title"`        // Optional. Title for the result
}

// ResultMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type ResultMpeg4Gif struct {
	BaseResultWithCaption
	URL      string  `json:"mpeg4_url"`      // A valid URL for the MP4 file. File size must not exceed 1MB
	Width    *int64  `json:"mpeg4_width"`    // Optional. Video width
	Height   *int64  `json:"mpeg4_height"`   // Optional. Video height
	Duration *int64  `json:"mpeg4_duration"` // Optional. Video duration
	ThumbURL string  `json:"thumb_url"`      // URL of the static thumbnail (jpeg or gif) for the result
	Title    *string `json:"title"`          // Optional. Title for the result
}

// ResultVideo represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type ResultVideo struct {
	BaseResultWithCaption
	URL         string  `json:"video_url"`      // A valid URL for the embedded video player or video file
	MimeType    string  `json:"mime_type"`      // Mime type of the content of video url, “text/html” or “video/mp4”
	ThumbURL    string  `json:"thumb_url"`      // URL of the thumbnail (jpeg only) for the video
	Title       string  `json:"title"`          // Title for the result
	Width       *int64  `json:"video_width"`    // Optional. Video width
	Height      *int64  `json:"video_height"`   // Optional. Video height
	Duration    *int64  `json:"video_duration"` // Optional. Video duration in seconds
	Description *string `json:"description"`    // Optional. Short description of the result
}

// ResultAudio represents a link to an mp3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type ResultAudio struct {
	BaseResultWithCaption
	URL       string  `json:"audio_url"`      // A valid URL for the audio file
	Title     string  `json:"title"`          // Title
	Performer *string `json:"performer"`      // Optional. Performer
	Duration  *int64  `json:"audio_duration"` // Optional. Audio duration in seconds
}

// ResultVoice represents a link to a voice recording in an .ogg container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type ResultVoice struct {
	BaseResultWithCaption
	URL      string `json:"voice_url"`      // A valid URL for the voice recording
	Title    string `json:"title"`          // Recording title
	Duration *int64 `json:"voice_duration"` // Optional. Recording duration in seconds
}

// ResultDocument represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type ResultDocument struct {
	BaseResultWithCaption
	Title       string  `json:"title"`        // Title for the result
	URL         string  `json:"document_url"` // A valid URL for the file
	MimeType    string  `json:"mime_type"`    // Mime type of the content of the file, either “application/pdf” or “application/zip”
	Description *string `json:"description"`  // Optional. Short description of the result
	ThumbURL    *string `json:"thumb_url"`    // Optional. URL of the thumbnail (jpeg only) for the file
	ThumbWidth  *string `json:"thumb_width"`  // Optional. Thumbnail width
	ThumbHeight *string `json:"thumb_height"` // Optional. Thumbnail height
}

// ResultLocation represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
type ResultLocation struct {
	BaseResult
	Latitude    float64 `json:"latitude"`     // Location latitude in degrees
	Longitude   float64 `json:"longitude"`    // Location longitude in degrees
	Title       string  `json:"title"`        // Location title
	LivePeriod  *int64  `json:"live_period"`  // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	ThumbURL    *int64  `json:"thumb_url"`    // Optional. Url of the thumbnail for the result
	ThumbWidth  *int64  `json:"thumb_width"`  // Optional. Thumbnail width
	ThumbHeight *int64  `json:"thumb_height"` // Optional. Thumbnail height
}

// ResultVenue represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type ResultVenue struct {
	BaseResult
	Latitude     float64 `json:"latitude"`      // Latitude of the venue location in degrees
	Longitude    float64 `json:"longitude"`     // Longitude of the venue location in degrees
	Title        string  `json:"title"`         // Title of the venue
	Address      string  `json:"address"`       // Address of the venue
	FoursquareID *string `json:"foursquare_id"` // Optional. Foursquare identifier of the venue if known
	ThumbURL     *int64  `json:"thumb_url"`     // Optional. Url of the thumbnail for the result
	ThumbWidth   *int64  `json:"thumb_width"`   // Optional. Thumbnail width
	ThumbHeight  *int64  `json:"thumb_height"`  // Optional. Thumbnail height
}

// ResultContact represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
type ResultContact struct {
	BaseResult
	PhoneNumber string  `json:"phone_number"` // Contact's phone number
	FirstName   string  `json:"first_name"`   // Contact's first name
	LastName    *string `json:"last_name"`    // Optional. Contact's last name
	ThumbURL    *int64  `json:"thumb_url"`    // Optional. Url of the thumbnail for the result
	ThumbWidth  *int64  `json:"thumb_width"`  // Optional. Thumbnail width
	ThumbHeight *int64  `json:"thumb_height"` // Optional. Thumbnail height
}

// ResultGame represents a Game.
type ResultGame struct {
	Type          string                      `json:"type"`            // Type of the result
	ID            string                      `json:"id"`              // Unique identifier for this result, 1-64 Bytes
	GameShortName string                      `json:"game_short_name"` // Short name of the game
	ReplyMarkup   *types.InlineKeyboardMarkup `json:"reply_markup"`    // Optional. Inline keyboard attached to the message
}

// ResultCachedPhoto represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type ResultCachedPhoto struct {
	BaseResultWithCaption
	PhotoFileID string  `json:"photo_file_id"` // A valid file identifier of the photo
	Title       *string `json:"title"`         // Optional. Title for the result
	Description *string `json:"description"`   // Optional. Short description of the result
}

// ResultCachedGif represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type ResultCachedGif struct {
	BaseResultWithCaption
	GifFileID string  `json:"gif_file_id"` // A valid file identifier for the GIF file
	Title     *string `json:"title"`       // Optional. Title for the result
}

// ResultCachedMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type ResultCachedMpeg4Gif struct {
	BaseResultWithCaption
	Mpeg4FileID string  `json:"mpeg_4_file_id"` // A valid file identifier for the MP4 file
	Title       *string `json:"title"`          // Optional. Title for the result
}

// ResultCachedSticker represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the
type ResultCachedSticker struct {
	BaseResult
	StickerFileID string `json:"sticker_file_id"` // A valid file identifier of the sticker
}

// ResultCachedDocument represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
type ResultCachedDocument struct {
	BaseResultWithCaption
	Title          string  `json:"title"`            // Title for the result
	DocumentFileID string  `json:"document_file_id"` // A valid file identifier for the file
	Description    *string `json:"description"`      // Optional. Short description of the result
}

// ResultCachedVideo represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type ResultCachedVideo struct {
	VideoFileID string  `json:"video_file_id"` // A valid file identifier for the video file
	Title       string  `json:"title"`         // Title for the result
	Description *string `json:"description"`   // Optional. Short description of the result
}

// ResultCachedVoice represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type ResultCachedVoice struct {
	BaseResultWithCaption
	VoiceFileID string `json:"voice_file_id"` // A valid file identifier for the voice message
	Title       string `json:"title"`         // Voice message title
}

// ResultCachedAudio represents a link to an mp3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type ResultCachedAudio struct {
	BaseResultWithCaption
	AudioFileID string `json:"audio_file_id"` // A valid file identifier for the audio file
}
