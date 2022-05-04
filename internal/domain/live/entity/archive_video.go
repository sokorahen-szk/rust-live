package entity

type ArchiveVideo struct {
	Id              *VideoId
	BroadcastId     *VideoBroadcastId
	Title           *VideoTitle
	Url             *VideoUrl
	Stremer         *VideoStremer
	ThumbnailImage  *ThumbnailImage
	StartedDatetime *StartedDatetime
	EndedDatetime   *EndedDatetime
}

func NewArchiveVideo(
	videoId *VideoId,
	broadcastId *VideoBroadcastId,
	videoTitle *VideoTitle,
	videoUrl *VideoUrl,
	videoStremer *VideoStremer,
	thumbnailImage *ThumbnailImage,
	startedDatetime *StartedDatetime,
	endedDatetime *EndedDatetime,
) *ArchiveVideo {
	return &ArchiveVideo{
		Id:              videoId,
		BroadcastId:     broadcastId,
		Title:           videoTitle,
		Url:             videoUrl,
		Stremer:         videoStremer,
		ThumbnailImage:  thumbnailImage,
		StartedDatetime: startedDatetime,
		EndedDatetime:   endedDatetime,
	}
}

func (ins ArchiveVideo) GetId() *VideoId {
	return ins.Id
}

func (ins ArchiveVideo) GetBroadcastId() *VideoBroadcastId {
	return ins.BroadcastId
}

func (ins ArchiveVideo) GetTitle() *VideoTitle {
	return ins.Title
}

func (ins ArchiveVideo) GetUrl() *VideoUrl {
	return ins.Url
}

func (ins ArchiveVideo) GetStremer() *VideoStremer {
	return ins.Stremer
}

func (ins ArchiveVideo) GetThumbnailImage() *ThumbnailImage {
	return ins.ThumbnailImage
}

func (ins ArchiveVideo) GetStartedDatetime() *StartedDatetime {
	return ins.StartedDatetime
}

func (ins ArchiveVideo) GetEndedDatetime() *EndedDatetime {
	return ins.EndedDatetime
}
