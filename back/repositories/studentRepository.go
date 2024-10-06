package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type StudentRepositoryInterface interface {
	Create(student *models.Student) error
	FindByID(id uint) (*models.Student, error)
	Update(student *models.Student) error
	Delete(id uint) error
	FindByParentID(parentID int) ([]models.Student, error) // Méthode pour trouver les étudiants par ID de parent
}

type StudentRepository struct {
	db *gorm.DB
}

// NewStudentRepository Retourne une nouvelle instance de StudentRepository
func NewStudentRepository(db *gorm.DB) StudentRepositoryInterface {
	return &StudentRepository{db: db}
}

// Create Créer un nouvel étudiant
func (r *StudentRepository) Create(student *models.Student) error {
	return r.db.Create(student).Error
}

// FindByID Trouver un étudiant par ID
func (r *StudentRepository) FindByID(id uint) (*models.Student, error) {
	var student models.Student
	if err := r.db.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

// Update Met à jour les informations d'un étudiant
func (r *StudentRepository) Update(student *models.Student) error {
	return r.db.Save(student).Error
}

// Delete Supprimer un étudiant par ID
func (r *StudentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}

// FindByParentID Trouver tous les étudiants d'un parent par ID
func (r *StudentRepository) FindByParentID(parentID int) ([]models.Student, error) {
	var students []models.Student
	if err := r.db.Where("parent_id = ?", parentID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
