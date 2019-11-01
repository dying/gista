package responses

import "github.com/dying/gista/models"

type PendingFollowRequests struct {
	Response
	User []models.User `json:"user,omitempty"`
}
