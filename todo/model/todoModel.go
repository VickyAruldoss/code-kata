package model

type Todo struct {
	UserId      int    `json:"userId"`
	Id          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"completed"`
}
