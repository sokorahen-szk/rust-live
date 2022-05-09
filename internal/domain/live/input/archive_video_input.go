package input

import (
	"database/sql"
	"time"
)

type ArchiveVideoInput struct {
	Id              int    `gorm:"primaryKey"`
	BroadcastId     string `gorm:"index"`
	Title           string
	Url             *sql.NullString
	Stremer         string
	Platform        int
	Status          int
	ThumbnailImage  string
	StartedDatetime *time.Time
	EndedDatetime   *sql.NullTime
}

func (m *ArchiveVideoInput) TableName() string {
	return "archive_videos"
}
