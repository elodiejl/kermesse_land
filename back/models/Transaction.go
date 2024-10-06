package models

// Transaction de jetons du parents
type Transaction struct {
	Base
	ParentID        int    `json:"parent_id" binding:"required"`     // Référence vers un parent
	Price           int    `json:"price" binding:"required"`         // Montant de la transaction (en monnaie réelle)
	TokensAmount    int    `json:"tokens_amount" binding:"required"` // Nombre de jetons achetés
	TransactionDate string `json:"transaction_date"`
}
