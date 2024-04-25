package controllers

import (
	"net/http"

	"github.com/Pure227/Grittaya_backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB: DB}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	ac := NewAuthController(uc.DB)
	// 1. Retrieve user data from token (assuming GetUserDataByToken is secure)
	userData, err := ac.GetUserDataByToken(ctx)
	if err != nil {
		// Handle error appropriately (e.g., return unauthorized or bad request)
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Invalid token"})
		return
	}

	// 2. Use a dedicated user model (improves data management and security)
	user := models.Admin{
		ID:       userData.ID,
		Username: userData.Username,
		Nickname: userData.Nickname,
		Position: userData.Position,
		// Add other relevant user data fields (if applicable)
	}

	// 3. Respond with user data (excluding potentially sensitive fields)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": user}})
}
