package types

const (
	Private = "private"
	Group = "group"
	Supergroup = "supergroup"
	Channel = "channel"
)


// Chat represents a chat.
type Chat struct {
	ID                          int64      `json:"id"`                             // Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	Type                        string     `json:"type"`                           // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title                       *string    `json:"title"`                          // Optional. Title, for supergroups, channels and group chats
	Username                    *string    `json:"username"`                       // Optional. Username, for private chats, supergroups and channels if available
	FirstName                   *string    `json:"first_name"`                     // Optional. First name of the other party in a private chat
	LastName                    *string    `json:"last_name"`                      // Optional. Last name of the other party in a private chat
	AllMembersAreAdministrators *bool      `json:"all_members_are_administrators"` // Optional. True if a group has ‘All Members Are Admins’ enabled.
	Photo                       *ChatPhoto `json:"photo"`                          // Optional. Chat photo. Returned only in getChat.
	Description                 *string    `json:"description"`                    // Optional. Description, for supergroups and channel chats. Returned only in getChat.
	InviteLink                  *string    `json:"invite_link"`                    // Optional. Chat invite link, for supergroups and channel chats. Returned only in getChat.
	PinnedMessage               *Message   `json:"pinned_message"`                 // Optional. Pinned message, for supergroups and channel chats. Returned only in getChat.
	StickerSetName              *string    `json:"sticker_set_name"`               // Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet            *bool      `json:"can_set_sticker_set"`            // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"` // Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download.
	BigFileID   string `json:"big_file_id"`   // Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download.
}

const (
	Creator = "creator"
	Administrator = "administrator"
	Member = "member"
	Restricted = "restricted"
	Left = "left"
	Kicked = "kicked"
)

// ChatMember contains information about one member of a chat.
type ChatMember struct {
	User                  User   `json:"user"`                      // Information about the user
	Status                string `json:"status"`                    // The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
	UntilDate             *int64 `json:"until_date"`                // Optional. Restricted and kicked only. Date when restrictions will be lifted for this user, unix time
	CanBeEdited           *bool  `json:"can_be_edited"`             // Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
	CanChangeInfo         *bool  `json:"can_change_info"`           // Optional. Administrators only. True, if the administrator can change the chat title, photo and other settings
	CanPostMessages       *bool  `json:"can_post_messages"`         // Optional. Administrators only. True, if the administrator can post in the channel, channels only
	CanEditMessages       *bool  `json:"can_edit_messages"`         // Optional. Administrators only. True, if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMessages     *bool  `json:"can_delete_messages"`       // Optional. Administrators only. True, if the administrator can delete messages of other users
	CanInviteUsers        *bool  `json:"can_invite_users"`          // Optional. Administrators only. True, if the administrator can invite new users to the chat
	CanRestrictMembers    *bool  `json:"can_restrict_members"`      // Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
	CanPinMessages        *bool  `json:"can_pin_messages"`          // Optional. Administrators only. True, if the administrator can pin messages, supergroups only
	CanPromoteMessages    *bool  `json:"can_promote_messages"`      // Optional. Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanSendMessages       *bool  `json:"can_send_messages"`         // Optional. Restricted only. True, if the user can send text messages, contacts, locations and venues
	CanSendMediaMessages  *bool  `json:"can_send_media_messages"`   // Optional. Restricted only. True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendOtherMessages  *bool  `json:"can_send_other_messages"`   // Optional. Restricted only. True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePreviews *bool  `json:"can_add_web_page_previews"` // Optional. Restricted only. True, if user may add web page previews to his messages, implies can_send_media_messages
}
