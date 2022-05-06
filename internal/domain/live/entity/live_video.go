package entity

type LiveVideo struct {
	Id              *VideoId          `json:"id"`
	BroadcastId     *VideoBroadcastId `json:"broadcast_id"`
	Title           *VideoTitle       `json:"title"`
	Url             *VideoUrl         `json:"url"`
	Stremer         *VideoStremer     `json:"stremer"`
	Viewer          *VideoViewer      `json:"viewer"`
	ThumbnailImage  *ThumbnailImage   `json:"thumbnail_image"`
	StartedDatetime *StartedDatetime  `json:"started_datetime"`
	ElapsedTimes    *ElapsedTimes     `json:"elapsed_times"`
}

func NewLiveVideo(
	videoId *VideoId,
	broadcastId *VideoBroadcastId,
	videoTitle *VideoTitle,
	videoUrl *VideoUrl,
	videoStremer *VideoStremer,
	videoViewer *VideoViewer,
	thumbnailImage *ThumbnailImage,
	startedDatetime *StartedDatetime,
	elapsedTimes *ElapsedTimes,
) *LiveVideo {
	return &LiveVideo{
		Id:              videoId,
		BroadcastId:     broadcastId,
		Title:           videoTitle,
		Url:             videoUrl,
		Stremer:         videoStremer,
		Viewer:          videoViewer,
		ThumbnailImage:  thumbnailImage,
		StartedDatetime: startedDatetime,
		ElapsedTimes:    elapsedTimes,
	}
}

func (ins LiveVideo) GetId() *VideoId {
	return ins.Id
}

func (ins LiveVideo) GetBroadcastId() *VideoBroadcastId {
	return ins.BroadcastId
}

func (ins LiveVideo) GetTitle() *VideoTitle {
	return ins.Title
}

func (ins LiveVideo) GetUrl() *VideoUrl {
	return ins.Url
}

func (ins LiveVideo) GetStremer() *VideoStremer {
	return ins.Stremer
}

func (ins LiveVideo) GetViewer() *VideoViewer {
	return ins.Viewer
}

func (ins LiveVideo) GetThumbnailImage() *ThumbnailImage {
	return ins.ThumbnailImage
}

func (ins LiveVideo) GetStartedDatetime() *StartedDatetime {
	return ins.StartedDatetime
}

func (ins LiveVideo) GetElapsedTimes() *ElapsedTimes {
	return ins.ElapsedTimes
}
