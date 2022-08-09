package users

import (
	"gomp/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
	authService auth.Service
}

func NewUserHandler(userService UserService, authService auth.Service) *UserHandler {
	return &UserHandler{userService, authService}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var input RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	user, err := uh.userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	token, err := uh.authService.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	userDTO := FormatUserDTO(user, token)

	c.JSON(http.StatusCreated, userDTO)
}

func (uh *UserHandler) LoginUser(c *gin.Context) {
	var input LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	loginUser, err := uh.userService.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}

	token, err := uh.authService.GenerateToken(loginUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	userDTO := FormatUserDTO(loginUser, token)

	c.JSON(http.StatusOK, userDTO)
}
