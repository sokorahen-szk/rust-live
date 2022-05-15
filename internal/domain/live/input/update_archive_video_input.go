package input

import "github.com/sokorahen-szk/rust-live/internal/domain/live/entity"

type UpdateArchiveVideoInput struct {
	Status        *entity.VideoStatus
	EndedDatetime *entity.EndedDatetime
}
