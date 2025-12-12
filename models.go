package main

// Data models for ad service

type Campaign struct {
	CampaignID string
	MinAge     int
	MaxAge     int
	Country    string
	BidCPM     float64
	Creatives  []Creative
}

type Creative struct {
	CreativeID string
	AssetURL   string
	ClickURL   string
}

type UserContext struct {
	UserID  string
	Age     int
	Country string
}

type PlacementContext struct {
	PlacementID string
}

// Return type
type AdDecision struct {
	CampaignID string  `json:"campaign_id"`
	CreativeID string  `json:"creative_id"`
	AssetURL   string  `json:"asset_url"`
	ClickURL   string  `json:"click_url"`
	Score      float64 `json:"score"`
}
