package responses

import "github.com/dying/gista/models"

type Sync struct {
	Response
	//Experiments []map[string]interface{} `json:"experiments"`
	Experiments []models.Experiment `json:"experiments"`
}
