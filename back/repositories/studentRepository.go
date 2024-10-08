package repositories

import (
	"back/models"
	"gorm.io/gorm"
)

type StudentRepository interface {
	Create(student *models.Student) error
	FindByID(id uint) (*models.Student, error)
	FindByParentID(parentID uint) ([]models.Student, error)
	Update(student *models.Student) error
	Delete(id uint) error
}

type studentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{DB: db}
}

func (r *studentRepository) Create(student *models.Student) error {
	return r.DB.Create(student).Error
}

func (r *studentRepository) FindByID(id uint) (*models.Student, error) {
	var student models.Student
	if err := r.DB.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) FindByParentID(parentID uint) ([]models.Student, error) {
	var students []models.Student

	err := r.DB.Table("students").Select("students.*, users.username").Joins("JOIN users ON users.id = students.user_id").
		Where("students.parent_id = ?", parentID).Scan(&students).Error

	return students, err
}

func (r *studentRepository) Update(student *models.Student) error {
	return r.DB.Save(student).Error
}

func (r *studentRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Student{}, id).Error
}
