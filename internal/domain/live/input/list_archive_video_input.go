package input

import "github.com/sokorahen-szk/rust-live/internal/domain/live/entity"

type SearchConditions string

const (
	SearchConditionsOR  SearchConditions = "OR"
	SearchConditionsAND SearchConditions = "AND"
)

type ListArchiveVideoInput struct {
	VideoStatuses []entity.VideoStatus

	SearchConditions *SearchConditions
}

func (in ListArchiveVideoInput) GetSearchConditions() string {
	if in.SearchConditions == nil {
		return string(SearchConditionsOR)
	}

	return string(*in.SearchConditions)
}
