package common

import (
	"time"
)

const (
	DateFormat      string = "2006-01-02"
	TimestampFormat string = "2006-01-02 15:04:05"
)

type Datetime struct {
	time *time.Time
}

func NewDatetimeFromTime(time *time.Time) *Datetime {
	return &Datetime{time: time}
}

func NewDatetime(value string) *Datetime {
	var tm time.Time
	if len(value) == 19 {
		t, err := time.Parse(TimestampFormat, value)
		if err != nil {
			panic(err)
		}

		tm = t
	} else {
		t, err := time.Parse(time.RFC3339, value)
		if err != nil {
			panic(err)
		}

		tm = t
	}

	return &Datetime{time: &tm}
}

func (ins *Datetime) RFC3339() string {
	return ins.time.Format(time.RFC3339)
}

func (ins *Datetime) YYYYMMDD() string {
	return ins.time.Format(DateFormat)
}

func (ins *Datetime) Timestamp() string {
	return ins.time.Format(TimestampFormat)
}

func (ins *Datetime) Time() *time.Time {
	return ins.time
}

func (ins *Datetime) TimeFunc() func() time.Time {
	return func() time.Time {
		return *ins.Time()
	}
}

func (ins *Datetime) DiffSeconds(tm *time.Time) int {
	diff := ins.Time().Sub(*tm)
	return int(diff.Seconds())
}

// a < b
func (ins *Datetime) Lt(tm *time.Time) bool {
	return ins.Time().Before(*tm)
}

// a > b
func (ins *Datetime) Gt(tm *time.Time) bool {
	return ins.Time().After(*tm)
}

// a = b
func (ins *Datetime) Equal(tm *time.Time) bool {
	return ins.Time().Equal(*tm)
}
