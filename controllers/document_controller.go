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

func (h *DocumentHandler) EditDocument(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	docID, _ := strconv.Atoi(c.Param("doc_id"))

	var req struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.docService.EditDocument(uint(docID), uint(userID), req.Name, req.Content)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	docID, _ := strconv.Atoi(c.Param("doc_id"))

	err := h.docService.DeleteDocument(uint(docID), uint(userID))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}

func (h *DocumentHandler) GrantAccess(c *gin.Context) {
	ownerID, _ := strconv.Atoi(c.Param("user_id"))
	docID, _ := strconv.Atoi(c.Param("doc_id"))
	targetUserID, _ := strconv.Atoi(c.Param("target_user_id"))

	err := h.docService.GrantAccess(uint(docID), uint(ownerID), uint(targetUserID))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access granted successfully"})
}

func (h *DocumentHandler) GetDocument(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	docID, _ := strconv.Atoi(c.Param("doc_id"))

	document, err := h.docService.GetDocument(uint(docID), uint(userID))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, document)
}
