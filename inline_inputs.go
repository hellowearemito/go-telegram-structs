package telegram

// InputMessageContent epresents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 4 types: InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent
type InputMessageContent interface{}

// InputTextMessageContent represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText          string  `json:"message_text"`            // Text of the message to be sent, 1-4096 characters
	ParseMode            *string `json:"parse_mode"`              // Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPageReview *bool   `json:"disable_web_page_review"` // Optional. Disables link previews for links in the sent message
}

// InputLocationMessageContent represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude   float64 `json:"latitude"`    // Latitude of the location in degrees
	Longitude  float64 `json:"longitude"`   // Longitude of the location in degrees
	LivePeriod *int64  `json:"live_period"` // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
}

// InputVenueMessageContent represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude     float64 `json:"latitude"`      // Latitude of the venue in degrees
	Longitude    float64 `json:"longitude"`     // Longitude of the venue in degrees
	Title        string  `json:"title"`         // Name of the venue
	Address      string  `json:"address"`       // Address of the venue
	FoursquareID *string `json:"foursquare_id"` // Optional. Foursquare identifier of the venue, if known
}

// InputContactMessageContent represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string  `json:"phone_number"` // Contact's phone number
	FirstName   string  `json:"first_name"`   // Contact's first name
	LastName    *string `json:"last_name"`    // Optional. Contact's last name
}

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`         // The unique identifier for the result that was chosen
	From            User      `json:"from"`              // The user that chose the result
	Location        *Location `json:"location"`          // Optional. Sender location, only for bots that require user location
	InlineMessageID *string   `json:"inline_message_id"` // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	Query           string    `json:"query"`             // The query that was used to obtain the result
}
