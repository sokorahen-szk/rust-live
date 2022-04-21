package list

type ListLiveVideosInput struct {
	searchKeywords string
}

func NewListLiveVideosInput(searchKeywords string) *ListLiveVideosInput {
	return &ListLiveVideosInput{
		searchKeywords: searchKeywords,
	}
}

func (ins ListLiveVideosInput) SearchKeywords() string {
	return ins.searchKeywords
}
