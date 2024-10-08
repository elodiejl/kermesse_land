package controllers

import (
	"back/config"
	"back/models"
	"back/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// Login godoc
// @Summary Login
// @Description Logs in a user
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.UserLogin true "User credentials"
// @Success 200 {string} string "Token"
// @Failure 400 {object} models.ErrorResponse "Invalid request"
// @Failure 401 {object} models.ErrorResponse "Invalid request"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /user/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var credentials models.UserLogin
	log.Println("Login")
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	user, err := ctrl.repo.GetByEmail(credentials.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	if !services.CheckPasswordHash(credentials.Password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := services.GenerateJWT(*user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

// Register godoc
// @Summary Register
// @Description Registers a new user
// @Tags auth
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Username"
// @Param last_name formData string true "Last Name"
// @Param first_name formData string true "First Name"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param roles formData string true "Role"
// @Success 201 {object} models.UserRegisterResponse "User registered"
// @Failure 400 {object} string "Invalid request"
// @Failure 500 {object} string "Internal server error"
// @Router /user/register [post]
func (ctrl *UserController) Register(c *gin.Context) {
	// Extract text fields
	username := c.PostForm("username")
	lastName := c.PostForm("last_name")
	firstName := c.PostForm("first_name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role := c.PostForm("roles")

	roleInt, err := strconv.Atoi(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role format"})
		return
	}

	// Validation du rôle
	validRoles := []uint8{
		config.RoleOrganizer,
		config.RoleStudent,
		config.RoleParent,
		config.RoleStandLeader,
	}

	isValidRole := false
	for _, validRole := range validRoles {
		if uint8(roleInt) == validRole {
			isValidRole = true
			break
		}
	}

	if !isValidRole {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	// Check if the email is already in use
	exists, err := ctrl.repo.ExistsByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check email existence", "details": err.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Check if the user already exists
	exists, err = ctrl.repo.ExistsByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check username existence", "details": err.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Hash password
	hashedPassword, err := services.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password", "details": err.Error()})
		return
	}

	// Create user model
	user := models.User{
		Username:  username,
		LastName:  lastName,
		FirstName: firstName,
		Email:     email,
		Password:  hashedPassword,
		Roles:     uint8(roleInt),
	}

	// Insert user into database
	if err := ctrl.repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user", "details": err.Error()})
		return
	}

	if uint8(roleInt) == config.RoleParent {
		parent := models.Parent{
			UserID:       int(user.ID),
			TokensAmount: 0,
		}

		if err := ctrl.parentRepo.Create(&parent); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create parent entry", "details": err.Error()})
			return
		}
	}

	// Prepare response
	response := models.UserRegisterResponse{
		Username: username,
		Email:    email,
	}
	c.JSON(http.StatusCreated, response)
}
