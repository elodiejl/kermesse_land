package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id uint, user *models.User) error
	UpdateUser(id string, user *models.User) error
	DeleteUser(id uint) error
	GetByEmail(email string) (*models.User, error)
	ExistsByEmail(email string) (bool, error)
	ExistsByUsername(username string) (bool, error)
	Create(user *models.User) error
	IsEmailTaken(email string, excludeUserID uint) (bool, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: DB}
}

func (r *UserRepositoryImpl) GetUserByID(id uint, user *models.User) error {
	err := r.db.First(&user, id).Error
	return err
}

// Créer un nouvel utilisateur
func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// Trouver un utilisateur par ID
func (r *UserRepositoryImpl) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Trouver un utilisateur par email
func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Supprimer un utilisateur
func (r *UserRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepositoryImpl) IsEmailTaken(email string, excludeUserID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).
		Where("email = ? AND id <> ?", email, excludeUserID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepositoryImpl) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepositoryImpl) UpdateUser(id string, user *models.User) error {
	// Update the user record
	//if err := r.DB.Updates(user).Error; err != nil {
	//	return err
	//}
	if err := r.db.Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepositoryImpl) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
