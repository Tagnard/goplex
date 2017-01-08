package filters

type Filter string

const (
	All            Filter = "all"
	Unwatched      Filter = "unwatched"
	Newest         Filter = "newest"
	RecentlyAdded  Filter = "recentlyAdded"
	RecentlyViewed Filter = "recentlyViewed"
	OnDeck         Filter = "onDeck"
)
