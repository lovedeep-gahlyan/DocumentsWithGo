package controllers

import (
	"docs/models"
	"docs/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	docService *service.DocumentService
}

func NewDocumentHandler(ds *service.DocumentService) *DocumentHandler {
	return &DocumentHandler{docService: ds}
}

func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	var req struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewResponseError("Invalid input", 400))
		return
	}

	document, err := h.docService.CreateDocument(uint(userID), req.Name, req.Content)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, document)
}
