package main

// Data models for ad service

type Campaign struct {
	CampaignID   string
	MinAge       int
	MaxAge       int
	Country      string
	BidCPM       float64
	PlacementIDs []string // which placements this campaign supports
	InterestsAny []string // match at least one
	DevicesAny   []string // "desktop", "mobile"
	LanguagesAny []string // "en", "de", etc.
	CreatorsAny  []string // optional: specific channels
	Creatives    []Creative
}

type Creative struct {
	CreativeID string
	AssetURL   string
	ClickURL   string
}

type UserContext struct {
	UserID    string
	Age       int
	Country   string
	Device    string   // "desktop" or "mobile"
	Language  string   // "en", "de", etc.
	Interests []string // ["valorant", "esports", etc]
}

type PlacementContext struct {
	PlacementID string // e.g. "stream_pre_roll"
	CreatorID   string // which streamer/channel
}

// Return type
type AdDecision struct {
	CampaignID string  `json:"campaign_id"`
	CreativeID string  `json:"creative_id"`
	AssetURL   string  `json:"asset_url"`
	ClickURL   string  `json:"click_url"`
	Score      float64 `json:"score"`
}
