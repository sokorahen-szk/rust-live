package input

import (
	"database/sql"
	"time"
)

type ArchiveVideoInput struct {
	Id              int    `gorm:"primaryKey"`
	BroadcastId     string `gorm:"index"`
	Title           string
	Url             string
	Stremer         string
	ThumbnailImage  string
	StartedDatetime *time.Time
	EndedDatetime   *sql.NullTime
}

func (m *ArchiveVideoInput) TableName() string {
	return "archive_videos"
}
