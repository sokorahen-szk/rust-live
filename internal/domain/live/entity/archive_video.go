package entity

type ArchiveVideo struct {
	id              *VideoId
	title           *VideoTitle
	stremer         *VideoStremer
	thumbnailImage  *ThumbnailImage
	startedDatetime *StartedDatetime
	endedDatetime   *EndedDatetime
}

func NewArchiveVideo(
	videoId *VideoId,
	videoTitle *VideoTitle,
	videoStremer *VideoStremer,
	thumbnailImage *ThumbnailImage,
	startedDatetime *StartedDatetime,
	endedDatetime *EndedDatetime,
) *ArchiveVideo {
	return &ArchiveVideo{
		id:              videoId,
		title:           videoTitle,
		stremer:         videoStremer,
		thumbnailImage:  thumbnailImage,
		startedDatetime: startedDatetime,
		endedDatetime:   endedDatetime,
	}
}

func (ins ArchiveVideo) Id() *VideoId {
	return ins.id
}

func (ins ArchiveVideo) Title() *VideoTitle {
	return ins.title
}

func (ins ArchiveVideo) Stremer() *VideoStremer {
	return ins.stremer
}

func (ins ArchiveVideo) ThumbnailImage() *ThumbnailImage {
	return ins.thumbnailImage
}

func (ins ArchiveVideo) StartedDatetime() *StartedDatetime {
	return ins.startedDatetime
}

func (ins ArchiveVideo) EndedDatetime() *EndedDatetime {
	return ins.endedDatetime
}
