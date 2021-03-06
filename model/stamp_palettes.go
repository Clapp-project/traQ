package model

import (
	"database/sql/driver"
	"errors"
	"github.com/gofrs/uuid"
	"strings"
	"time"
)

type UUIDs []uuid.UUID

func (arr UUIDs) Value() (driver.Value, error) {
	idStr := []string{}
	for _, id := range arr {
		idStr = append(idStr, id.String())
	}
	return strings.Join(idStr, ","), nil
}

func (arr *UUIDs) Scan(src interface{}) error {
	switch s := src.(type) {
	case nil:
		*arr = UUIDs{}
	case string:
		idSlice := strings.Split(s, ",")
		for _, value := range idSlice {
			ID, err := uuid.FromString(value)
			if err != nil {
				continue
			}
			*arr = append(*arr, ID)
		}
	case []byte:
		str := string(s)
		idSlice := strings.Split(str, ",")
		for _, value := range idSlice {
			ID, err := uuid.FromString(value)
			if err != nil {
				continue
			}
			*arr = append(*arr, ID)
		}
	default:
		return errors.New("failed to scan UUIDs")
	}
	return nil
}

func (arr UUIDs) ToUUIDSlice() ([]uuid.UUID) {
	uuidSlice := []uuid.UUID{}
	for _, id := range arr {
		uuidSlice = append(uuidSlice, id)
	}
	return uuidSlice
}

type StampPalette struct {
	ID          uuid.UUID `gorm:"type:char(36);not null;primary_key" json:"id"`
	Name        string    `gorm:"type:varchar(30);not null" json:"name"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Stamps      UUIDs     `gorm:"type:text;not null" json:"stamps"`
	CreatorID   uuid.UUID `gorm:"type:char(36);not null;index" json:"creatorId"`
	CreatedAt   time.Time `gorm:"precision:6" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"precision:6" json:"updatedAt"`
}

// TableName StampPalettes構造体のテーブル名
func (*StampPalette) TableName() string {
	return "stamp_palettes"
}
