package entity

type LiveVideo struct {
	id              *VideoId
	title           *VideoTitle
	url             *VideoUrl
	stremer         *VideoStremer
	thumbnailImage  *ThumbnailImage
	startedDatetime *StartedDatetime
	elapsedTimes    *ElapsedTimes
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
		id:              videoId,
		title:           videoTitle,
		url:             videoUrl,
		stremer:         videoStremer,
		thumbnailImage:  thumbnailImage,
		startedDatetime: startedDatetime,
		elapsedTimes:    elapsedTimes,
	}
}

func (ins LiveVideo) Id() *VideoId {
	return ins.id
}

func (ins LiveVideo) Title() *VideoTitle {
	return ins.title
}

func (ins LiveVideo) Url() *VideoUrl {
	return ins.url
}

func (ins LiveVideo) Stremer() *VideoStremer {
	return ins.stremer
}

func (ins LiveVideo) ThumbnailImage() *ThumbnailImage {
	return ins.thumbnailImage
}

func (ins LiveVideo) StartedDatetime() *StartedDatetime {
	return ins.startedDatetime
}

func (ins LiveVideo) ElapsedTimes() *ElapsedTimes {
	return ins.elapsedTimes
}
