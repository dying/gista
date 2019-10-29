package responses

import "github.com/dying/gista/models"

type FriendshipsShow struct {
	Response
	models.FriendshipStatus
}
