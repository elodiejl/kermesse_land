package models

type Parent struct {
	Base
	UserID       int `json:"user_id"`                 // Référence vers la table User
	TokensAmount int `json:"tokens_amount_available"` // Solde de jetons disponibles
}
