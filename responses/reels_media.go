package responses

import (
	"github.com/dying/gista/models"
)

type ReelsMedia struct {
	Response
	ReelsMedia *[]models.Reel          `json:"reels_media,omitempty"`
	Reels      *map[string]models.Reel `json:"reels,omitempty"`
}
