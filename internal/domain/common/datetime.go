package common

import "time"

const (
	DateFormat = "2006-01-02"
)

type Datetime struct {
	time *time.Time
}

func NewDatetime(value string) *Datetime {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}

	return &Datetime{time: &t}
}

func (ins *Datetime) RFC3339() string {
	return ins.time.Format(time.RFC3339)
}

func (ins *Datetime) YYYYMMDD() string {
	return ins.time.Format(DateFormat)
}

func (ins *Datetime) Time() *time.Time {
	return ins.time
}
