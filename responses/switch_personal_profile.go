package responses

import "github.com/dying/gista/models"

type SwitchPersonalProfile struct {
	Response
	Users []models.User `json:"users,omitempty"`
}
