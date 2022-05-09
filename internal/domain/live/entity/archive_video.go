package entity

type ArchiveVideo struct {
	Id              *VideoId
	BroadcastId     *VideoBroadcastId
	Title           *VideoTitle
	Url             *VideoUrl
	Stremer         *VideoStremer
	Platform        *Platform
	Status          *VideoStatus
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
	platform *Platform,
	videoStatus *VideoStatus,
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
		Platform:        platform,
		Status:          videoStatus,
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
	if ins.Url == nil {
		return nil
	}
	return ins.Url
}

func (ins ArchiveVideo) GetStremer() *VideoStremer {
	return ins.Stremer
}

func (ins ArchiveVideo) GetPlatform() *Platform {
	return ins.Platform
}

func (ins ArchiveVideo) GetStatus() *VideoStatus {
	return ins.Status
}

func (ins ArchiveVideo) GetThumbnailImage() *ThumbnailImage {
	return ins.ThumbnailImage
}

func (ins ArchiveVideo) GetStartedDatetime() *StartedDatetime {
	return ins.StartedDatetime
}

func (ins ArchiveVideo) GetEndedDatetime() *EndedDatetime {
	if ins.EndedDatetime == nil {
		return nil
	}
	return ins.EndedDatetime
}
