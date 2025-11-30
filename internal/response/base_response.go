package response

// BaseResponse adalah struktur respons global
type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // nullable
}
