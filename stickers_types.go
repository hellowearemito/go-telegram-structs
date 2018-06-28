package telegram

// Sticker represents a sticker.
type Sticker struct {
	FileID       string       `json:"file_id"`       // Unique identifier for this file
	Width        int64        `json:"width"`         // Sticker width
	Height       int64        `json:"height"`        // Sticker height
	Thumb        *PhotoSize   `json:"thumb"`         // Optional. Sticker thumbnail in the .webp or .jpg format
	Emoji        *string      `json:"emoji"`         // Optional. Emoji associated with the sticker
	SetName      *string      `json:"set_name"`      // Optional. Name of the sticker set to which the sticker belongs
	MaskPosition MaskPosition `json:"mask_position"` // Optional. For mask stickers, the position where the mask should be placed
	FileSize     *int64       `json:"file_size"`     // Optional. File size
}

// StickerSet represents a sticker set.
type StickerSet struct {
	Name          string    `json:"name"`           // Sticker set name
	Title         string    `json:"title"`          // Sticker set title
	ContainsMasks bool      `json:"contains_masks"` // True, if the sticker set contains masks
	Stickers      []Sticker `json:"stickers"`       // List of all set stickers
}

const (
	// Forehead is a point of mask position.
	Forehead = "forehead"
	// Eyes is a point of mask position.
	Eyes = "eyes"
	// Mouth is a point of mask position.
	Mouth = "mouth"
	// Chin is a point of mask position.
	Chin = "chin"
)

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`   // The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	Scale  float64 `json:"scale"`   // Mask scaling coefficient. For example, 2.0 means double size.
}
