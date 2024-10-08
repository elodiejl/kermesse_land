package controllers

import (
	"back/models"
	"back/repositories"
	"back/services"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	repo       repositories.UserRepository
	parentRepo repositories.ParentRepository
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{repo: repo}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object to create"
// @Success 201 {object} models.User "User created"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /user [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetMe godoc
// @Summary Get the current user
// @Description Retrieve the currently authenticated user's details using their token
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.User "The currently authenticated user"
// @Failure 401 {object} string "Unauthorized: Invalid token"
// @Failure 404 {object} string "User not found"
// @Router /user/me [get]
func (ctrl *UserController) GetMe(c *gin.Context) {
	userID, err := ctrl.getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	var user models.User
	if err := ctrl.repo.GetUserByID(userID, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateMe godoc
// @Summary Update current user's profile
// @Description Update the profile information of the currently authenticated user, including password.
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param first_name formData string false "First name of the user"
// @Param last_name formData string false "Last name of the user"
// @Param email formData string false "Email address of the user"
// @Param profile_picture formData file false "Profile picture file"
// @Param old_password formData string true "Current password for verification"
// @Param new_password formData string false "New password for the user"
// @Param skills formData string false "Array of skill IDs"
// @Success 200 {object} models.PublicUser "Successfully updated user profile"
// @Failure 400 {object} string "Bad request if the provided data is incorrect"
// @Failure 401 {object} string "Unauthorized if the user's old password is incorrect or token is invalid"
// @Failure 404 {object} string "Not Found if the user does not exist"
// @Failure 500 {object} string "Internal Server Error for any server errors"
// @Router /user/me [put]
func (ctrl *UserController) UpdateMe(c *gin.Context) {
	userID, err := ctrl.getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
		return
	}

	var user models.User
	if err := ctrl.repo.GetUserByID(userID, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Retrieve form data
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	email := c.PostForm("email")
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	// Check if the old password is correct before updating
	if oldPassword != "" && newPassword != "" {
		if !services.CheckPasswordHash(oldPassword, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Old password is incorrect"})
			return
		}
		hashedPassword, err := services.HashPassword(newPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash new password", "details": err.Error()})
			return
		}
		user.Password = hashedPassword
	}

	// Update fields
	if firstName != "" {
		user.FirstName = firstName
	}
	if lastName != "" {
		user.LastName = lastName
	}
	if email != "" && email != user.Email {
		isTaken, err := ctrl.repo.IsEmailTaken(email, user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email availability", "details": err.Error()})
			return
		}
		if isTaken {
			c.JSON(http.StatusBadRequest, gin.H{"error": "This email is already in use by another user"})
			return
		}
		user.Email = email
	}

	if err := ctrl.repo.UpdateUser(strconv.Itoa(int(userID)), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete the currently authenticated user
// @Description Delete the user account of the currently authenticated user.
// @Tags users
// @Security ApiKeyAuth
// @Success 200 {string} string "User deleted successfully"
// @Failure 401 {object} string "Unauthorized: Invalid token"
// @Failure 404 {object} string "User not found"
// @Router /user [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := ctrl.getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.repo.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// Helper method to get user ID from token
func (ctrl *UserController) getUserIDFromToken(c *gin.Context) (uint, error) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return 0, errors.New("no authorization token provided")
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	userID, err := services.GetUserIDFromToken(token)
	if err != nil {
		return 0, errors.New("unauthorized: invalid token")
	}

	return userID, nil
}
