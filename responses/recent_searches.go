package responses

import "github.com/dying/gista/models"

type RecentSearches struct {
	Response
	Recent []models.Suggested `json:"recent"`
}
