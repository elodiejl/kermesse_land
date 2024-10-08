package controllers

import (
	"back/models"
	"back/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StandController struct {
	repo repositories.StandRepository
}

func NewStandController(repo repositories.StandRepository) *StandController {
	return &StandController{repo: repo}
}

// CreateStand godoc
// @Summary Créer un nouveau stand
// @Description Crée un stand pour une kermesse
// @Tags stands
// @Accept json
// @Produce json
// @Param stand body models.Stand true "Stand à créer"
// @Success 201 {object} models.Stand "Stand créé"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Could not create stand"
// @Router /stands [post]
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

// GetStandByID godoc
// @Summary Récupérer un stand par ID
// @Description Récupère un stand spécifique par son ID
// @Tags stands
// @Produce json
// @Param id path int true "Stand ID"
// @Success 200 {object} models.Stand "Stand trouvé"
// @Failure 400 {object} string "Invalid stand ID"
// @Failure 404 {object} string "Stand not found"
// @Router /stands/detail/{id} [get]
func (ctrl *StandController) GetStandByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stand ID"})
		return
	}

	stand, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stand not found"})
		return
	}

	c.JSON(http.StatusOK, stand)
}

// GetStandsByKermesse godoc
// @Summary Récupérer tous les stands pour une kermesse
// @Description Récupère tous les stands associés à une kermesse spécifique
// @Tags stands
// @Produce json
// @Param kermesse_id path int true "Kermesse ID"
// @Success 200 {array} models.Stand "Liste des stands"
// @Failure 400 {object} string "Invalid kermesse ID"
// @Failure 500 {object} string "Could not retrieve stands"
// @Router /kermesse/{kermesse_id}/stands [get]
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

// DeleteStand godoc
// @Summary Supprimer un stand par ID
// @Description Supprime un stand spécifique par son ID
// @Tags stands
// @Security ApiKeyAuth
// @Param id path int true "Stand ID"
// @Success 200 {object} string "Stand supprimé"
// @Failure 400 {object} string "Invalid stand ID"
// @Failure 404 {object} string "Stand not found"
// @Failure 500 {object} string "Could not delete stand"
// @Router /stands/delete/{id} [delete]
func (ctrl *StandController) DeleteStand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stand ID"})
		return
	}

	if err := ctrl.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete stand"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stand deleted"})
}
