package input

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type CreateArchiveVideoInput struct {
	//Id              uint   `gorm:"primaryKey"`
	BroadcastId     string `gorm:"index"`
	Title           string
	Url             string
	Stremer         string
	ThumbnailImage  string
	StartedDatetime *time.Time
	EndedDatetime   sql.NullTime
	//CreatedAt       time.Time
	//UpdatedAt       time.Time
	gorm.Model
}

func (m *CreateArchiveVideoInput) TableName() string {
	return "archive_videos"
}
