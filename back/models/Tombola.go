package models

type Tombola struct {
	Base
	KermesseID int    `json:"kermesse_id" binding:"required"` // Référence vers la Kermesse
	PrizeName  string `json:"prize_name" binding:"required"`  // Nom du lot
	DrawnAt    string `json:"drawn_at" binding:"required"`    // Date du tirage
	WinnerID   int    `json:"winner_id"`                      // Référence vers le gagnant (élève ou parent) User
	TicketId   int    `json:"ticket_id"`                      // Référence vers Ticket gagnant
}
