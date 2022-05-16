package list

import "github.com/sokorahen-szk/rust-live/internal/domain/live/entity"

type ListLiveVideoInput struct {
	searchKeywords string
	sortKey        *entity.LiveVideoSortKey
	page           int
	limit          int
}

func NewListLiveVideoInput(searchKeywords string, sortKey *entity.LiveVideoSortKey, page int, limit int) *ListLiveVideoInput {
	return &ListLiveVideoInput{
		searchKeywords: searchKeywords,
		sortKey:        sortKey,
		page:           page,
		limit:          limit,
	}
}

func (ins ListLiveVideoInput) SearchKeywords() string {
	return ins.searchKeywords
}

func (ins ListLiveVideoInput) SortKey() *entity.LiveVideoSortKey {
	if ins.sortKey == nil {
		return nil
	}
	return ins.sortKey
}

func (ins ListLiveVideoInput) Page() int {
	return ins.page
}

func (ins ListLiveVideoInput) Limit() int {
	return ins.limit
}
