package model

type Todo struct {
	ID     string `json:"id"`
	Todo   string `json:"todo"`
	Date   string `json:"date"`
}

type TodoRequest struct {
    Todo string `json:"todo"`
}