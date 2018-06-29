package telegram

import "encoding/json"

// Decode decodes the json data and set the type of message.
func Decode(data json.RawMessage) (*Message, error) {
	message := Message{}

	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}

	if message.Text != nil {
		message.Type = TextType
	} else if message.Audio != nil {
		message.Type = AudioType
	} else if message.Video != nil {
		message.Type = VideoType
	} else if message.Document != nil {
		message.Type = DocumentType
	} else if message.Game != nil {
		message.Type = GameType
	} else if message.Photo != nil {
		message.Type = PhotoType
	} else if message.Sticker != nil {
		message.Type = StickerType
	} else if message.Voice != nil {
		message.Type = VoiceType
	} else if message.VideoNote != nil {
		message.Type = VideoNoteType
	} else if message.Contact != nil {
		message.Type = ContactType
	} else if message.Location != nil {
		message.Type = LocationType
	} else if message.Venue != nil {
		message.Type = VenueType
	} else if message.Invoice != nil {
		message.Type = InvoiceType
	} else if message.SuccessfulPayment != nil {
		message.Type = SuccessfulPaymentType
	}

	return &message, nil
}
