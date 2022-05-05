package entity

import (
	"time"

	"github.com/sokorahen-szk/rust-live/internal/domain/common"
)

type EndedDatetime struct {
	*common.Datetime
}

func NewEndedDatetimeFromTime(time *time.Time) *EndedDatetime {
	datetime := common.NewDatetimeFromTime(time)
	m := EndedDatetime{datetime}
	return &m
}

func NewEndedDatetime(value string) *EndedDatetime {
	datetime := common.NewDatetime(value)
	m := EndedDatetime{datetime}
	return &m
}
