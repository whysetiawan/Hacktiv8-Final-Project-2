package dto

type CreateTodoDto struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTodoDto struct {
	Status bool `json:"status" binding:"required"`
}

type RegisterDto struct {
	Age      uint8  `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
