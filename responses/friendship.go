package responses

import "github.com/dying/gista/models"

type Friendship struct {
	Response
	FriendshipStatus models.FriendshipStatus `json:"friendship_status"`
}
