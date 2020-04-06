package response

// Response struct
type Response struct {
	StatusCode uint        `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
