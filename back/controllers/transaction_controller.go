package controllers

import (
	"back/models"
	"back/repositories"
	"back/services"
	"net/http"
	"os"
	"strconv"
	_ "strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"gorm.io/gorm"

	"github.com/stripe/stripe-go/v72/paymentintent"
)

// TransactionController gère les transactions
type TransactionController struct {
	repo repositories.TransactionRepository
}

// NewTransactionController crée une nouvelle instance de TransactionController
func NewTransactionController(repo repositories.TransactionRepository) *TransactionController {
	return &TransactionController{repo: repo}
}

// CreateTransaction crée une nouvelle transaction
// @Summary      Créer une transaction
// @Description  Crée une nouvelle transaction pour un parent
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        request body models.Transaction true "Données de la requête"
// @Success      201 {object} models.Transaction
// @Failure      400 {object} map[string]string
// @Router       /transactions [post]
func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := tc.repo.Create(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// GetTransaction récupère une transaction par ID
// @Summary      Obtenir une transaction
// @Description  Récupère une transaction par son ID
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        id path int true "Transaction ID"
// @Success      200 {object} models.Transaction
// @Failure      404 {object} map[string]string
// @Router       /transactions/{id} [get]
func (tc *TransactionController) GetTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	transaction, err := tc.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// UpdateTransaction met à jour une transaction
// @Summary      Mettre à jour une transaction
// @Description  Met à jour les informations d'une transaction
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        id path int true "Transaction ID"
// @Param        request body models.Transaction true "Données de la requête"
// @Success      200 {object} models.Transaction
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /transactions/{id} [put]
func (tc *TransactionController) UpdateTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	transaction.ID = uint(id) // Assigner l'ID à la transaction
	if err := tc.repo.Update(&transaction); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction supprime une transaction par ID
// @Summary      Supprimer une transaction
// @Description  Supprime une transaction par son ID
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        id path int true "Transaction ID"
// @Success      204 {object} nil
// @Failure      404 {object} map[string]string
// @Router       /transactions/{id} [delete]
func (tc *TransactionController) DeleteTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No authorization token provided"})
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	_, err = services.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	if err := tc.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetTransactionsByParentID récupère toutes les transactions d'un parent par ID
// @Summary      Obtenir toutes les transactions d'un parent
// @Description  Récupère toutes les transactions associées à un parent
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        parent_id query int true "Parent ID"
// @Success      200 {array} models.Transaction
// @Failure      404 {object} map[string]string
// @Router       /transactions [get]
func (tc *TransactionController) GetTransactionsByParentID(c *gin.Context) {
	parentID, err := strconv.Atoi(c.Query("parent_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent ID"})
		return
	}
	transactions, err := tc.repo.FindByParentID(parentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No transactions found"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

type PaymentIntentRequest struct {
	Amount   int    `json:"amount" binding:"required"`
	Currency string `json:"currency" binding:"required"`
	ParentID string `json:"parent_id" binding:"required"`
}

// CreatePaymentIntent gère la création d'un PaymentIntent Stripe
// @Summary      Créer un PaymentIntent
// @Description  Crée un PaymentIntent pour le paiement de jetons
// @Tags         payments
// @Accept       json
// @Produce      json
//
//	@Param       request body PaymentIntentRequest true "Données de la requête"
//
// @Success      200 {object} map[string]string "Client secret"
// @Failure      400 {object} map[string]string "Invalid request"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /api/create-payment-intent [post]
func CreatePaymentIntent(c *gin.Context) {
	var requestData PaymentIntentRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")
	if stripe.Key == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Stripe API key is not set"})
		return
	}

	// Créer le PaymentIntent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(requestData.Amount)),
		Currency: stripe.String(requestData.Currency),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating payment intent"})
		return
	}

	// Envoyer le clientSecret au frontend
	response := map[string]string{"clientSecret": pi.ClientSecret}
	c.JSON(http.StatusOK, response)

	/*parentID, err := strconv.Atoi(requestData.ParentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent ID"})
		return
	}

	completePurchaseData := struct {
		ParentID     int `json:"parent_id"`
		TokensAmount int `json:"tokens_amount"`
		Price        int `json:"price"`
	}{
		ParentID:     parentID,
		TokensAmount: requestData.Amount,
		Price:        requestData.Amount,
	}

	completePurchase(c, struct {
		ParentID     int
		TokensAmount int
		Price        int
	}(completePurchaseData))*/
}

// CompletePurchaseHandler gère la requête pour compléter un achat
func (tc *TransactionController) CompletePurchaseHandler(c *gin.Context) {
	var requestData struct {
		ParentID     int `json:"parent_id"`
		TokensAmount int `json:"tokens_amount"`
		Price        int `json:"price"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	completePurchase(c, struct {
		ParentID     int
		TokensAmount int
		Price        int
	}(requestData))
}

// completePurchase gère la validation de l'achat de jetons
// @Summary      Compléter l'achat de jetons
// @Description  Enregistre la transaction et met à jour le nombre de jetons du parent
// @Tags         payments
// @Accept       json
// @Produce      json
// @Param       request body struct { ParentID int `json:"parent_id"`; TokensAmount int `json:"tokens_amount"`; Price int `json:"price"` } true "Données de la requête"
// @Success      200 {object} map[string]string "Transaction completed successfully"
// @Failure      400 {object} map[string]string "Invalid request"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /api/complete-purchase [post]
func completePurchase(c *gin.Context, requestData struct {
	ParentID     int
	TokensAmount int
	Price        int
}) {
	db := c.MustGet("db").(*gorm.DB)

	transaction := models.Transaction{
		ParentID:        requestData.ParentID,
		Price:           requestData.Price,
		TokensAmount:    requestData.TokensAmount,
		TransactionDate: time.Now().Format(time.RFC3339),
	}

	db.Create(&transaction)

	var parent models.Parent
	db.First(&parent, requestData.ParentID)
	parent.TokensAmount += requestData.TokensAmount
	db.Save(&parent)

	c.JSON(http.StatusOK, gin.H{"message": "Transaction completed successfully"})
}
