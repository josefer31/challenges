package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"polaris/internal/application/domain"
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
	findAdService   service.FindAdService
}

func (adController *AdController) HandlerCreationAd(context *gin.Context) {
	var bodyInput AdDtoInput

	if err := context.BindJSON(&bodyInput); err != nil {
		log.Printf("Error parsing input to JSON : %v", err.Error())
		return
	}

	createAdRequest := getCreateAdRequestFor(bodyInput)
	createdAdResponse, err := adController.createAdService.Execute(createAdRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		log.Printf("Error creating ad : %v", err.Error())
		return
	}

	context.JSON(http.StatusCreated, AdDtoResponse{
		Id:          createdAdResponse.Id,
		Title:       createdAdResponse.Title,
		Description: createdAdResponse.Description,
		Price:       createdAdResponse.Price,
		CreatedAt:   createdAdResponse.CreatedAt,
	})
}

func (adController *AdController) HandlerFindAd(context *gin.Context) {
	adId := context.Param("adId")
	if foundAd, err := adController.findAdService.Execute(service.FindAdRequest{Id: adId}); err != nil {
		handleError(context, err)
		return
	} else {
		context.JSON(http.StatusOK, AdDtoResponse{
			Id:          foundAd.Id,
			Title:       foundAd.Title,
			Description: foundAd.Description,
			Price:       foundAd.Price,
			CreatedAt:   foundAd.CreatedAt,
		})
	}
}

func handleError(context *gin.Context, err error) {
	if _, ok := err.(domain.AdNotFoundError); ok {
		context.JSON(http.StatusNotFound, http.NoBody)
	} else {
		context.JSON(http.StatusInternalServerError, http.NoBody)
	}
}

func getCreateAdRequestFor(bodyInput AdDtoInput) service.CreateAdRequest {
	return service.CreateAdRequest{
		Title:       bodyInput.Title,
		Description: bodyInput.Description,
		Price:       bodyInput.Price,
	}
}

func NewAdController(
	createAdService service.CreateAdService,
	findAdService service.FindAdService,
) AdController {
	return AdController{createAdService, findAdService}
}
