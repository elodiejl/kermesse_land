package models

type Token struct {
	Base
	StudentID int `json:"student_id" binding:"required"` // Référence vers la table Student
	ParentID  int `json:"parent_id" binding:"required"`  // Référence vers la table Parent
	Amount    int `json:"amount" binding:"required"`     // Nombre de jetons distribués
}
