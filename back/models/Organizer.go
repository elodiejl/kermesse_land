package models

type Organizer struct {
	Base
	UserID     int `json:"user_id"`     // Référence vers un User ayant le rôle 'organisateur'
	KermesseID int `json:"kermesse_id"` // Référence vers une Kermesse
}
