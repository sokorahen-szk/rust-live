package form

import (
	pb "github.com/sokorahen-szk/rust-live/api/proto"
)

type ListLiveVideosForm struct {
	SearchKeywords string                        `validate:"omitempty,max=20"`
	Platform       []pb.VideoPlatform            `validate:"omitempty,dive,is_video_platform"`
	Sort           pb.ListLiveVideosRequest_Sort `validate:"omitempty,is_list_live_video_sort"`
	Page           int32                         `validate:"omitempty,gte=1"`
	Limit          int32                         `validate:"omitempty,gte=1"`
}

func NewListLiveVideosForm(req *pb.ListLiveVideosRequest) ListLiveVideosForm {
	return ListLiveVideosForm{
		SearchKeywords: req.GetSearchKeywords(),
		Sort:           req.GetSort(),
		Platform:       req.GetVideoPlatforms(),
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

func (ins ListLiveVideosForm) GetPlatforms() []int {
	platforms := make([]int, 0)
	for _, p := range ins.Platform {
		platforms = append(platforms, int(p.Number()))
	}

	return platforms
}

func (ins ListLiveVideosForm) GetPage() int {
	return int(ins.Page)
}

func (ins ListLiveVideosForm) GetLimit() int {
	return int(ins.Limit)
}
