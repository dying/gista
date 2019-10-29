package responses

import "github.com/dying/gista/models"

type Comment struct {
	Response
	Comment *models.Comment `json:"comment"`
}
