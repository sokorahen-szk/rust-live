package entity

const (
	PlatformTwitch  Platform = 1
	PlatformYoutube Platform = 2
)

type Platform int

var platformValues = map[Platform]string{
	PlatformTwitch:  "Twitch",
	PlatformYoutube: "Youtube",
}

func NewPlatform(p Platform) *Platform {
	return &p
}

func NewPlatformFromInt(n int) *Platform {
	for p, _ := range platformValues {
		if int(p) == n {
			return &p
		}
	}

	return nil
}

func (p Platform) Int() int {
	return int(p)
}

func (p Platform) String() string {
	s := platformValues[p]
	if s == "" {
		panic("platformの情報が有効ではありません。")
	}
	return s
}
