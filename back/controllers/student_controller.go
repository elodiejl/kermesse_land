package controllers

import (
	"back/models"
	"back/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	repo repositories.StudentRepository
}

// NewStudentController crée une nouvelle instance de StudentController
func NewStudentController(repo repositories.StudentRepository) *StudentController {
	return &StudentController{repo: repo}
}

// CreateStudent gère la création d'un étudiant
// @Summary Créer un étudiant
// @Description Créer une nouvelle entrée étudiant
// @Tags students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student data"
// @Success 201 {object} models.Student "Student created"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /students [post]
func (ctrl *StudentController) CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := ctrl.repo.Create(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// GetStudentByID gère la récupération d'un étudiant par ID
// @Summary Trouver un étudiant par ID
// @Description Récupérer un étudiant par son ID
// @Tags students
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.Student "Student found"
// @Failure 400 {object} string "Invalid ID format"
// @Failure 404 {object} string "Student not found"
// @Failure 500 {object} string "Internal server error"
// @Router /students/{id} [get]
func (ctrl *StudentController) GetStudentByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	student, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// GetStudentsByParentID gère la récupération des étudiants d'un parent
// @Summary Trouver les étudiants d'un parent par ID
// @Description Récupérer tous les étudiants d'un parent donné
// @Tags students
// @Produce json
// @Param parent_id path int true "Parent ID"
// @Success 200 {array} models.Student "Students found"
// @Failure 400 {object} string "Invalid parent ID"
// @Failure 500 {object} string "Internal server error"
// @Router /students/parent/{parent_id} [get]
func (ctrl *StudentController) GetStudentsByParentID(c *gin.Context) {
	parentIDParam := c.Param("parent_id")
	parentID, err := strconv.Atoi(parentIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent ID"})
		return
	}

	students, err := ctrl.repo.FindByParentID(uint(parentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// UpdateStudent gère la mise à jour des informations d'un étudiant
// @Summary Mettre à jour un étudiant
// @Description Mettre à jour les informations d'un étudiant
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body models.Student true "Student data"
// @Success 200 {object} models.Student "Student updated"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "Student not found"
// @Failure 500 {object} string "Internal server error"
// @Router /students/{id} [put]
func (ctrl *StudentController) UpdateStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	existingStudent, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Mettre à jour les champs de l'étudiant
	existingStudent.UserID = student.UserID
	existingStudent.ParentID = student.ParentID
	existingStudent.TokenAmount = student.TokenAmount

	if err := ctrl.repo.Update(existingStudent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingStudent)
}

// DeleteStudent gère la suppression d'un étudiant
// @Summary Supprimer un étudiant
// @Description Supprimer un étudiant par son ID
// @Tags students
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {string} string "Student deleted"
// @Failure 400 {object} string "Invalid ID format"
// @Failure 404 {object} string "Student not found"
// @Failure 500 {object} string "Internal server error"
// @Router /students/{id} [delete]
func (ctrl *StudentController) DeleteStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	_, err = ctrl.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := ctrl.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
