package controllers

import (
	"back/models"
	"back/repositories"
	"back/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	repo repositories.TicketRepository
}

// NewTicketController Créer un nouveau contrôleur
func NewTicketController(repo repositories.TicketRepository) *TicketController {
	return &TicketController{repo: repo}
}

// CreateTicket godoc
// @Summary Créer un nouveau ticket
// @Description Crée un billet de tombola pour un élève ou un parent.
// @Tags Tickets
// @Accept json
// @Produce json
// @Param ticket body models.Ticket true "Ticket à créer"
// @Success 201 {object} models.Ticket "Ticket créé avec succès"
// @Failure 400 {object} string "Requête invalide"
// @Failure 500 {object} string "Erreur interne du serveur"
// @Router /tickets [post]
func (ctrl *TicketController) CreateTicket(c *gin.Context) {
	var ticket models.Ticket

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Générer un numéro de ticket unique
	ticket.TicketNumber = services.GenerateTicketNumber()
	ticket.PurchasedAt = services.GetCurrentTime()

	if err := ctrl.repo.Create(&ticket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create ticket"})
		return
	}

	c.JSON(http.StatusCreated, ticket)
}

// GetTicketByID godoc
// @Summary Récupérer un ticket par ID
// @Description Récupère un ticket spécifique par son ID.
// @Tags Tickets
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 200 {object} models.Ticket "Ticket trouvé"
// @Failure 400 {object} string "ID de ticket invalide"
// @Failure 404 {object} string "Ticket non trouvé"
// @Router /tickets/{id} [get]
func (ctrl *TicketController) GetTicketByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	ticket, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// GetTicketsByTombola godoc
// @Summary Récupérer tous les tickets pour une tombola
// @Description Récupère tous les tickets associés à une tombola spécifique.
// @Tags Tickets
// @Produce json
// @Param tombola_id path int true "Tombola ID"
// @Success 200 {array} models.Ticket "Liste des tickets"
// @Failure 400 {object} string "ID de tombola invalide"
// @Failure 500 {object} string "Erreur lors de la récupération des tickets"
// @Router /tombola/{tombola_id}/tickets [get]
func (ctrl *TicketController) GetTicketsByTombola(c *gin.Context) {
	tombolaID, err := strconv.Atoi(c.Param("tombola_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tombola ID"})
		return
	}

	tickets, err := ctrl.repo.FindAllByTombolaID(tombolaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve tickets"})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// DeleteTicket godoc
// @Summary Supprimer un ticket par ID
// @Description Supprime un ticket spécifique par son ID.
// @Tags Tickets
// @Security ApiKeyAuth
// @Param id path int true "Ticket ID"
// @Success 200 {object} string "Ticket supprimé avec succès"
// @Failure 400 {object} string "ID de ticket invalide"
// @Failure 404 {object} string "Ticket non trouvé"
// @Failure 500 {object} string "Erreur lors de la suppression du ticket"
// @Router /tickets/{id} [delete]
func (ctrl *TicketController) DeleteTicket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	if err := ctrl.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}
