package model

import (
	"github.com/gofrs/uuid"
	"time"
)

// Message データベースに格納するmessageの構造体
type Message struct {
	ID        uuid.UUID  `gorm:"type:char(36);not null;primary_key"`
	UserID    uuid.UUID  `gorm:"type:char(36);not null;"`
	ChannelID uuid.UUID  `gorm:"type:char(36);not null;index"`
	Text      string     `sql:"type:TEXT COLLATE utf8mb4_bin NOT NULL"`
	CreatedAt time.Time  `gorm:"precision:6;index"`
	UpdatedAt time.Time  `gorm:"precision:6"`
	DeletedAt *time.Time `gorm:"precision:6"`

	Stamps []MessageStamp `gorm:"association_autoupdate:false;association_autocreate:false;preload:false;foreignkey:MessageID"`
	Pin    *Pin           `gorm:"association_autoupdate:false;association_autocreate:false;preload:false;foreignkey:MessageID"`
}

// TableName DBの名前を指定するメソッド
func (m *Message) TableName() string {
	return "messages"
}

// ChannelLatestMessage チャンネル別最新メッセージ
type ChannelLatestMessage struct {
	ChannelID uuid.UUID `gorm:"type:char(36);not null;primary_key"`
	MessageID uuid.UUID `gorm:"type:char(36);not null;"`
	DateTime  time.Time `gorm:"precision:6;index"`
}

// TableName テーブル名
func (m *ChannelLatestMessage) TableName() string {
	return "channel_latest_messages"
}

// Unread 未読レコード
type Unread struct {
	UserID     uuid.UUID `gorm:"type:char(36);not null;primary_key"`
	MessageID  uuid.UUID `gorm:"type:char(36);not null;primary_key"`
	Noticeable bool      `gorm:"type:boolean;not null;default:false"`
	CreatedAt  time.Time `gorm:"precision:6"`
}

// TableName テーブル名
func (unread *Unread) TableName() string {
	return "unreads"
}

// ArchivedMessage 編集前のアーカイブ化されたメッセージの構造体
type ArchivedMessage struct {
	ID        uuid.UUID `gorm:"type:char(36);not null;primary_key"`
	MessageID uuid.UUID `gorm:"type:char(36);not null;index"`
	UserID    uuid.UUID `gorm:"type:char(36);not null"`
	Text      string    `sql:"type:TEXT COLLATE utf8mb4_bin NOT NULL"`
	DateTime  time.Time `gorm:"precision:6"`
}

// TableName ArchivedMessage構造体のテーブル名
func (am *ArchivedMessage) TableName() string {
	return "archived_messages"
}
