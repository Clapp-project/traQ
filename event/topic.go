package event

const (
	// UserCreated ユーザーが追加された
	// 	Fields:
	//      user_id: uuid.UUID
	// 		user: *model.User
	UserCreated = "user.created"
	// UserUpdated ユーザーが更新された
	// 	Fields:
	//      user_id: uuid.UUID
	UserUpdated = "user.updated"
	// UserIconUpdated ユーザーのアイコンが更新された
	// 	Fields:
	// 		user_id: uuid.UUID
	// 		file_id: uuid.UUID
	UserIconUpdated = "user.icon_updated"
	// UserOnline ユーザーがオンラインになった
	// 	Fields:
	//      user_id: uuid.UUID
	UserOnline = "user.online"
	// UserOffline ユーザーがオフラインになった
	// 	Fields:
	//      user_id: uuid.UUID
	// 		datetime: time.Time
	UserOffline = "user.offline"

	// UserTagAdded ユーザーにタグが追加された
	// 	Fields:
	//      user_id: uuid.UUID
	// 		tag_id: uuid.UUID
	UserTagAdded = "user_tag.added"
	// UserTagUpdated ユーザーのタグが更新された
	// 	Fields:
	//      user_id: uuid.UUID
	// 		tag_id: uuid.UUID
	UserTagUpdated = "user_tag.updated"
	// UserTagRemoved ユーザーからタグが削除された
	// 	Fields:
	//      user_id: uuid.UUID
	// 		tag_id: uuid.UUID
	UserTagRemoved = "user_tag.deleted"

	// UserGroupCreated ユーザーグループが作成された
	// 	Fields:
	// 		group_id: uuid.UUID
	// 		group: *model.UserGroup
	UserGroupCreated = "user_group.created"
	// UserGroupDeleted ユーザーグループが削除された
	// 	Fields:
	// 		group_id: uuid.UUID
	UserGroupDeleted = "user_group.deleted"
	// UserGroupMemberAdded ユーザーがグループに追加された
	// 	Fields:
	// 		group_id: uuid.UUID
	// 		user_id: uuid.UUID
	UserGroupMemberAdded = "user_group.member.added"
	// UserGroupMemberRemoved ユーザーがグループから削除された
	// 	Fields:
	// 		group_id: uuid.UUID
	// 		user_id: uuid.UUID
	UserGroupMemberRemoved = "user_group.member.removed"

	// MessageCreated メッセージが作成された
	// 	Fields:
	// 		message_id: uuid.UUID
	//  	message: *model.Message
	//  	embedded: []*message.EmbeddedInfo
	//      plain: string
	MessageCreated = "message.created"
	// MessageUpdated メッセージが更新された
	// 	Fields:
	// 		message_id: uuid.UUID
	//  	message: *model.Message
	//  	old_message: *model.Message
	MessageUpdated = "message.updated"
	// MessageDeleted メッセージが削除された
	// 	Fields:
	// 		message_id: uuid.UUID
	//  	message: *model.Message
	MessageDeleted = "message.deleted"
	// MessageStamped メッセージにスタンプが押された
	// 	Fields:
	// 		message_id: uuid.UUID
	// 		user_id: uuid.UUID
	// 		stamp_id: uuid.UUID
	// 		count: int
	// 		created_at: time.Time
	MessageStamped = "message.stamped"
	// MessageUnstamped メッセージからスタンプが消された
	// 	Fields:
	// 		message_id: uuid.UUID
	// 		user_id: uuid.UUID
	// 		stamp_id: uuid.UUID
	MessageUnstamped = "message.unstamped"
	// MessagePinned メッセージがピンされた
	// 	Fields:
	// 		message_id: uuid.UUID
	// 		channel_id: uuid.UUID
	// 		pin_id: uuid.UUID
	MessagePinned = "message.pinned"
	// MessageUnpinned メッセージがピンから外れた
	// 	Fields:
	// 		message_id: uuid.UUID
	// 		channel_id: uuid.UUID
	// 		pin_id: uuid.UUID
	MessageUnpinned = "message.unpinned"

	// ChannelCreated チャンネルが作成された
	// 	Fields:
	// 		channel_id: uuid.UUID
	// 		channel: *model.Channel
	// 		private: bool
	ChannelCreated = "channel.created"
	// ChannelUpdated チャンネルが更新された
	// 	Fields:
	// 		channel_id: uuid.UUID
	// 		private: bool
	ChannelUpdated = "channel.updated"
	// ChannelTopicUpdated チャンネルトピックが更新された
	// 	Fields:
	// 		channel_id: uuid.UUID
	// 		topic: string
	// 		updater_id: uuid.UUID
	ChannelTopicUpdated = "channel.topic.updated"
	// ChannelDeleted チャンネルが削除された
	// 	Fields:
	// 		channel_id: uuid.UUID
	// 		private: bool
	ChannelDeleted = "channel.deleted"
	// ChannelRead チャンネルのメッセージが既読された
	// 	Fields:
	// 		user_id: uuid.UUID
	// 		channel_id: uuid.UUID
	ChannelRead = "channel.read"
	// ChannelStared チャンネルがスターされた
	// 	Fields:
	// 		user_id: uuid.UUID
	// 		channel_id: uuid.UUID
	ChannelStared = "channel.stared"
	// ChannelUnstared チャンネルのスターが解除された
	// 	Fields:
	// 		user_id: uuid.UUID
	// 		channel_id: uuid.UUID
	ChannelUnstared = "channel.unstared"
	// ChannelViewersChanged チャンネルの閲覧者が変化した
	// 	Fields:
	// 		channel_id: uuid.UUID
	// 		viewers: map[uuid.UUID]realtime.ViewState
	ChannelViewersChanged = "channel.viewers_changed"

	// StampCreated スタンプが作成された
	// 	Fields:
	// 		stamp_id: uuid.UUID
	// 		stamp: *model.Stamp
	StampCreated = "stamp.created"
	// StampUpdated スタンプが更新された
	// 	Fields:
	// 		stamp_id: uuid.UUID
	StampUpdated = "stamp.updated"
	// StampDeleted スタンプが削除された
	// 	Fields:
	// 		stamp_id: uuid.UUID
	StampDeleted = "stamp.deleted"

	// StampPaletteCreated スタンプパレットが作成された
	// 	Fields:
	//		user_id: uuid.UUID
	// 		stamp_palette_id: uuid.UUID
	// 		stamp_palette: *model.StampPalette
	StampPaletteCreated = "stamp_palette.created"
	// StampPaletteUpdated スタンプパレットが更新された
	// 	Fields:
	//		user_id: uuid.UUID
	// 		stamp_palette_id: uuid.UUID
	StampPaletteUpdated = "stamp_palette.updated"
	// StampPaletteDeleted スタンプパレットが削除された
	// 	Fields:
	//		user_id: uuid.UUID
	// 		stamp_palette_id: uuid.UUID
	StampPaletteDeleted = "stamp_palette.deleted"

	// WebhookCreated Webhookが作成された
	// 	Fields:
	// 		webhook_id: uuid.UUID
	// 		webhook: Webhook
	WebhookCreated = "webhook.created"
	// WebhookUpdated Webhookが更新された
	// 	Fields:
	// 		webhook_id: uuid.UUID
	WebhookUpdated = "webhook.updated"
	// WebhookDeleted Webhookが削除された
	// 	Fields:
	// 		webhook_id: uuid.UUID
	WebhookDeleted = "webhook.deleted"

	// BotCreated Botが作成された
	// 	Fields:
	// 		bot_id: uuid.UUID
	// 		bot: *model.Bot
	BotCreated = "bot.created"
	// BotUpdated Botが更新された
	// 	Fields:
	// 		bot_id: uuid.UUID
	BotUpdated = "bot.updated"
	// BotDeleted Botが削除された
	// 	Fields:
	// 		bot_id: uuid.UUID
	BotDeleted = "bot.deleted"
	// BotStateChanged Botの状態が変化した
	// 	Fields:
	// 		bot_id: uuid.UUID
	// 		state: model.BotState
	BotStateChanged = "bot.state_changed"
	// BotPingRequest BotのPingがリクエストされた
	// 	Fields:
	// 		bot_id: uuid.UUID
	// 		bot: *model.Bot
	BotPingRequest = "bot.ping"
	// BotJoined Botがチャンネルに参加した
	// 	Fields:
	// 		bot_id: uuid.UUID
	// 		channel_id: uuid.UUID
	BotJoined = "bot.joined"
	// BotLeft Botがチャンネルから退出した
	// 	Fields:
	// 		bot_id: uuid.UUID
	// 		channel_id: uuid.UUID
	BotLeft = "bot.left"

	// UserWebRTCStateChanged ユーザーのWebRTCの状態が変化した
	// 	Fields:
	// 		user_id: uuid.UUID
	// 		channel_id: uuid.UUID
	// 		state: string
	UserWebRTCStateChanged = "user.webrtc.state_changed"

	// SSEConnected ユーザーがSSEストリームに接続した
	// 	Fields:
	// 		user_id: uuid.UUID
	//		req: *http.Request
	SSEConnected = "sse.connected"
	// SSEDisconnected ユーザーがSSEストリームから切断した
	// 	Fields:
	// 		user_id: uuid.UUID
	//		req: *http.Request
	SSEDisconnected = "sse.disconnected"

	// WSConnected ユーザーがWSストリームに接続した
	// 	Fields:
	// 		user_id: uuid.UUID
	//		req: *http.Request
	WSConnected = "ws.connected"
	// WSDisconnected ユーザーがWSストリームから切断した
	// 	Fields:
	// 		user_id: uuid.UUID
	//		req: *http.Request
	WSDisconnected = "ws.disconnected"

	// ClipFolderCreated クリップフォルダーが作成された
	// 	Fields:
	// 		user_id: uuid.UUID
	//		clip_folder_id: uuid.UUID
	//		clip_folder: *model.ClipFolder
	ClipFolderCreated = "clip_folder.created"
	// ClipFolderUpdated クリップフォルダーが更新された
	// 	Fields:
	// 		user_id: uuid.UUID
	//		clip_folder_id: uuid.UUID
	//		old_clip_folder: *model.ClipFolder
	//		clip_folder: *model.ClipFolder
	ClipFolderUpdated = "clip_folder.updated"
	// ClipFolderDeleted クリップフォルダーが削除された
	// 	Fields:
	// 		user_id: uuid.UUID
	//		clip_folder_id: uuid.UUID
	//		clip_folder: *model.ClipFolder
	ClipFolderDeleted = "clip_folder.deleted"
	// ClipFolderMessageDeleted クリップフォルダーのメッセージが除外された
	// 	Fields:
	// 		user_id: uuid.UUID
	//		clip_folder_id: uuid.UUID
	//		clip_folder_message_id: uuid.UUID
	//		clip_folder_message: *model.ClipFolderMessage
	ClipFolderMessageDeleted = "clip_folder_message.deleted"
	// ClipFolderMessageAdded クリップフォルダーにメッセージが追加された
	// 	Fields:
	// 		user_id: uuid.UUID
	//		clip_folder_id: uuid.UUID
	//		clip_folder_message_id: uuid.UUID
	//		clip_folder_message: *model.ClipFolderMessage
	ClipFolderMessageAdded = "clip_folder_message.added"
)
