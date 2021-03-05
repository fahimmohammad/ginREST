package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerInterface interface {
	createTodoHandler(c *gin.Context)
	readAllTodoHandler(c *gin.Context)
	readSingleTodoHandler(c *gin.Context)
	deleteTodoHandler(c *gin.Context)
	updateTodoHandler(c *gin.Context)
}

type handlerFields struct {
	todoService *Service
}

// MakeHTTPHandlers - initializes routing
func MakeHTTPHandlers(router *gin.RouterGroup, todoService *Service) {
	h := handlerFields{
		todoService: todoService,
	}

	router.POST("todo", h.createTodoHandler)
	router.GET("todo", h.readAllTodoHandler)
	router.GET("todo/:id", h.readSingleTodoHandler)
	router.DELETE("todo/:id", h.deleteTodoHandler)
	router.PUT("todo/:id", h.updateTodoHandler)
}

type createTodoRequest struct {
	Todo Todo
}

type createTodoResponse struct {
	Todo Todo   `json:"todo"`
	Err  string `json:"err"`
}

func (h *handlerFields) createTodoHandler(c *gin.Context) {
	var req createTodoRequest

	if err := c.ShouldBindJSON(&req.Todo); err != nil {
		c.JSON(http.StatusInternalServerError, createTodoResponse{
			Todo: Todo{},
			Err:  "Error in data binding",
		})
		return
	}

	postResult, err := h.todoService.PostTodo(req.Todo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, createTodoResponse{
			Todo: Todo{},
			Err:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, createTodoResponse{
		Todo: postResult,
		Err:  "",
	})
	return
}

type readAllTodoResponse struct {
	Todo []Todo `json:"todos"`
	Err  string `json:"err"`
}

func (h *handlerFields) readAllTodoHandler(c *gin.Context) {

	getResult, err := h.todoService.GetAllTodo()

	if err != nil {
		c.JSON(http.StatusInternalServerError, readAllTodoResponse{
			Todo: []Todo{},
			Err:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, readAllTodoResponse{
		Todo: getResult,
		Err:  "",
	})
	return
}

type readSingleTodoResponse struct {
	Todo Todo   `json:"todo"`
	Err  string `json:"err"`
}

func (h *handlerFields) readSingleTodoHandler(c *gin.Context) {
	todoID := c.Param("id")
	getResult, err := h.todoService.GetSingleTodo(todoID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, readSingleTodoResponse{
			Todo: Todo{},
			Err:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, readSingleTodoResponse{
		Todo: getResult,
		Err:  "",
	})
	return
}

type deleteTodoResponse struct {
	Err string `json:"err"`
}

func (h *handlerFields) deleteTodoHandler(c *gin.Context) {
	todoID := c.Param("id")

	err := h.todoService.RemoveTodo(todoID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, deleteTodoResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, deleteTodoResponse{
		Err: "",
	})
	return
}

type updateTodoRequest struct {
	Todo Todo
}

type updateTodoResponse struct {
	Todo Todo   `json:"todo"`
	Err  string `json:"err"`
}

func (h *handlerFields) updateTodoHandler(c *gin.Context) {
	var req updateTodoRequest
	todoID := c.Param("id")
	if err := c.ShouldBindJSON(&req.Todo); err != nil {
		c.JSON(http.StatusInternalServerError, updateTodoResponse{
			Todo: Todo{},
			Err:  "Error in data binding",
		})
		return
	}

	updateResult, err := h.todoService.EditTodo(todoID, req.Todo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, updateTodoResponse{
			Todo: Todo{},
			Err:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updateTodoResponse{
		Todo: updateResult,
		Err:  "",
	})
	return
}
