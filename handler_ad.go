package main

import (
	"fmt"
	"net/http"

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
	

	decision := selectAd(userCtx, placementCtx, adCandidatesCtx)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	if decision == nil {
		// No eligible campaign for this user/context
		w.WriteHeader(http.StatusNoContent) // 204
		return
	}

	type AdResponce struct {
		CampaignID string `json:"campaign_id"`
		CreativeID string `json:"creative_id"`
		AssetUrl string `json:"asset_url"`
		ClickUrl string `json:"click_url"`
	}

	respondWithJSON(w, 200, AdResponce{
		CampaignID: decision.CampaignID,
		CreativeID: decision.CreativeID,
		AssetUrl: decision.AssetURL,
		ClickUrl: decision.ClickURL,
	})
}
