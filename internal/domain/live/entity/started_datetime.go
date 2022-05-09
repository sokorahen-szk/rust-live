package entity

import (
	"encoding/json"
	"strings"
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

func (ins *StartedDatetime) UnmarshalJSON(data []byte) error {
	s := strings.ReplaceAll(string(data), "\"", "")
	datetime := common.NewDatetime(s)
	*ins = StartedDatetime{datetime}
	return nil
}

func (ins StartedDatetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ins.RFC3339())
}
