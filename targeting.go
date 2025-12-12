package main

import (
	"math/rand"
)

var campaigns []Campaign

func seedCampaigns() {
	campaigns = []Campaign{
		{
			CampaignID: "cmp_riot",
			MinAge:     18,
			MaxAge:     35,
			Country:    "US",
			BidCPM:     4.0,
			Creatives: []Creative{
				{
					CreativeID: "crt_riot_1",
					AssetURL:   "https://cdn.example.com/riot-15s.mp4",
					ClickURL:   "https://riotgames.com/signup",
				},
			},
		},
		{
			CampaignID: "cmp_snack",
			MinAge:     16,
			MaxAge:     45,
			Country:    "US",
			BidCPM:     2.0,
			Creatives: []Creative{
				{
					CreativeID: "crt_snack_1",
					AssetURL:   "https://cdn.example.com/snack-30s.mp4",
					ClickURL:   "https://snackbrand.com",
				},
			},
		},
	}
}


func selectAd(user UserContext, placement PlacementContext) *AdDecision {
	var best *AdDecision

	// Basic ad targeting for v0
	for _, c := range campaigns {
		if user.Country != c.Country {
			continue
		}

		if user.Age < c.MinAge || user.Age > c.MaxAge {
			continue
		}

		if len(c.Creatives) == 0 {
			continue
		}

		// Simple score: BidCPM + a little random jitter so ties break
		score := c.BidCPM + (rand.Float64() * 0.1)

		if best == nil || score > best.Score {
			creative := c.Creatives[0]
			decision := AdDecision{
				CampaignID: c.CampaignID,
				CreativeID: creative.CreativeID,
				AssetURL:   creative.AssetURL,
				ClickURL:   creative.ClickURL,
				Score:      score,
			}
			best = &decision
		}
	}

	return best
}
