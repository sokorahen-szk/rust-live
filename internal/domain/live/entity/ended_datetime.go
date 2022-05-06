package entity

import (
	"encoding/json"
	"strings"
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

func (ins *EndedDatetime) UnmarshalJSON(data []byte) error {
	s := strings.ReplaceAll(string(data), "\"", "")
	datetime := common.NewDatetime(s)
	*ins = EndedDatetime{datetime}
	return nil
}

func (ins EndedDatetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ins.RFC3339())
}
