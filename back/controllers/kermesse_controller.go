package controllers

import (
	"back/models"
	"back/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KermesseController struct {
	repo repositories.KermesseRepository
}

// NewKermesseController Retourne une nouvelle instance de KermesseController
func NewKermesseController(repo repositories.KermesseRepository) *KermesseController {
	return &KermesseController{repo: repo}
}

// @Summary Créer une nouvelle kermesse
// @Description Crée une kermesse.
// @Tags Kermesses
// @Accept  json
// @Produce  json
// @Param kermesse body models.Kermesse true "Kermesse à créer"
// @Success 201 {object} models.Kermesse
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Could not create kermesse"
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

// @Summary Récupérer une kermesse par ID
// @Description Récupère une kermesse par son ID.
// @Tags Kermesses
// @Produce  json
// @Param id path int true "Kermesse ID"
// @Success 200 {object} models.Kermesse
// @Failure 400 {object} string "Invalid kermesse ID"
// @Failure 404 {object} string "Kermesse not found"
// @Router /kermesses/{id} [get]
func (ctrl *KermesseController) GetKermesseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kermesse ID"})
		return
	}

	kermesse, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kermesse not found"})
		return
	}

	c.JSON(http.StatusOK, kermesse)
}

// @Summary Récupérer toutes les kermesses
// @Description Récupère toutes les kermesses.
// @Tags Kermesses
// @Produce  json
// @Success 200 {array} models.Kermesse
// @Failure 500 {object} string "Could not retrieve kermesses"
// @Router /kermesses [get]
func (ctrl *KermesseController) GetAllKermesses(c *gin.Context) {
	kermesses, err := ctrl.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve kermesses"})
		return
	}

	c.JSON(http.StatusOK, kermesses)
}

// @Summary Supprimer une kermesse par ID
// @Description Supprime une kermesse par son ID.
// @Tags Kermesses
// @Produce  json
// @Param id path int true "Kermesse ID"
// @Success 200 {object} string "Kermesse deleted"
// @Failure 400 {object} string "Invalid kermesse ID"
// @Failure 500 {object} string "Could not delete kermesse"
// @Router /kermesses/{id} [delete]
func (ctrl *KermesseController) DeleteKermesse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kermesse ID"})
		return
	}

	if err := ctrl.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete kermesse"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kermesse deleted"})
}
