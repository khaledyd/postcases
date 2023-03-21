package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khaledyd/postcases/internal/database"
	"github.com/khaledyd/postcases/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

// sign up

func SignUp(c *gin.Context) {
	// Define a struct that represents the JSON payload
	type signUpInput struct {
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	// Bind the request body to the signUpInput struct
	var input signUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	// Create the user in the database
	user := model.User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: string(hashedPassword),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	// Create the response JSON object
	response := gin.H{
		"id":        user.ID,
		"full_name": user.FullName,
		"email":     user.Email,
	}
	c.JSON(http.StatusOK, response)
}

//sign in

func SignIn(c *gin.Context) {
	type SignInInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	var input SignInInput
	// validate the fields
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Query the user
	var user model.User
	if err := database.DB.Where(&model.User{Email: input.Email}).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}

	// Compare the hashed passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}

	// Prepare the response
	response := gin.H{
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	}
	c.JSON(http.StatusOK, response)
}
