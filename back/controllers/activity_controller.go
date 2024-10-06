package controllers

import (
	"back/models"
	"back/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	repo *repositories.ActivityRepositoryImpl
}

// Créer un nouveau contrôleur d'activités
func NewActivityController(repo *repositories.ActivityRepositoryImpl) *ActivityController {
	return &ActivityController{repo: repo}
}

// Créer une nouvelle activité
func (ctrl *ActivityController) CreateActivity(c *gin.Context) {
	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.repo.CreateActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create activity"})
		return
	}

	c.JSON(http.StatusCreated, activity)
}

// Récupérer une activité par ID
func (ctrl *ActivityController) GetActivityByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	activity, err := ctrl.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// Supprimer une activité par ID
func (ctrl *ActivityController) DeleteActivity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	if err := ctrl.repo.DeleteActivity(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted"})
}

// Récupérer toutes les activités pour un stand
func (ctrl *ActivityController) GetActivitiesByStandID(c *gin.Context) {
	standID, err := strconv.Atoi(c.Param("stand_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stand ID"})
		return
	}

	activities, err := ctrl.repo.FindAllByStandID(standID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve activities"})
		return
	}

	c.JSON(http.StatusOK, activities)
}
