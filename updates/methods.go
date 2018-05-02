package updates

import "github.com/hellowearemito/go-telegram-structs/types"

// GetUpdates
type GetUpdates struct {
	Offset *int64 `json:"offset"` // Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Limit *int64 `json:"limit"` // Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
	Timeout *int64 `json:"timeout"` // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []string `json:"allowed_updates"` // List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.  Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
}

// SetWebhook : Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns true. If you'd like to make sure that the Webhook request comes from Telegram, we recommend using a secret path in the URL, e.g. https://www.example.com/<token>. Since nobody else knows your bot‘s token, you can be pretty sure it’s us.
type SetWebhook struct {
	URL string `json:"url"` // HTTPS url to send updates to. Use an empty string to remove webhook integration
	Certificate *types.InputFile `json:"certificate"` // Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	MaxConnections *int64 `json:"max_connections"` // Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server, and higher values to increase your bot’s throughput.
	AllowedUpdates []string `json:"allowed_updates"` // List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
}

// WebhookInfo
type WebhookInfo struct {
	URL string `json:"url"` // Webhook URL, may be empty if webhook is not set up
	HasCustomCertificate bool `json:"has_custom_certificate"` // True, if a custom certificate was provided for webhook certificate checks
	PendingUpdateCount int64 `json:"pending_update_count"` // Number of updates awaiting delivery
	LastErrorDate *int64 `json:"last_error_date"` // Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage *string `json:"last_error_message"` // Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	MaxConnections *int64 `json:"max_connections"` // Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	AllowedUpdates []string `json:"allowed_updates"` // Optional. A list of update types the bot is subscribed to. Defaults to all update types
}