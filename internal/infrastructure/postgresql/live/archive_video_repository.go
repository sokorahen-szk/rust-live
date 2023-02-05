package live

import (
	"context"
	"database/sql"

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
	resultTx := repository.conn.Debug().Create(in)
	if resultTx.Error != nil {
		return resultTx.Error
	}

	return nil
}
func (repository *archiveVideoRepository) GetByBroadcastId(ctx context.Context, broadcastId *entity.VideoBroadcastId) (*entity.ArchiveVideo, error) {
	achiveVideoInput := input.ArchiveVideoInput{}
	resultTx := repository.conn.Debug().Where("broadcast_id = ?", broadcastId).First(&achiveVideoInput)
	if resultTx.Error != nil {
		return nil, resultTx.Error
	}

	return repository.scan(achiveVideoInput), nil
}

func (repository *archiveVideoRepository) List(ctx context.Context, listInput *input.ListArchiveVideoInput) ([]*entity.ArchiveVideo, error) {
	achiveVideoInputs := []input.ArchiveVideoInput{}

	tx := repository.conn.Debug()
	if listInput.GetSearchConditions() == "AND" {
		if len(listInput.VideoStatuses) > 0 {
			tx = repository.conn.Where("status IN ?", listInput.VideoStatuses)
		}

		if len(listInput.BroadcastIds) > 0 {
			tx = repository.conn.Where("broadcast_id IN ?", listInput.BroadcastIds)
		}
	} else {
		if len(listInput.VideoStatuses) > 0 {
			tx = repository.conn.Or("status IN ?", listInput.VideoStatuses)
		}

		if len(listInput.BroadcastIds) > 0 {
			tx = repository.conn.Or("broadcast_id IN ?", listInput.BroadcastIds)
		}
	}

	resultTx := tx.Find(&achiveVideoInputs)
	if resultTx.Error != nil {
		return nil, resultTx.Error
	}

	return repository.scans(achiveVideoInputs), nil
}

func (repository *archiveVideoRepository) Update(ctx context.Context, id *entity.VideoId, updateInput *input.UpdateArchiveVideoInput) (*entity.ArchiveVideo, error) {
	achiveVideoInput := input.ArchiveVideoInput{}

	repository.conn.Debug().First(&achiveVideoInput, id)
	if updateInput.Status != nil {
		achiveVideoInput.Status = updateInput.Status.Int()
	}
	if updateInput.EndedDatetime != nil {
		achiveVideoInput.EndedDatetime = &sql.NullTime{
			Time:  *updateInput.EndedDatetime.Time(),
			Valid: true,
		}
	}

	resultTx := repository.conn.Save(&achiveVideoInput)
	if resultTx.Error != nil {
		return nil, resultTx.Error
	}

	return repository.scan(achiveVideoInput), nil
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
	status := entity.NewVideoStatusFromInt(input.Status)
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
