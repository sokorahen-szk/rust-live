package entity

import (
	"github.com/sokorahen-szk/rust-live/internal/domain/common"
)

type StartedDatetime struct {
	*common.Datetime
}

func NewStartedDatetime(value string) *StartedDatetime {
	datetime := common.NewDatetime(value)
	m := StartedDatetime{datetime}
	return &m
}
