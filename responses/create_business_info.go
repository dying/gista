package responses

import "github.com/dying/gista/models"

type CreateBusinessInfo struct {
	Response
	Users []models.User `json:"users,omitempty"`
}
