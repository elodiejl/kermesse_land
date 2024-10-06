package models

type ActivityParticipation struct {
	ID                uint   `json:"id"`
	UserID            int    `json:"user_id"`       // Référence vers un utilisateur (élève ou parent)
	ActivityID        int    `json:"activity_id"`   // Référence vers une activité (stand d'activité)
	PointsEarned      int    `json:"points_earned"` // Points gagnés lors de la participation à l'activité
	IsValidated       bool   `json:"is_validated"`
	ParticipationDate string `json:"participation_date"` // Date de la participation
}
