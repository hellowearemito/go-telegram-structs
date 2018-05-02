package types

import (
	"github.com/hellowearemito/go-telegram-structs/games"
	"github.com/hellowearemito/go-telegram-structs/payments"
	"github.com/hellowearemito/go-telegram-structs/stickers"
)

// Message represents a message.
type Message struct {
	MessageID             int64              `json:"message_id"`              // Unique message identifier inside this chat
	From                  *User              `json:"from"`                    // Optional. Sender, empty for messages sent to channels
	Date                  int64              `json:"date""`                   // Date the message was sent in Unix time
	Chat                  Chat               `json:"chat"`                    // Conversation the message belongs to
	ForwardFrom           *User              `json:"forward_from"`            // Optional. For forwarded messages, sender of the original message
	ForwardFromChat       *Chat              `json:"forward_from_chat"`       // Optional. For messages forwarded from channels, information about the original channel
	ForwardFromMessageID  *int64             `json:"forward_from_message_id"` // Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardSignature      *string            `json:"forward_signature"`       // Optional. For messages forwarded from channels, signature of the post author if present
	ForwardDate           *int64             `json:"forward_date"`            // Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage        *Message           `json:"reply_to_message"`        // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	EditDate              *int64             `json:"edit_date"`               // Optional. Date the message was last edited in Unix time
	MediaGroupID          *string            `json:"media_group_id"`          // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature       *string            `json:"author_signature"`        // Optional. Signature of the post author for messages in channels
	Text                  *string            `json:"text"`                    // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Entities              []MessageEntity     `json:"entities"`                // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	CaptionEntities       []MessageEntity     `json:"caption_entities"`        // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	Audio                 *Audio             `json:"audio"`                   // Optional. Message is an audio file, information about the file
	Document              *Document          `json:"document"`                // Optional. Message is a general file, information about the file
	Game                  *games.Game              `json:"game"`                    // Optional. Message is a game, information about the game.
	Photo                 []PhotoSize             `json:"photo"`                   // Optional. Message is a photo, available sizes of the photo
	Sticker               *stickers.Sticker           `json:"sticker"`                 // Optional. Message is a sticker, information about the sticker
	Video                 *Video             `json:"video"`                   // Optional. Message is a video, information about the video
	Voice                 *Voice             `json:"voice"`                   // Optional. Message is a voice message, information about the file
	VideoNote             *VideoNote         `json:"video_note"`              // Optional. Message is a video note, information about the video message
	Caption               *string            `json:"caption"`                 // Optional. Caption for the audio, document, photo, video or voice, 0-200 characters
	Contact               *Contact           `json:"contact"`                 // Optional. Message is a shared contact, information about the contact
	Location              *Location          `json:"location"`                // Optional. Message is a shared location, information about the location
	Venue                 *Venue             `json:"venue"`                   // Optional. Message is a venue, information about the venue
	NewChatMembers        []User             `json:"new_chat_members"`        // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember        *User              `json:"left_chat_member"`        // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle          *string            `json:"new_chat_title"`          // Optional. A chat title was changed to this value
	NewChatPhoto          []PhotoSize        `json:"new_chat_photo"`          // Optional. A chat photo was change to this value
	DeleteChatPhoto       *bool              `json:"delete_chat_photo"`       // Optional. Service message: the chat photo was deleted
	GroupChatCreated      *bool              `json:"group_chat_created"`      // Optional. Service message: the group has been created
	SupergroupChatCreated *bool              `json:"supergroup_chat_created"` // Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated    *bool              `json:"channel_chat_created"`    // Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MigrateToChatID       *int64             `json:"migrate_to_chat_id"`      // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatID     *int64             `json:"migrate_from_chat_id"`    // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	PinnedMessage         *Message           `json:"pinned_message"`          // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	Invoice               *payments.Invoice           `json:"invoice"`                 // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment     *payments.SuccessfulPayment `json:"successful_payment"`      // Optional. Message is a service message about a successful payment, information about the payment.
	ConnectedWebsite      *string            `json:"connected_website"`       // Optional. The domain name of the website on which the user has logged in.
}

const (
	Hashtag = "hashtag"
	BotCommand = "bot_command"
	Url = "url"
	Email = "email"
	Bold = "bold"
	Italic = "italic"
	Code = "code"
	Pre = "pre"
	TextLink = "text_link"
	TextMention = "text_mention"
	)

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type   string  `json:"type"`   // Type of the entity. Can be mention (@username), hashtag, bot_command, url, email, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
	Offset int64   `json:"offset"` // Offset in UTF-16 code units to the start of the entity
	Length int64   `json:"length"` // Length of the entity in UTF-16 code units
	URL    *string `json:"url"`    // Optional. For “text_link” only, url that will be opened after user taps on the text
	User   *User   `json:"user"`   // Optional. For “text_mention” only, the mentioned user
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID   string `json:"file_id"`   // Unique identifier for this file
	Width    int64  `json:"width"`     // Photo width
	Height   int64  `json:"height"`    // Photo height
	FileSize *int64 `json:"file_size"` // Optional. File size
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID    string  `json:"file_id"`   // Unique identifier for this file
	Duration  int64   `json:"duration"`  // Duration of the audio in seconds as defined by sender
	Performer *string `json:"performer"` // Optional. Performer of the audio as defined by sender or by audio tags
	Title     *string `json:"title"`     // Optional. Title of the audio as defined by sender or by audio tags
	MimeType  *string `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	FileSize  *int64  `json:"file_size"` // Optional. File size
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID   string     `json:"file_id"`   // Unique file identifier
	Thumb    *PhotoSize `json:"thumb"`     // Optional. Document thumbnail as defined by sender
	FileName *string    `json:"file_name"` // Optional. Original filename as defined by sender
	MimeType *string    `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	FileSize *int64     `json:"file_size"` // Optional. File size
}

// Video represents a video file.
type Video struct {
	FileID   string     `json:"file_id"`   // Unique identifier for this file
	Width    int64      `json:"width"`     // Video width as defined by sender
	Height   int64      `json:"height"`    // Video height as defined by sender
	Duration int64      `json:"duration"`  // Duration of the video in seconds as defined by sender
	Thumb    *PhotoSize `json:"thumb"`     // Optional. Video thumbnail
	MimeType *string    `json:"mime_type"` // Optional. Mime type of a file as defined by sender
	FileSize *int64     `json:"file_size"` // Optional. File size
}

// Voice represents a voice note.
type Voice struct {
	FileID   string  `json:"file_id"`   // Unique identifier for this file
	Duration int64   `json:"duration"`  // Duration of the audio in seconds as defined by sender
	MimeType *string `json:"mime_type"` // Optional. MIME type of the file as defined by sender
	FileSize *int64  `json:"file_size"` // Optional. File size
}

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileID   string     `json:"file_id"`   // Unique identifier for this file
	Length   int64      `json:"length"`    // Video width and height as defined by sender
	Duration int64      `json:"duration"`  // Duration of the video in seconds as defined by sender
	Thumb    *PhotoSize `json:"thumb"`     // Optional. Video thumbnail
	FileSize *int64     `json:"file_size"` // Optional. File size
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string  `json:"phone_number"` // Contact's phone number
	FirstName   string  `json:"first_name"`   // Contact's first name
	LastName    *string `json:"last_name"`    // Optional. Contact's last name
	UserID      *int64  `json:"user_id"`      // Optional. Contact's user identifier in Telegram
}

// Location represents a point on the map.
type Location struct {
	Longitude float64 `json:"longitude"` // Longitude as defined by sender
	Latitude  float64 `json:"latitude"`  // Latitude as defined by sender
}

// Venue represents a venue.
type Venue struct {
	Location     Location `json:"location"`      // Venue location
	Title        string   `json:"title"`         // Name of the venue
	Address      string   `json:"address"`       // Address of the venue
	FoursquareID *string  `json:"foursquare_id"` // Optional. Foursquare identifier of the venue
}

// UserProfilePhotos represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int64         `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}

// File represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
type File struct {
	FileID   string  `json:"file_id"`   // Unique identifier for this file
	FileSize *int64  `json:"file_size"` // Optional. File size, if known
	FilePath *string `json:"file_path"` // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

//
