package models

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

type JwtToken struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}
