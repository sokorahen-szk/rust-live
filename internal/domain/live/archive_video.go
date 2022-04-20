package live

type ArchiveVideo struct {
	id *ArchiveVideoId
}

func NewArchiveVideo(archiveVideoId *ArchiveVideoId) *ArchiveVideo {
	return &ArchiveVideo{
		id: archiveVideoId,
	}
}

func (ins ArchiveVideo) Id() *ArchiveVideoId {
	return ins.id
}
