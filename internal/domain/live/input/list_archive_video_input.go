package input

type SearchConditions string

const (
	SearchConditionsOR  SearchConditions = "OR"
	SearchConditionsAND SearchConditions = "AND"
)

type ListArchiveVideoInput struct {
	VideoStatuses []int

	SearchConditions *SearchConditions
}

func (in ListArchiveVideoInput) GetSearchConditions() string {
	if in.SearchConditions == nil {
		return string(SearchConditionsOR)
	}

	return string(*in.SearchConditions)
}
