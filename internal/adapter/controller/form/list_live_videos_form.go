package form

import (
	pb "github.com/sokorahen-szk/rust-live/api/proto"
)

type ListLiveVideosForm struct {
	SearchKeywords string                        `validate:"omitempty,max=20"`
	Sort           pb.ListLiveVideosRequest_Sort `validate:"omitempty,is_list_live_video_sort"`
	Page           int32                         `validate:"required,gte=1"`
	Limit          int32                         `validate:"omitempty,gte=1"`
}

func NewListLiveVideosForm(req *pb.ListLiveVideosRequest) ListLiveVideosForm {
	return ListLiveVideosForm{
		SearchKeywords: req.GetSearchKeywords(),
		Sort:           req.GetSort(),
		Page:           req.GetPage(),
		Limit:          req.GetLimit(),
	}
}

func (ins ListLiveVideosForm) GetSearchKeywords() string {
	return ins.SearchKeywords
}

func (ins ListLiveVideosForm) GetSort() int {
	return int(ins.Sort.Number())
}

func (ins ListLiveVideosForm) GetPage() int {
	return int(ins.Page)
}

func (ins ListLiveVideosForm) GetLimit() int {
	return int(ins.Limit)
}
