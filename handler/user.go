package handler

import (
	"fmt"
	"net/http"
	"ngevent-api/auth"
	"ngevent-api/helper"
	"ngevent-api/user"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     user.Service
	authService auth.Service
}

func NewUserHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{service, authService}
}

func (h *userHandler) RegisterNewUser(c *gin.Context) {
	var input user.CreateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		errorValidation := helper.FormatValidationError(err)
		errors := gin.H{"errors": errorValidation}

		response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/users/%d.png", time.Now().UnixNano())
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	registeredUser, err := h.service.CreateUser(input, path)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success create user", http.StatusOK, "success", registeredUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.service.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.Id)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	fetchUser, err := h.service.GetUserById(currentUser.Id)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(fetchUser, "")
	response := helper.APIResponse("fetch user data", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateDataUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	var input user.UpdateUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.User = currentUser
	updatedUser, err := h.service.UpdateUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(updatedUser, "")
	response := helper.APIResponse("success update user", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdatePasswordUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	var input user.UpdatePasswordUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.User = currentUser
	updatedUser, err := h.service.UpdatePassword(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(updatedUser, "")
	response := helper.APIResponse("success update password user", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateAvatarUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	file, err := c.FormFile("file")
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/users/%d.png", time.Now().UnixNano())
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateAvatar, err := h.service.UpdateAvatar(currentUser.Id, path)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = os.Remove(currentUser.Avatar)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(updateAvatar, "")
	response := helper.APIResponse("success update avatar user", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
