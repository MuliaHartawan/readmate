package response

type ResponseDTO struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Body       any    `json:"body"`
}
