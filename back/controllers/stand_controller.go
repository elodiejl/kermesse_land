package controllers

import (
	"back/models"
	"back/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StandController struct {
	repo *repositories.StandRepository
}

func NewStandController(repo *repositories.StandRepository) *StandController {
	return &StandController{repo: repo}
}

// Créer un nouveau stand
func (ctrl *StandController) CreateStand(c *gin.Context) {
	var stand models.Stand

	if err := c.ShouldBindJSON(&stand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.repo.Create(&stand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create stand"})
		return
	}

	c.JSON(http.StatusCreated, stand)
}

// Récupérer un stand par ID
func (ctrl *StandController) GetStandByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stand ID"})
		return
	}

	stand, err := ctrl.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stand not found"})
		return
	}

	c.JSON(http.StatusOK, stand)
}

// Récupérer tous les stands pour une kermesse
func (ctrl *StandController) GetStandsByKermesse(c *gin.Context) {
	kermesseID, err := strconv.Atoi(c.Param("kermesse_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kermesse ID"})
		return
	}

	stands, err := ctrl.repo.FindAllByKermesseID(kermesseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve stands"})
		return
	}

	c.JSON(http.StatusOK, stands)
}

// Supprimer un stand
func (ctrl *StandController) DeleteStand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stand ID"})
		return
	}

	if err := ctrl.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete stand"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stand deleted"})
}
