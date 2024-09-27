package models

type Prize struct {
	ID        int    `json:"id"`
	TombolaID int    `json:"tombola_id"` // Référence vers la Tombola
	Name      string `json:"name"`       // Nom du lot (ex: "Bicycle", "Télévision")
	//WinnerID  int    `json:"winner_id"`  // Référence vers l'élève ou parent gagnant, table User
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
