package responses

import "github.com/dying/gista/models"

type TokenResult struct {
	Response
	Token models.Token
}
