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
	repo *repositories.TicketRepository
}

// NewTicketController Créer un nouveau contrôleur
func NewTicketController(repo *repositories.TicketRepository) *TicketController {
	return &TicketController{repo: repo}
}

// @Summary Créer un nouveau ticket
// @Description Crée un billet de tombola pour un élève ou un parent.
// @Tags Tickets
// @Accept  json
// @Produce  json
// @Param ticket body models.Ticket true "Ticket à créer"
// @Success 201 {object} models.Ticket
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Bad request"
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

// Récupérer un ticket par ID
func (ctrl *TicketController) GetTicketByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	ticket, err := ctrl.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// GetTicketsByTombola Récupérer tous les tickets pour une tombola
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

// DeleteTicket Supprimer un ticket par ID
func (ctrl *TicketController) DeleteTicket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	if err := ctrl.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete ticket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}
