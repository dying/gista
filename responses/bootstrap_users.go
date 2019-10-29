package responses

import "github.com/dying/gista/models"

type BootstrapUsers struct {
	Response
	Surfaces []models.Surface `json:"surfaces"`
	Users    []models.User    `json:"users"`
}
