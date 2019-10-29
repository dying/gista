package responses

import "github.com/dying/gista/models"

type UserInfo struct {
	Response
	Megaphone *interface{} `json:"megaphone,omitempty"`
	User      *models.User `json:"user,omitempty"`
}
