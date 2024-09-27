package seeders

import (
	"back/config"
	"back/models"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	users := []models.User{
		{
			Username: "organizer1",
			Email:    "organizer1@example.com",
			Password: "hashed_password_organizer1", // Remplacer par un vrai hash de mot de passe
			Roles:    config.RoleOrganizer,
		},
		{
			Username: "admin1",
			Email:    "admin1@example.com",
			Password: "hashed_password_admin1",
			Roles:    config.RoleAdmin,
		},
		{
			Username: "student1",
			Email:    "student1@example.com",
			Password: "hashed_password_student1",
			Roles:    config.RoleStudent,
		},
		{
			Username: "student2",
			Email:    "student2@example.com",
			Password: "hashed_password_student2",
			Roles:    config.RoleStudent,
		},
		{
			Username: "student3",
			Email:    "student3@example.com",
			Password: "hashed_password_student3",
			Roles:    config.RoleStudent,
		},
		{
			Username: "parent1",
			Email:    "parent1@example.com",
			Password: "hashed_password_parent1",
			Roles:    config.RoleParent,
		},
		{
			Username: "parent2",
			Email:    "parent2@example.com",
			Password: "hashed_password_parent2",
			Roles:    config.RoleParent,
		},
		{
			Username: "stand_leader1",
			Email:    "stand_leader1@example.com",
			Password: "hashed_password_stand_leader1",
			Roles:    config.RoleStandLeader,
		},
		{
			Username: "stand_leader2",
			Email:    "stand_leader2@example.com",
			Password: "hashed_password_stand_leader1",
			Roles:    config.RoleStandLeader,
		},
		{
			Username: "stand_leader3",
			Email:    "stand_leader3@example.com",
			Password: "hashed_password_stand_leader1",
			Roles:    config.RoleStandLeader,
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
