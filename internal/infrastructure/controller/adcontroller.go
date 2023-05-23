package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"polaris/internal/application/service"
)

type AdDtoInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
}

type AdDtoResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	CreatedAt   string `json:"createdAt"`
}

type AdController struct {
	createAdService service.CreateAdService
}

func (adController *AdController) HandlerCreationAd(context *gin.Context) {
	var bodyInput AdDtoInput

	if errorBinding := context.BindJSON(&bodyInput); errorBinding != nil {
		return
	}

	createdAdResponse := adController.createAdService.Execute(service.CreateAdRequest{
		Title:       bodyInput.Title,
		Description: bodyInput.Description,
		Price:       bodyInput.Price,
	})

	context.JSON(http.StatusCreated, AdDtoResponse{
		Id:          createdAdResponse.Id,
		Title:       createdAdResponse.Title,
		Description: createdAdResponse.Description,
		Price:       createdAdResponse.Price,
		CreatedAt:   createdAdResponse.CreatedAt,
	})
}

func NewAdController(createAdService service.CreateAdService) AdController {
	return AdController{createAdService}
}
