package controllers

import (
	"docs/models"
	"docs/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponseError("Invalid input", 400))
		return
	}

	user, err := h.userService.CreateUser(req.Username)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))

	user, err := h.userService.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
