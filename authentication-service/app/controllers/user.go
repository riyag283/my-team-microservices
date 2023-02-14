package controllers

import (
	"authentication/app/models"
	"authentication/db"
	"authentication/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var reqBody models.User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	// validate request body
	if reqBody.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Username is required",
		})
		return
	}
	// check if username already exists
	var count int
	err := db.DBClient.QueryRow("SELECT COUNT(*) FROM users WHERE username=$1", reqBody.Username).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to check username",
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Username already exists",
		})
		return
	}
	if reqBody.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Password is required",
		})
		return
	}
	if reqBody.Role != models.RoleAdmin && reqBody.Role != models.RoleUser {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid role",
		})
		return
	}

	// encrypt password
	encryptedPwd, err := utils.EncryptPassword(reqBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to encrypt password",
		})
		return
	}

	var user models.UserCreated
	err = db.DBClient.QueryRow("INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id, username, role",
		reqBody.Username,
		encryptedPwd,
		reqBody.Role,
	).Scan(&user.ID, &user.Username, &user.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error": false,
		"user":  user,
	})
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var reqBody LoginRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	// retrieve user by username
	var user models.User
	err := db.DBClient.QueryRow("SELECT * FROM users WHERE username=$1", reqBody.Username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Invalid username",
		})
		return
	}
	fmt.Println(user.Password, reqBody.Password)

	// verify password
	if err := utils.VerifyPassword(user.Password, reqBody.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Incorrect password",
		})
		return
	}
	
	// generate JWT token
	token, err := utils.GenerateToken(user.ID, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"token": token,
	})
}
