package models

type Activity struct {
	Base
	Name          string `json:"name"`
	StandID       int    `json:"stand_id"`       // Référence vers un Stand
	PointsAwarded int    `json:"points_awarded"` // Points attribués pour l'activité
	//InteractionDate string `json:"interaction_date"`
}
