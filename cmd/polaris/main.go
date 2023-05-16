package main

import (
	"fmt"
	. "polaris/internal/application/domain"
	. "polaris/internal/application/service"
	. "polaris/internal/infrastructure/repository"
)

var (
	adRepository    = ProvideInMemoryRepository()
	createAdService = ProvideCreateAdService(adRepository)
	findAdService   = ProvideFindAdService(adRepository)
	findAdsService  = ProvideFindAdsService(adRepository)
)

func main() {

	fmt.Println("Go to save new ad")
	fmt.Println("--------------------------------")

	createdAd := createAnAd("titulo1", "descripcion 1", 12)
	fmt.Printf("Your new Ad %v was created at %v\n", createdAd.Id, createdAd.CreatedAt)
	fmt.Println("--------------------------------")

	foundAd, adNotFoundError := findAdService.Execute(FindAdRequest{Id: createdAd.Id})
	if adNotFoundError != nil {
		fmt.Printf("Error, Ad %v not found\n", createdAd.Id)
	} else {
		fmt.Printf("Found Ad  %v\n", foundAd)
	}
	fmt.Println("--------------------------------")

	createAnAd("titulo2", "descripcion 2", 12)
	createAnAd("titulo3", "descripcion 3", 12)
	createAnAd("titulo4", "descripcion 4", 12)
	createAnAd("titulo5", "descripcion 5", 12)
	createAnAd("titulo6", "descripcion 6", 12)
	createAnAd("titulo7", "descripcion 7", 12)
	foundAdResponse := findAdsService.Execute()

	fmt.Println("--------------------------------")
	fmt.Printf("Found Ads  %v\n", foundAdResponse.Ads)

}

func createAnAd(
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

func ProvideCreateAdService(adRepository AdRepository) CreateAdService {
	return CreateAdService{
		AdRepository: adRepository,
	}
}

func ProvideFindAdService(adRepository AdRepository) FindAdService {
	return FindAdService{
		AdRepository: adRepository,
	}
}

func ProvideFindAdsService(adRepository AdRepository) FindAdsService {
	return FindAdsService{
		AdRepository: adRepository,
	}
}

func ProvideInMemoryRepository() AdRepository {
	return &InMemoryAdRepository{}
}
