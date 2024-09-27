package models

type Tombola struct {
	Base
	KermesseID int    `json:"kermesse_id"` // Référence vers la Kermesse
	PrizeName  string `json:"prize_name"`  // Nom du lot
	DrawnAt    string `json:"drawn_at"`    // Date du tirage
	WinnerID   int    `json:"winner_id"`   // Référence vers le gagnant (élève ou parent) User
	TicketId   int    `json:"ticket_id"`   // Référence vers Ticket gagant
}
