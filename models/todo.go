package models

type Todo struct {
	ID     int `json:"id"`
	Todo   string `json:"todo"`
	Date   string `json:"date"`
}

type TodoRequest struct {
    Todo string `json:"todo"`
}