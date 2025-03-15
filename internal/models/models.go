package models

type Task struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
