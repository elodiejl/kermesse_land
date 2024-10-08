package models

import "gorm.io/gorm"

type User struct {
	Base
	Username    string `gorm:"unique" json:"username" binding:"required" example:"jdoe"`
	LastName    string `json:"last_name" binding:"required" example:"Doe"`
	FirstName   string `json:"first_name" binding:"required" example:"John"`
	Email       string `gorm:"unique" json:"email" binding:"required" example:"john.doe@exmple.com"`
	Password    string `gorm:"not null" json:"password" binding:"required" example:"password"`
	Roles       uint8  `json:"roles" example:"0"` // 1 = organizer, 2 = admin, 4 = student , 8 = parents, 16 = stand_leader
	CreatedByID *uint  `json:"created_by_id"`
	CreatedBy   *User  `gorm:"foreignKey:CreatedByID"`
}

type UserRegister struct {
	Username  string `json:"username" binding:"required" example:"jdoe"`
	LastName  string `json:"last_name" binding:"required" example:"Doe"`
	FirstName string `json:"first_name" binding:"required" example:"John"`
	Email     string `json:"email" binding:"required" example:"john.doe@exmple.com"`
	Password  string `json:"password" binding:"required" example:"password"`
	Roles     string `json:"roles" binding:"required"`
}

type UserRegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required" example:"john.doe@exmple.com"`
	Password string `json:"password" binding:"required" example:"password"`
}

// PublicUser omits sensitive data from user model
type PublicUser struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
