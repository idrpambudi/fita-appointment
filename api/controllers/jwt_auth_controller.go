package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/idrpambudi/fita-appointment/lib"
	"github.com/idrpambudi/fita-appointment/services"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      lib.Logger
	service     services.JWTAuthService
	userService services.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger lib.Logger,
	service services.JWTAuthService,
	userService services.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// SignIn signs in user
func (jwt JWTAuthController) SignIn(c *gin.Context) {
	jwt.logger.Info("SignIn route called")
	// Currently not checking for username and password
	// Can add the logic later if necessary.
	user, _ := jwt.userService.GetOneUser(uint(1))
	token := jwt.service.CreateToken(user)
	c.JSON(200, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}

// Register registers user
func (jwt JWTAuthController) Register(c *gin.Context) {
	jwt.logger.Info("Register route called")
	c.JSON(200, gin.H{
		"message": "register route",
	})
}
