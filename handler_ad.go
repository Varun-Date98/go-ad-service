package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Varun-Date98/go-ad-service/internal/database"
)


func (db *dbAPI) adHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	userId := q.Get("user_id")
	user, err := db.DB.GetUserById(r.Context(), userId)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting user with id %v, %v", userId, err))
	}

	placementId := q.Get("placement_id")
	creatorId := q.Get("creator_id")

	userCtx := UserContext{
		UserID: user.UserID,
		Age: int(user.Age),
		Country: user.Country,
		Device: user.Device,
		Language: user.Language,
		Interests: user.Interests,
	}

	placementCtx := PlacementContext{
		PlacementID: placementId,
		CreatorID: creatorId,
	}

	adCandidatesCtx, err := db.DB.GetCandidateAd(r.Context(), database.GetCandidateAdParams{
		Country: user.Country,
		Age: user.Age,
		PlacementID: placementId,
		Device: user.Device,
		Language: user.Language,
		CreatorID: creatorId,
		Interests: user.Interests,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Could not get active campaigns %v", err))
		return
	}

	allow := func(campaignId string) bool {return true}

	if db.Redis != nil {
		allow = func(campaignId string) bool {
			key := fmt.Sprintf("cap:%s:%s", user.UserID, campaignId)
			ok, err := db.Redis.Exists(r.Context(), key).Result()

			if err != nil {
				return true
			}

			return ok == 0
		}
	}

	decision := selectAd(userCtx, placementCtx, adCandidatesCtx, allow)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	if decision == nil {
		// No eligible campaign for this user/context
		w.WriteHeader(http.StatusNoContent) // 204
		return
	}

	if db.Redis != nil {
		key := fmt.Sprintf("cap:%s:%s", user.UserID, decision.CampaignID)
		_, err := db.Redis.Set(r.Context(), key, 1, time.Minute).Result()

		if err != nil {
			log.Panicf("Could not save ad to Redis %v", err)
		}
	}

	type AdResponse struct {
		CampaignID string `json:"campaign_id"`
		CreativeID string `json:"creative_id"`
		AssetUrl string `json:"asset_url"`
		ClickUrl string `json:"click_url"`
	}

	respondWithJSON(w, 200, AdResponse{
		CampaignID: decision.CampaignID,
		CreativeID: decision.CreativeID,
		AssetUrl: decision.AssetURL,
		ClickUrl: decision.ClickURL,
	})
}
