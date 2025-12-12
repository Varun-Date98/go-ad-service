package main

import (
	"math/rand"
	"strings"
)

var campaigns []Campaign

func seedCampaigns() {
	campaigns = []Campaign{
		{
			CampaignID:   "cmp_riot",
			MinAge:       18,
			MaxAge:       35,
			Country:      "US",
			BidCPM:       4.0,
			PlacementIDs: []string{"stream_pre_roll", "homepage_banner"},
			InterestsAny: []string{"valorant", "league", "esports"},
			DevicesAny:   []string{"desktop", "mobile"},
			LanguagesAny: []string{"en"},
			CreatorsAny:  []string{"shroud", "tenz"},
			Creatives: []Creative{
				{
					CreativeID: "crt_riot_1",
					AssetURL:   "https://cdn.example.com/riot-15s.mp4",
					ClickURL:   "https://riotgames.com/signup",
				},
			},
		},
		{
			CampaignID:   "cmp_snack",
			MinAge:       16,
			MaxAge:       45,
			Country:      "US",
			BidCPM:       2.0,
			PlacementIDs: []string{"stream_mid_roll"},
			InterestsAny: []string{"variety", "just chatting"},
			DevicesAny:   []string{"desktop"},
			LanguagesAny: []string{"en"},
			CreatorsAny:  []string{}, // any creator
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

		if len(c.PlacementIDs) > 0 && !stringInSlice(placement.PlacementID, c.PlacementIDs) {
			continue
		}

		if len(c.DevicesAny) > 0 && !stringInSlice(user.Device, c.DevicesAny) {
			continue
		}

		if len(c.LanguagesAny) > 0 && !stringInSlice(user.Language, c.LanguagesAny) {
			continue
		}

		if len(c.CreatorsAny) > 0 && !stringInSlice(placement.CreatorID, c.CreatorsAny) {
			continue
		}

		if len(c.Creatives) == 0 {
			continue
		}

		interestMatches := countInterestMatches(c.InterestsAny, user.Interests)
		matchingFactor := 1.0 + 0.1 * float64(interestMatches)
		score := c.BidCPM * matchingFactor + (rand.Float64() * 0.01)

		if best == nil || best.Score < score {
			creative := c.Creatives[0]
			decision := AdDecision{
				CampaignID: c.CampaignID,
				CreativeID: creative.CreativeID,
				AssetURL: creative.AssetURL,
				ClickURL: creative.ClickURL,
				Score: score,
			}

			best = &decision
		}
	}

	return best
}


func stringInSlice(s string, list []string) bool {
	for _, v := range list {
		if strings.EqualFold(v, s) {
			return true
		}
	}
	return false
}


func countInterestMatches(adInterests, userInterests []string) int {
	count := 0
	set := make(map[string]struct{})

	for _, v := range adInterests {
		set[strings.ToLower(v)] = struct {}{}
	}

	for _, v := range userInterests {
		if _, ok := set[strings.ToLower(v)]; ok {
			count++
		}
	}

	return count
}