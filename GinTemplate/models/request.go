package models

type Claims struct {
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Role     []int64 `json:"role"`
}

type Paging struct {
	Page int64 `json:"page,omitempty"`
	Size int64 `json:"size,omitempty"`
}
