package list

type ListLiveVideoInput struct {
	searchKeywords string
}

func NewListLiveVideoInput(searchKeywords string) *ListLiveVideoInput {
	return &ListLiveVideoInput{
		searchKeywords: searchKeywords,
	}
}

func (ins ListLiveVideoInput) SearchKeywords() string {
	return ins.searchKeywords
}
