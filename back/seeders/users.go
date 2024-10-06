package seeders

import (
	"back/config"
	"back/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Failed to hash the password: ", err)
		return nil
	}
	users := []models.User{
		{
			Username:  "organizer1",
			FirstName: "organizer",
			LastName:  "organizer",
			Email:     "organizer1@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleOrganizer,
		},
		{
			Username:  "admin1",
			FirstName: "admin",
			LastName:  "admin",
			Email:     "admin1@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleAdmin,
		},
		{
			Username:  "student1",
			FirstName: "student1",
			LastName:  "student1",
			Email:     "student1@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleStudent,
		},
		{
			Username:  "student2",
			FirstName: "student2",
			LastName:  "student2",
			Email:     "student2@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleStudent,
		},
		{
			Username:  "student3",
			FirstName: "student3",
			LastName:  "student3",
			Email:     "student3@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleStudent,
		},
		{
			Username:  "parent1",
			FirstName: "parent1",
			LastName:  "parent1",
			Email:     "parent1@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleParent,
		},
		{
			Username:  "parent2",
			FirstName: "parent2",
			LastName:  "parent2",
			Email:     "parent2@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleParent,
		},
		{
			Username:  "stand_leader1",
			FirstName: "standlead",
			LastName:  "standlead",
			Email:     "stand_leader1@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleStandLeader,
		},
		{
			Username:  "stand_leader2",
			FirstName: "standlead2",
			LastName:  "standlead2",
			Email:     "stand_leader2@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleStandLeader,
		},
		{
			Username:  "stand_leader3",
			FirstName: "standlead3",
			LastName:  "standlead3",
			Email:     "stand_leader3@example.com",
			Password:  string(hashedPassword),
			Roles:     config.RoleStandLeader,
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
