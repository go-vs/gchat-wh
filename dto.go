package gchat_wh

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ErrorResponse struct {
	Error *Error `json:"error"`
}
