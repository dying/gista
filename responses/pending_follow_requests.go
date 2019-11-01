package responses

import "github.com/dying/gista/models"

type PendingFollowRequests struct {
	Response
	Users []models.User `json:"users,omitempty"`
}
