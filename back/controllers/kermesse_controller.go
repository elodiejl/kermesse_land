package controllers

import (
	"back/models"
	"back/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KermesseController struct {
	repo *repositories.KermesseRepository
}

func NewKermesseController(repo *repositories.KermesseRepository) *KermesseController {
	return &KermesseController{repo: repo}
}

// @Summary Créer une nouvelle kermesse
// @Description Crée une kermesse.
// @Tags Kermesses
// @Accept  json
// @Produce  json
// @Param ticket body models.Kermesse true "Kermesse à créer"
// @Success 201 {object} models.Kermesse
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Bad request"
// @Router /kermesses [post]
func (ctrl *KermesseController) CreateKermesse(c *gin.Context) {
	var kermesse models.Kermesse

	if err := c.ShouldBindJSON(&kermesse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.repo.Create(&kermesse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create kermesse"})
		return
	}

	c.JSON(http.StatusCreated, kermesse)
}

// Récupérer une kermesse par ID
func (ctrl *KermesseController) GetKermesseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kermesse ID"})
		return
	}

	kermesse, err := ctrl.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kermesse not found"})
		return
	}

	c.JSON(http.StatusOK, kermesse)
}

// Récupérer toutes les kermesses
func (ctrl *KermesseController) GetAllKermesses(c *gin.Context) {
	kermesses, err := ctrl.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve kermesses"})
		return
	}

	c.JSON(http.StatusOK, kermesses)
}

// Supprimer une kermesse par ID
func (ctrl *KermesseController) DeleteKermesse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kermesse ID"})
		return
	}

	if err := ctrl.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete kermesse"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kermesse deleted"})
}
