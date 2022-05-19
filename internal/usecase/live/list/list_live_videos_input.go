package list

import "github.com/sokorahen-szk/rust-live/internal/domain/live/entity"

type ListLiveVideoInput struct {
	searchKeywords string
	platforms      []*entity.Platform
	sortKey        *entity.LiveVideoSortKey
	page           int
	limit          int
}

func NewListLiveVideoInput(searchKeywords string, platforms []*entity.Platform,
	sortKey *entity.LiveVideoSortKey, page int, limit int) *ListLiveVideoInput {
	return &ListLiveVideoInput{
		searchKeywords: searchKeywords,
		platforms:      platforms,
		sortKey:        sortKey,
		page:           page,
		limit:          limit,
	}
}

func (ins ListLiveVideoInput) SearchKeywords() string {
	return ins.searchKeywords
}

func (ins ListLiveVideoInput) Platforms() []*entity.Platform {
	if ins.platforms == nil {
		return nil
	}

	return ins.platforms
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
