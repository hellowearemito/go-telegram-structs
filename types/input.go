package types

// InputMedia represents the content of a media message to be sent. It should be one of InputMediaPhoto or InputMediaVideo
type InputMedia interface {}

// InputMediaPhoto represents a photo to be sent.
type InputMediaPhoto struct {
	Type      string  `json:"type"`       // Type of the result, must be photo
	Media     string  `json:"media"`      // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name.
	Caption   *string `json:"caption"`    // Optional. Caption of the photo to be sent, 0-200 characters
	ParseMode *string `json:"parse_mode"` // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
}

// InputMediaVideo represents a video to be sent.
type InputMediaVideo struct {
	Type              string  `json:"type"`               // Type of the result, must be video
	Media             string  `json:"media"`              // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name.
	Caption           *string `json:"caption"`            // Optional. Caption of the video to be sent, 0-200 characters
	ParseMode         *string `json:"parse_mode"`         // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Width             *int64  `json:"width"`              // Optional. Video width
	Height            *int64  `json:"height"`             // Optional. Video height
	Duration          *int64  `json:"duration"`           // Optional. Video duration
	SupportsStreaming *bool   `json:"supports_streaming"` // Optional. Pass True, if the uploaded video is suitable for streaming
}

// InputFile represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile interface{}
