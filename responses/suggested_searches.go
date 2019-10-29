package responses

import "github.com/dying/gista/models"

type SuggestedSearches struct {
	Response
	Suggested []models.Suggested `json:"suggested"`
	RankToken string             `json:"rank_token"`
}
