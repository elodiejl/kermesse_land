package models

type Student struct {
	Base
	UserID      int  `json:"user_id" binding:"required"` // Référence vers la table User
	ParentID    int  `json:"parent_id"`                  // Référence vers la table Parent (parent)
	TokenAmount int  `json:"token_amount"`
	User        User `json:"user"`
}
