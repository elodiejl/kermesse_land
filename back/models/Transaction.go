package models

// Transaction de jetons du parents
type Transaction struct {
	Base
	ParentID        int    `json:"parent_id"`     // Référence vers un parent
	Price           int    `json:"price"`         // Montant de la transaction (en monnaie réelle)
	TokensAmount    int    `json:"tokens_amount"` // Nombre de jetons achetés
	TransactionDate string `json:"transaction_date"`
}
