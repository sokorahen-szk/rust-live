package entity

import (
	"time"

	"github.com/sokorahen-szk/rust-live/internal/domain/common"
)

type StartedDatetime struct {
	*common.Datetime
}

func NewStartedDatetimeFromTime(time *time.Time) *StartedDatetime {
	datetime := common.NewDatetimeFromTime(time)
	m := StartedDatetime{datetime}
	return &m
}

func NewStartedDatetime(value string) *StartedDatetime {
	datetime := common.NewDatetime(value)
	m := StartedDatetime{datetime}
	return &m
}
