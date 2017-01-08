package models

type SectionMetrics struct {
	Id               int
	Total            int
	Unwatched        int
	Watched          int
	WatchedPercent   float64
	UnwatchedPercent float64
}

func NewSectionMetrics(id int, total, unwatched int) *SectionMetrics {
	watched := total - unwatched
	watched_percent := (float64(watched) / float64(total)) * 100
	unwatched_percent := (float64(unwatched) / float64(total)) * 100

	return &SectionMetrics{
		Id: id,
		Total: total,
		Unwatched: unwatched,
		Watched: total - unwatched,
		WatchedPercent: watched_percent,
		UnwatchedPercent: unwatched_percent,
	}
}