package controllers

import (
	"back/models"
	"back/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityParticipationController struct {
	repo *repositories.ActivityParticipationRepositoryImpl
}

// Créer un nouveau contrôleur pour les participations à des activités
func NewActivityParticipationController(repo *repositories.ActivityParticipationRepositoryImpl) *ActivityParticipationController {
	return &ActivityParticipationController{repo: repo}
}

// Créer une nouvelle participation
func (ctrl *ActivityParticipationController) CreateParticipation(c *gin.Context) {
	var participation models.ActivityParticipation
	if err := c.ShouldBindJSON(&participation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.repo.CreateParticipation(&participation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create participation"})
		return
	}

	c.JSON(http.StatusCreated, participation)
}

// Récupérer une participation par ID
func (ctrl *ActivityParticipationController) GetParticipationByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid participation ID"})
		return
	}

	participation, err := ctrl.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Participation not found"})
		return
	}

	c.JSON(http.StatusOK, participation)
}

// Récupérer toutes les participations d'un utilisateur
func (ctrl *ActivityParticipationController) GetParticipationsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	participations, err := ctrl.repo.FindAllByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve participations"})
		return
	}

	c.JSON(http.StatusOK, participations)
}

// Supprimer une participation par ID
func (ctrl *ActivityParticipationController) DeleteParticipation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid participation ID"})
		return
	}

	if err := ctrl.repo.DeleteParticipation(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete participation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Participation deleted"})
}
