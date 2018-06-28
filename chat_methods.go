package telegram

// KickChatMember : Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type KickChatMember struct {
	ChatID    string `json:"chat_id"`    // String or Integer. Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserID    int64  `json:"user_id"`    // Unique identifier of the target user
	UntilDate *int64 `json:"until_date"` // Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
}

// UnbanChatMember : Use this method to unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. Returns True on success.
type UnbanChatMember struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
	UserID int64  `json:"user_id"` // Unique identifier of the target user
}

// RestrictChatMember : Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights. Pass True for all boolean parameters to lift restrictions from a user. Returns True on success.
type RestrictChatMember struct {
	ChatID                string `json:"chat_id"`                   // String or integer. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID                int64  `json:"user_id"`                   // Unique identifier of the target user
	UntilDate             *int64 `json:"until_date"`                // Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
	CanSendMessages       *bool  `json:"can_send_messages"`         // Pass True, if the user can send text messages, contacts, locations and venues
	CanSendMediaMessages  *bool  `json:"can_send_media_messages"`   // Pass True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendOtherMessages  *bool  `json:"can_send_other_messages"`   // Pass True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePreviews *bool  `json:"can_add_web_page_previews"` // Pass True, if the user may add web page previews to their messages, implies can_send_media_messages
}

// PromoteChatMember : Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters to demote a user. Returns True on success.
type PromoteChatMember struct {
	ChatID             string `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID             int64  `json:"user_id"`              // Unique identifier of the target user
	CanChaneInfo       *bool  `json:"can_chane_info"`       // Pass True, if the administrator can change chat title, photo and other settings
	CanPostMessages    *bool  `json:"can_post_messages"`    // Pass True, if the administrator can create channel posts, channels only
	CanEditMessages    *bool  `json:"can_edit_message"`     // Pass True, if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMessages  *bool  `json:"can_delete_messages"`  // Pass True, if the administrator can delete messages of other users
	CanInviteUsers     *bool  `json:"can_invite_users"`     // Pass True, if the administrator can invite new users to the chat
	CanRestrictMembers *bool  `json:"can_restrict_members"` // Pass True, if the administrator can restrict, ban or unban chat members
	CanPinMessages     *bool  `json:"can_pin_messages"`     // Pass True, if the administrator can pin messages, supergroups only
	CanPromoteMembers  *bool  `json:"can_promote_members"`  // Pass True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by him)
}

// ExportChatInviteLink : Use this method to generate a new invite link for a chat; any previously generated link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the new invite link as String on success.
type ExportChatInviteLink struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// SetChatPhoto : Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatPhoto struct {
	ChatID string    `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  InputFile `json:"photo"`   // New chat photo, uploaded using multipart/form-data
}

// DeleteChatPhoto : Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type DeleteChatPhoto struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// SetChatTitle : Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatTitle struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string `json:"title"`   // New chat title, 1-255 characters
}

// SetChatDescription : Use this method to change the description of a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatDescription struct {
	ChatID      string  `json:"chat_id"`     // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Description *string `json:"description"` // New chat description, 0-255 characters
}

// PinChatMessage : Use this method to pin a message in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
type PinChatMessage struct {
	ChatID              string `json:"chat_id"`              // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageID           int64  `json:"message_id"`           // Identifier of a message to pin
	DisableNotification *bool  `json:"disable_notification"` // Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels.
}

// UnpinChatMessage : Use this method to unpin a message in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
type UnpinChatMessage struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

// LeaveChat : Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
type LeaveChat struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChat : Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
type GetChat struct {
	ChatID string `json:"chat_id"` // 	Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatAdministrators : Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
type GetChatAdministrators struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatMembersCount : Use this method to get the number of members in a chat. Returns Int on success.
type GetChatMembersCount struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

// GetChatMember : Use this method to get information about a member of a chat. Returns a ChatMember object on success.
type GetChatMember struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the t
	UserID int64  `json:"user_id"` // Unique identifier of the target user
}

// SetChatStickerSet : Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type SetChatStickerSet struct {
	ChatID         string `json:"chat_id"`          // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	StickerSetName string `json:"sticker_set_name"` // Name of the sticker set to be set as the group sticker set
}

// DeleteChatStickerSet : Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type DeleteChatStickerSet struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

// AnswerCallbackQuery : Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
type AnswerCallbackQuery struct {
	CallbackQueryID string  `json:"callback_query_id"` // Unique identifier for the query to be answered
	Text            *string `json:"text"`              // Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert       *bool   `json:"show_alert"`        // If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	URL             *string `json:"url"`               // URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime       *int64  `json:"cache_time"`        // The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}
