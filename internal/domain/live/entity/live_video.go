package entity

type LiveVideo struct {
	id              *VideoId
	title           *VideoTitle
	stremer         *VideoStremer
	thumbnailImage  *ThumbnailImage
	startedDatetime *StartedDatetime
	elapsedTimes    *ElapsedTimes
}

func NewLiveVideo(
	videoId *VideoId,
	videoTitle *VideoTitle,
	videoStremer *VideoStremer,
	thumbnailImage *ThumbnailImage,
	startedDatetime *StartedDatetime,
	elapsedTimes *ElapsedTimes,
) *LiveVideo {
	return &LiveVideo{
		id:              videoId,
		title:           videoTitle,
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
