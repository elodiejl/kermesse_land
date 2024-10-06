package models

type Stand struct {
	Base
	Name              string `json:"name" binding:"required"`
	StandType         string `json:"stand_type" binding:"required"`         // ENUM: 'nourriture', 'boisson', 'activité'
	ParticipationCost int    `json:"participation_cost" binding:"required"` // Coût en jetons pour participer
	TeneurID          int    `json:"teneur_id"`                             // Référence vers la table User (teneur de stand)
	KermesseID        int    `json:"kermesse_id" binding:"required"`        // Référence vers la table Kermesse
	Stock             int    `json:"stock" binding:"required"`              // Stock pour les stands de nourriture ou boisson
}
