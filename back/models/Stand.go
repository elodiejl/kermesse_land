package models

type Stand struct {
	Base
	Name              string `json:"name"`
	StandType         string `json:"stand_type"`         // ENUM: 'nourriture', 'boisson', 'activité'
	ParticipationCost int    `json:"participation_cost"` // Coût en jetons pour participer
	TeneurID          int    `json:"teneur_id"`          // Référence vers la table User (teneur de stand)
	KermesseID        int    `json:"kermesse_id"`        // Référence vers la table Kermesse
	Stock             int    `json:"stock"`              // Stock pour les stands de nourriture ou boisson
}
