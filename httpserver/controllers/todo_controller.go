package controllers

import (
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/services"
	"final-project-2/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	CreateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
}

type todoController struct {
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) *todoController {
	return &todoController{
		todoService,
	}
}

// CreateTodo godoc
// @Tags    Todo
// @Summary create a todo
// @Param   todo body     dto.CreateTodoDto true "Create Todo DTO"
// @Success 200  {object} utils.HttpSuccess[any]
// @Failure 400  {object} utils.HttpError
// @Failure 500  {object} utils.HttpError
// @Router  /todo [post]
func (c *todoController) CreateTodo(ctx *gin.Context) {
	var dto dto.CreateTodoDto
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	todo, err := c.todoService.CreateTodo(dto)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess("Todo Created", todo))
}

// DeleteTodo godoc
// @Tags    Todo
// @Summary create a todo
// @Param   id  path     int true "Todo ID"
// @Success 200 {object} utils.HttpSuccess[models.UserModel]
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router  /todo/{id} [delete]
func (c *todoController) DeleteTodo(ctx *gin.Context) {
	idDto := ctx.Param("id")
	id, err := strconv.Atoi(idDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewHttpError("Bad Request", err.Error()))
		return
	}

	err = c.todoService.DeleteTodo(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewHttpError("Internal Server Error", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewHttpSuccess[any]("Todo Deleted", nil))
}

// UpdateTodo godoc
// @Tags    Todo
// @Summary update a todo
// @Param   id  path     int true "Todo ID"
// @Success 200 {object} utils.HttpSuccess[models.UserModel]
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router  /todo/update/{id} [put]
func (c *todoController) UpdateTodo(ctx *gin.Context) {
	respData := &utils.ResponseData{
		Status: "Fail",
	}

	var param dto.UpdateTodoDto

	todo_id := ctx.Param("id")
	id, _ := strconv.Atoi(todo_id)

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		respData.Message = map[string]string{
			"error":  err.Error(),
			"status": "error binding json",
		}
		ctx.JSON(http.StatusInternalServerError, respData)
		return
	}

	msg, err := c.todoService.UpdateTodo(int64(id), param)
	if err != nil {
		respData.Message = map[string]string{
			"error":  err.Error(),
			"status": "error update data",
		}

		ctx.JSON(http.StatusBadRequest, msg)
		return
	}

	respData = &utils.ResponseData{
		Status:  http.StatusOK,
		Message: msg,
		Details: nil,
	}

	ctx.JSON(http.StatusOK, respData)

}
