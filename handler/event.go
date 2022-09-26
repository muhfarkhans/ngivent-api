package handler

import (
	"errors"
	"fmt"
	"net/http"
	"ngevent-api/auth"
	"ngevent-api/event"
	"ngevent-api/helper"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type eventHandler struct {
	service     event.Service
	authService auth.Service
}

func NewEventHandler(service event.Service, authService auth.Service) *eventHandler {
	return &eventHandler{service, authService}
}

func (h *eventHandler) CreateNewEvent(c *gin.Context) {
	var input event.CreateEventInput
	err := c.ShouldBind(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error create event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/events/%d.png", time.Now().UnixNano())
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newEvent, err := h.service.CreateEvent(input, path)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error create event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success create event", http.StatusOK, "success", newEvent)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) GetAllEvent(c *gin.Context) {
	events, err := h.service.GetEvents()
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := event.FormatEvents(events)
	response := helper.APIResponse("success get event", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) GetEventById(c *gin.Context) {
	var input event.GetEventDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error get event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	detailEvent, err := h.service.FindEvent(input)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if detailEvent.Id == 0 {
		errors := gin.H{"errors": errors.New("event not found")}

		response := helper.APIResponse("event not found", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := event.FormatEventDetail(detailEvent)
	response := helper.APIResponse("success get event", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) CreateNewImageEvent(c *gin.Context) {
	var input event.GetEventDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error create event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	path := fmt.Sprintf("images/events/%d.png", time.Now().UnixNano())
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file upload error", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newImageEvent, err := h.service.CreateEventImage(input, path)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error create image event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success create image event", http.StatusOK, "success", newImageEvent)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) UpdateDataEvent(c *gin.Context) {
	var input event.CreateEventInput
	err := c.ShouldBind(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error update event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputUri event.GetEventDetailInput
	err = c.ShouldBindUri(&inputUri)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error update event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedEvent, err := h.service.UpdateEvent(input, inputUri)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error update event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := event.FormatEventDetail(updatedEvent)
	response := helper.APIResponse("success update event", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) MakeImagePrimary(c *gin.Context) {
	var input event.UpdatePrimaryImageInput
	err := c.ShouldBind(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error create event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputUri event.GetEventImageDetailInput
	err = c.ShouldBindUri(&inputUri)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error update image event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedEvent, err := h.service.UpdateImagePrimary(input, inputUri)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error update image event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := event.FormatEventDetail(updatedEvent)
	response := helper.APIResponse("success update image event", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) DeleteEvent(c *gin.Context) {
	var inputUri event.GetEventDetailInput
	err := c.ShouldBindUri(&inputUri)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error delete event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deletedEvent, err := h.service.DeleteEvent(inputUri)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error update image event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	for _, image := range deletedEvent.EventImages {
		err = os.Remove(image.Image)
	}

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file delete error", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := gin.H{"message": "deleted"}
	response := helper.APIResponse("success delete event", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *eventHandler) DeleteEventImage(c *gin.Context) {
	var inputUri event.GetEventImageDetailInput
	err := c.ShouldBindUri(&inputUri)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error delete event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deletedEvent, err := h.service.DeleteEventImage(inputUri)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error update image event", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = os.Remove(deletedEvent.Image)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("file delete error", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := gin.H{"message": "deleted"}
	response := helper.APIResponse("success delete event image", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
