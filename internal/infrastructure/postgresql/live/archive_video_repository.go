package live

import (
	"context"

	"github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/input"
	"github.com/sokorahen-szk/rust-live/internal/domain/live/repository"
	"github.com/sokorahen-szk/rust-live/internal/infrastructure/postgresql"
)

type archiveVideoRepository struct {
	conn *postgresql.PostgreSql
}

func NewArchiveVideoRepository(conn *postgresql.PostgreSql) repository.ArchiveVideoRepositoryInterface {
	return &archiveVideoRepository{
		conn: conn,
	}
}

func (repository *archiveVideoRepository) Create(ctx context.Context, in *input.ArchiveVideoInput) error {
	err := repository.conn.Create(in)
	if err != nil {
		return err
	}

	return nil
}
func (repository *archiveVideoRepository) GetByBroadcastId(ctx context.Context, broadcastId *entity.VideoBroadcastId) (*entity.ArchiveVideo, error) {
	achiveVideoInput := input.ArchiveVideoInput{}
	err := repository.conn.Get(&achiveVideoInput, "broadcast_id = ?", broadcastId)
	if err != nil {
		return nil, err
	}

	return repository.scan(achiveVideoInput), nil
}

func (repository *archiveVideoRepository) List(ctx context.Context, listInput *input.ListArchiveVideoInput) ([]*entity.ArchiveVideo, error) {
	achiveVideoInputs := []input.ArchiveVideoInput{}

	postgresqlQuery := postgresql.NewPostgresqlQuery(listInput.GetSearchConditions())

	if len(listInput.VideoStatuses) > 0 {
		postgresqlQuery.Add("status IN ?", listInput.VideoStatuses)
	}

	err := repository.conn.List(
		&achiveVideoInputs,
		postgresqlQuery.GetQueries(),
		postgresqlQuery.GetArgs(),
	)
	if err != nil {
		return nil, err
	}

	return repository.scans(achiveVideoInputs), nil
}

func (repository *archiveVideoRepository) scans(inputs []input.ArchiveVideoInput) []*entity.ArchiveVideo {
	resultArchiveVideos := make([]*entity.ArchiveVideo, 0)
	for _, input := range inputs {
		resultArchiveVideos = append(resultArchiveVideos, repository.scan(input))
	}

	return resultArchiveVideos
}

func (repository *archiveVideoRepository) scan(input input.ArchiveVideoInput) *entity.ArchiveVideo {
	videoId := entity.NewVideoId(input.Id)
	broadcastId := entity.NewVideoBroadcastId(input.BroadcastId)
	videoTitle := entity.NewVideoTitle(input.Title)
	videoStremer := entity.NewVideoStremer(input.Stremer)
	platform := entity.NewPlatformFromInt(input.Platform)
	status := entity.NewVideoStatus(entity.VideoStatusStreaming)
	thumbnailImage := entity.NewThumbnailImage(input.ThumbnailImage)
	startedDatetime := entity.NewStartedDatetimeFromTime(input.StartedDatetime)

	var endedDatetime *entity.EndedDatetime
	if input.EndedDatetime != nil {
		endedDatetime = entity.NewEndedDatetimeFromTime(&input.EndedDatetime.Time)
	}

	var videoUrl *entity.VideoUrl
	if input.Url != nil {
		videoUrl = entity.NewVideoUrl(input.Url.String)
	}

	return entity.NewArchiveVideo(
		videoId,
		broadcastId,
		videoTitle,
		videoUrl,
		videoStremer,
		platform,
		status,
		thumbnailImage,
		startedDatetime,
		endedDatetime,
	)
}
