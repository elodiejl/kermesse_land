package controllers

import (
	"back/models"
	"back/repositories"
	"back/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TombolaController struct {
	repo *repositories.TombolaRepository
}

func NewTombolaController(repo *repositories.TombolaRepository) *TombolaController {
	return &TombolaController{repo: repo}
}

// Créer une nouvelle tombola
func (ctrl *TombolaController) CreateTombola(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	var tombola models.Tombola

	if err := c.ShouldBindJSON(&tombola); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.repo.Create(&tombola); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create tombola"})
		return
	}

	c.JSON(http.StatusCreated, tombola)
}

// Récupérer une tombola par ID
func (ctrl *TombolaController) GetTombolaByID(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tombola ID"})
		return
	}

	tombola, err := ctrl.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tombola not found"})
		return
	}

	c.JSON(http.StatusOK, tombola)
}

// Récupérer toutes les tombolas pour une kermesse
func (ctrl *TombolaController) GetTombolasByKermesse(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	kermesseID, err := strconv.Atoi(c.Param("kermesse_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kermesse ID"})
		return
	}

	tombolas, err := ctrl.repo.FindAllByKermesseID(kermesseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve tombolas"})
		return
	}

	c.JSON(http.StatusOK, tombolas)
}

// Supprimer une tombola par ID
func (ctrl *TombolaController) DeleteTombola(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tombola ID"})
		return
	}

	if err := ctrl.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete tombola"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tombola deleted"})
}
