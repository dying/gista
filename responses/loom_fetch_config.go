package responses

import "github.com/dying/gista/models"

type LoomFetchConfig struct {
	Response
	SystemControl models.SystemControl `json:"system_control"`
	TraceControl  models.TraceControl  `json:"trace_control"`
	Id            int                  `json:"id"`
}
