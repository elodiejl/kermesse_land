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
	repo     repositories.TombolaRepository
	userRepo repositories.UserRepository
}

func NewTombolaController(repo repositories.TombolaRepository) *TombolaController {
	return &TombolaController{repo: repo}
}

// CreateTombola godoc
// @Summary Create a new tombola
// @Description Create a new tombola
// @Tags tombola
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param tombola body models.Tombola true "Tombola object to create"
// @Success 201 {object} models.Tombola "Tombola created"
// @Failure 400 {object} string "Invalid input"
// @Failure 401 {object} string "Unauthorized: No authorization token provided"
// @Failure 500 {object} string "Could not create tombola"
// @Router /tombola [post]
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

	var user models.User
	if err := ctrl.userRepo.GetUserByID(userID, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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

// GetTombolaByID godoc
// @Summary Get tombola by ID
// @Description Retrieve a tombola by its ID
// @Tags tombola
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Tombola ID"
// @Success 200 {object} models.Tombola
// @Failure 400 {object} string "Invalid tombola ID"
// @Failure 401 {object} string "Unauthorized: No authorization token provided"
// @Failure 404 {object} string "Tombola not found"
// @Router /tombola/{id} [get]
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

	var user models.User
	if err := ctrl.userRepo.GetUserByID(userID, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tombola ID"})
		return
	}

	tombola, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tombola not found"})
		return
	}

	c.JSON(http.StatusOK, tombola)
}

// GetTombolasByKermesse godoc
// @Summary Get all tombolas for a kermesse
// @Description Retrieve all tombolas for a specific kermesse
// @Tags tombola
// @Produce json
// @Security ApiKeyAuth
// @Param kermesse_id path int true "Kermesse ID"
// @Success 200 {array} models.Tombola
// @Failure 400 {object} string "Invalid kermesse ID"
// @Failure 401 {object} string "Unauthorized: No authorization token provided"
// @Failure 500 {object} string "Could not retrieve tombolas"
// @Router /kermesse/{kermesse_id}/tombolas [get]
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

	var user models.User
	if err := ctrl.userRepo.GetUserByID(userID, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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

// DeleteTombola godoc
// @Summary Delete a tombola by ID
// @Description Delete a specific tombola by its ID
// @Tags tombola
// @Security ApiKeyAuth
// @Param id path int true "Tombola ID"
// @Success 200 {string} string "Tombola deleted successfully"
// @Failure 400 {object} string "Invalid tombola ID"
// @Failure 401 {object} string "Unauthorized: No authorization token provided"
// @Failure 404 {object} string "Tombola not found"
// @Failure 500 {object} string "Could not delete tombola"
// @Router /tombola/{id} [delete]
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

	var user models.User
	if err := ctrl.userRepo.GetUserByID(userID, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tombola ID"})
		return
	}

	if err := ctrl.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete tombola"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tombola deleted"})
}
