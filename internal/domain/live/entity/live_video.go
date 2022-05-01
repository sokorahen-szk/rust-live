package entity

type LiveVideo struct {
	Id              *VideoId         `json:"id"`
	Title           *VideoTitle      `json:"title"`
	Url             *VideoUrl        `json:"url"`
	Stremer         *VideoStremer    `json:"stremer"`
	ThumbnailImage  *ThumbnailImage  `json:"thumbnail_image"`
	StartedDatetime *StartedDatetime `json:"started_datetime"`
	ElapsedTimes    *ElapsedTimes    `json:"ended_datetime"`
}

func NewLiveVideo(
	videoId *VideoId,
	videoTitle *VideoTitle,
	videoUrl *VideoUrl,
	videoStremer *VideoStremer,
	thumbnailImage *ThumbnailImage,
	startedDatetime *StartedDatetime,
	elapsedTimes *ElapsedTimes,
) *LiveVideo {
	return &LiveVideo{
		Id:              videoId,
		Title:           videoTitle,
		Url:             videoUrl,
		Stremer:         videoStremer,
		ThumbnailImage:  thumbnailImage,
		StartedDatetime: startedDatetime,
		ElapsedTimes:    elapsedTimes,
	}
}

func (ins LiveVideo) GetId() *VideoId {
	return ins.Id
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

func (ins LiveVideo) GetThumbnailImage() *ThumbnailImage {
	return ins.ThumbnailImage
}

func (ins LiveVideo) GetStartedDatetime() *StartedDatetime {
	return ins.StartedDatetime
}

func (ins LiveVideo) GetElapsedTimes() *ElapsedTimes {
	return ins.ElapsedTimes
}
