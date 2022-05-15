package mockInput

import (
	"database/sql"
	"fmt"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
)

func NewMockArchiveVideoInput(i int, url string, platform *entity.Platform,
	status *entity.VideoStatus, startedDatetime *entity.StartedDatetime,
	endedDatetime *entity.EndedDatetime) *input.ArchiveVideoInput {
	var inUrl *sql.NullString
	var inEndedDatetime *sql.NullTime
	if url != "" {
		inUrl = &sql.NullString{
			String: url,
			Valid:  true,
		}
	}
	if endedDatetime != nil {
		inEndedDatetime = &sql.NullTime{
			Time:  *endedDatetime.Time(),
			Valid: true,
		}
	}

	return &input.ArchiveVideoInput{
		Id:              i,
		BroadcastId:     fmt.Sprintf("%010d", i),
		Title:           fmt.Sprintf("配信%d", i),
		Url:             inUrl,
		Stremer:         fmt.Sprintf("配信者%d", i),
		Platform:        platform.Int(),
		Status:          status.Int(),
		ThumbnailImage:  fmt.Sprintf("https://example.com/test%d.jpg", i),
		StartedDatetime: startedDatetime.Time(),
		EndedDatetime:   inEndedDatetime,
	}
}
