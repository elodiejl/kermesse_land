package models

type Token struct {
	Base
	ID        int `json:"id"`
	StudentID int `json:"student_id"` // Référence vers la table Student
	ParentID  int `json:"parent_id"`  // Référence vers la table Parent
	Amount    int `json:"amount"`     // Nombre de jetons distribués
}
