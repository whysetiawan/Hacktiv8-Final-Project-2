package dto

type CreateTodoDto struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTodoDto struct {
	Status bool `json:"status" binding:"required"`
}
