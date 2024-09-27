package models

type Message struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"sender_id"`   // Référence vers un User envoyant le message avec rôle teneur ou organizer
	ReceiverID int    `json:"receiver_id"` // Référence vers un User recevant le message avec rôle teneur ou organizer
	Message    string `json:"message"`     // Contenu du message
	SentAt     string `json:"sent_at"`     // Date d'envoi
}
