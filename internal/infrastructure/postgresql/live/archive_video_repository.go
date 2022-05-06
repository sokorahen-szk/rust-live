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
	achiveVideoInput := &input.ArchiveVideoInput{}
	err := repository.conn.Get(achiveVideoInput, "broadcast_id = ?", broadcastId)
	if err != nil {
		return nil, err
	}

	return repository.scan(achiveVideoInput), nil
}

func (repository *archiveVideoRepository) List(ctx context.Context) ([]*entity.ArchiveVideo, error) {
	// TODO: 開発する
	return nil, nil
}

func (repository *archiveVideoRepository) scan(in *input.ArchiveVideoInput) *entity.ArchiveVideo {

	videoId := entity.NewVideoId(in.Id)
	broadcastId := entity.NewVideoBroadcastId(in.BroadcastId)
	videoTitle := entity.NewVideoTitle(in.Title)
	videoStremer := entity.NewVideoStremer(in.Stremer)
	platform := entity.NewPlatformFromInt(in.Platform)
	thumbnailImage := entity.NewThumbnailImage(in.ThumbnailImage)
	startedDatetime := entity.NewStartedDatetimeFromTime(in.StartedDatetime)

	var endedDatetime *entity.EndedDatetime
	if in.EndedDatetime != nil {
		endedDatetime = entity.NewEndedDatetimeFromTime(&in.EndedDatetime.Time)
	}

	var videoUrl *entity.VideoUrl
	if in.Url != nil {
		videoUrl = entity.NewVideoUrl(in.Url.String)
	}

	return entity.NewArchiveVideo(
		videoId,
		broadcastId,
		videoTitle,
		videoUrl,
		videoStremer,
		platform,
		thumbnailImage,
		startedDatetime,
		endedDatetime,
	)
}
