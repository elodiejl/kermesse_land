package models

// Ticket Représente un billet de tombola acheté par un élève.
type Ticket struct {
	Base
	StudentID    int    `json:"student_id"`    // Référence vers un Student
	KermesseID   int    `json:"kermesse_id"`   // Référence vers la Kermesse
	TombolaID    int    `json:"tombola_id"`    // Référence vers la Tombola
	TicketNumber string `json:"ticket_number"` // Numéro unique du ticket
	PurchasedAt  string `json:"purchased_at"`  // Date d'achat des tickets
}
