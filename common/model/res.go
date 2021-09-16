package model

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"message"`
	Data   interface{} `json:"data"`
}
