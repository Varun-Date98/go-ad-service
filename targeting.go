package main

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/Varun-Date98/go-ad-service/internal/database"
)


func selectAd(user UserContext, placement PlacementContext, candidates []database.GetCandidateAdRow) *AdDecision {
	var best *AdDecision

	// Basic ad targeting for v0
	for _, c := range candidates {
		if user.Country != c.Country {
			continue
		}

		if user.Age < int(c.MinAge) || user.Age > int(c.MaxAge) {
			continue
		}

		if len(c.PlacementIds) > 0 && !stringInSlice(placement.PlacementID, c.PlacementIds) {
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

		bidCpm, err := strconv.ParseFloat(c.BidCpm, 64)

		if err != nil {
			continue
		}

		interestMatches := countInterestMatches(c.InterestsAny, user.Interests)
		matchingFactor := 1.0 + 0.1 * float64(interestMatches)
		score := bidCpm * matchingFactor + (rand.Float64() * 0.01)

		if best == nil || best.Score < score {
			decision := AdDecision{
				CampaignID: c.CampaignID,
				CreativeID: c.CreativeID,
				AssetURL: c.AssetUrl,
				ClickURL: c.ClickUrl,
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