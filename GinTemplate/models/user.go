package models

type User struct {
	Name     string  `json:"name"`
	Password string  `json:"password,omitempty"`
	Role     []int64 `json:"role"`
	Status   bool    `json:"status"`
	LastTime int64   `json:"lasttime,omitempty"`
}
