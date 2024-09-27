package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id uint, user *models.User) error {
	err := r.db.First(&user, id).Error
	return err
}

// Créer un nouvel utilisateur
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Trouver un utilisateur par ID
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Trouver un utilisateur par email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Supprimer un utilisateur
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepository) IsEmailTaken(email string, excludeUserID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).
		Where("email = ? AND id <> ?", email, excludeUserID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepository) UpdateUser(id string, user *models.User) error {
	// Update the user record
	//if err := r.DB.Updates(user).Error; err != nil {
	//	return err
	//}
	if err := r.db.Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
