package entity

type VideoStatus int

const (
	VideoStatusStreaming VideoStatus = 1
	VideoStatusEnded     VideoStatus = 2
)

var videoStatusValues = map[VideoStatus]string{
	VideoStatusStreaming: "Streaming",
	VideoStatusEnded:     "Ended",
}

func NewVideoStatus(v VideoStatus) *VideoStatus {
	return &v
}

func NewVideoStatusFromInt(value int) *VideoStatus {
	for v, _ := range videoStatusValues {
		if int(v) == value {
			return &v
		}
	}
	return nil
}

func (ins VideoStatus) Int() int {
	return int(ins)
}

func (v VideoStatus) String() string {
	s := videoStatusValues[v]
	if s == "" {
		panic("videoStatusの情報が有効ではありません。")
	}
	return s
}
