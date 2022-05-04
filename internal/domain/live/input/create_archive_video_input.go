package input

import (
	"database/sql"
	"time"
)

type CreateArchiveVideoInput struct {
	Id              uint   `gorm:"primaryKey"`
	BroadcastId     string `gorm:"index"`
	Title           string
	Url             string
	Stremer         string
	ThumbnailImage  string
	StartedDatetime *time.Time
	EndedDatetime   sql.NullTime
}

func (m *CreateArchiveVideoInput) TableName() string {
	return "archive_videos"
}
