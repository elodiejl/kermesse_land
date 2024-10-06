package controllers

import (
	"back/models"
	"back/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ParentController struct {
	repo repositories.ParentRepository
}

// NewParentController crée une nouvelle instance de ParentController
func NewParentController(repo repositories.ParentRepository) *ParentController {
	return &ParentController{repo: repo}
}

// CreateParent gère la création d'un parent
// @Summary Créer un parent
// @Description Créer une nouvelle entrée parent
// @Tags parents
// @Accept json
// @Produce json
// @Param parent body models.Parent true "Parent data"
// @Success 201 {object} models.Parent "Parent created"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /parents [post]
func (ctrl *ParentController) CreateParent(c *gin.Context) {
	var parent models.Parent

	if err := c.ShouldBindJSON(&parent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := ctrl.repo.Create(&parent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create parent", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, parent)
}

// GetParentByID gère la récupération d'un parent par ID
// @Summary Trouver un parent par ID
// @Description Récupérer un parent par son ID
// @Tags parents
// @Produce json
// @Param id path int true "Parent ID"
// @Success 200 {object} models.Parent "Parent found"
// @Failure 400 {object} string "Invalid ID format"
// @Failure 404 {object} string "Parent not found"
// @Failure 500 {object} string "Internal server error"
// @Router /parents/{id} [get]
func (ctrl *ParentController) GetParentByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	parent, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		return
	}

	c.JSON(http.StatusOK, parent)
}

// UpdateParent gère la mise à jour des informations d'un parent
// @Summary Mettre à jour un parent
// @Description Mettre à jour les informations d'un parent
// @Tags parents
// @Accept json
// @Produce json
// @Param id path int true "Parent ID"
// @Param parent body models.Parent true "Parent data"
// @Success 200 {object} models.Parent "Parent updated"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "Parent not found"
// @Failure 500 {object} string "Internal server error"
// @Router /parents/{id} [put]
func (ctrl *ParentController) UpdateParent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var parent models.Parent
	if err := c.ShouldBindJSON(&parent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	existingParent, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		return
	}

	// Mettre à jour les champs du parent
	existingParent.TokensAmount = parent.TokensAmount

	if err := ctrl.repo.Update(existingParent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update parent", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingParent)
}

// DeleteParent gère la suppression d'un parent
// @Summary Supprimer un parent
// @Description Supprimer un parent par son ID
// @Tags parents
// @Produce json
// @Param id path int true "Parent ID"
// @Success 200 {string} string "Parent deleted"
// @Failure 400 {object} string "Invalid ID format"
// @Failure 404 {object} string "Parent not found"
// @Failure 500 {object} string "Internal server error"
// @Router /parents/{id} [delete]
func (ctrl *ParentController) DeleteParent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	_, err = ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		return
	}

	if err := ctrl.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete parent", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Parent deleted"})
}
