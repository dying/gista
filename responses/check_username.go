package responses

type CheckUsername struct {
	Response
	Username  string `json:"username,omitempty"`
	Available bool   `json:"available,omitempty"`
	Error     string `json:"error,omitempty"`
	ErrorType string `json:"error_type,omitempty"`
}
