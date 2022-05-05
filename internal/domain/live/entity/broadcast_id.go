package entity

type VideoBroadcastId string

func NewVideoBroadcastId(value string) *VideoBroadcastId {
	if value == "" {
		panic("入力された値が無効")
	}

	m := VideoBroadcastId(value)
	return &m
}

func (ins VideoBroadcastId) String() string {
	return string(ins)
}
