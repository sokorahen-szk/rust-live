package entity

import (
	"github.com/sokorahen-szk/rust-live/internal/domain/common"
)

type EndedDatetime struct {
	*common.Datetime
}

func NewEndedDatetime(value string) *EndedDatetime {
	datetime := common.NewDatetime(value)
	m := EndedDatetime{datetime}
	return &m
}
