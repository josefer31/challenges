package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "polaris/internal/application/domain"
	. "polaris/internal/application/service"
	"polaris/internal/infrastructure/controller"
	. "polaris/internal/infrastructure/repository"
)

var (
	ads             = ProvideInMemoryAds()
	idGenerator     = ProvideIdGenerator()
	clock           = ProvideClock()
	createAdService = ProvideCreateAdService()
	findAdService   = ProvideFindAdService()
	findAdsService  = ProvideFindAdsService()
	adController    = ProvideAdController()
)

func main() {

	fmt.Println("Go to save new ad")
	fmt.Println("--------------------------------")

	createdAd := createAd("titulo1", "descripcion 1", 12)
	fmt.Printf("Your new Ad %v was created at %v\n", createdAd.Id, createdAd.CreatedAt)
	fmt.Println("--------------------------------")

	foundAd := findAdService.Execute(FindAdRequest{Id: createdAd.Id})
	if foundAd == nil {
		fmt.Printf("Error, Ad %v not found\n", createdAd.Id)
	} else {
		fmt.Printf("Found Ad  %v\n", foundAd)
	}
	fmt.Println("--------------------------------")

	createAd("titulo2", "descripcion 2", 12)
	createAd("titulo3", "descripcion 3", 12)
	createAd("titulo4", "descripcion 4", 12)
	createAd("titulo5", "descripcion 5", 12)
	createAd("titulo6", "descripcion 6", 12)
	createAd("titulo7", "descripcion 7", 12)
	foundAdResponse := findAdsService.Execute()

	fmt.Println("--------------------------------")
	fmt.Printf("Found Ads  %v\n", foundAdResponse.Ads)

	router := SetupRouter()
	router.Run(":8080")

}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/ads", func(ctx *gin.Context) {
		createdAd, err := adController.CreateAd(ctx)
		if err == nil {
			ctx.JSON(201, createdAd)
		}
	})
	return router
}

func createAd(
	title string,
	description string,
	price uint,
) CreateAdResponse {
	request := CreateAdRequest{
		Title:       title,
		Description: description,
		Price:       price,
	}

	createdAd := createAdService.Execute(request)
	return createdAd
}

func ProvideCreateAdService() CreateAdService {
	return NewCreateAdService(ads, idGenerator, clock)
}

func ProvideFindAdService() FindAdService {
	return NewFindAdService(ads)
}

func ProvideFindAdsService() FindAdsService {
	return NewFindAdsService(ads)
}

func ProvideInMemoryAds() Ads {
	return NewInMemoryAds()
}

func ProvideIdGenerator() IdGenerator {
	return NewUUIDGenerator()
}

func ProvideClock() Clock {
	return NewClock()
}

func ProvideAdController() controller.AdController {
	return controller.NewAdController(createAdService)
}
