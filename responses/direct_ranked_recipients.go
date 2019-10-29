package responses

import "github.com/dying/gista/models"

type DirectRankedRecipients struct {
	Response
	Expires          int                            `json:"expires"`
	RankedRecipients []models.DirectRankedRecipient `json:"ranked_recipients"`
	Filtered         bool                           `json:"filtered"`
	RequestId        string                         `json:"request_id"`
	RankToken        string                         `json:"rank_token"`
}
