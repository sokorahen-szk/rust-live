package form

import (
	pb "github.com/sokorahen-szk/rust-live/api/proto"
)

type ListLiveVideosForm struct {
	SearchKeywords string `validate:"max=20"`
}

func NewListLiveVideosForm(req *pb.ListLiveVideosRequest) ListLiveVideosForm {
	return ListLiveVideosForm{
		SearchKeywords: req.GetSearchKeywords(),
	}
}
