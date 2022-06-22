package live

type AnalyticsOutput struct {
	Viewers int
}

func NewAnalyticsOutput(viewers int) *AnalyticsOutput {
	return &AnalyticsOutput{
		Viewers: viewers,
	}
}

func (out *AnalyticsOutput) GetViewers() int {
	return out.Viewers
}
