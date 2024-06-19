package v1

type RespMessage struct {
	Data    any    `json:"data"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
