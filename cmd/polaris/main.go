package main

import (
	"fmt"
	. "polaris/internal/application/domain"
	. "polaris/internal/application/service"
	. "polaris/internal/infrastructure/repository"
)

func main() {
	adRepository := ProvideInMemoryRepository()
	createAdService := ProvideCreateAdService(adRepository)
	findAdService := ProvideFindAdService(adRepository)

	fmt.Println("Go to save new ad")

	request := CreateAdRequest{
		Title:       "Car Ad",
		Description: "New Ferrary car",
		Price:       200000,
	}

	createdAd := createAdService.Execute(request)
	fmt.Printf("Your Ad %v was created at %v\n", createdAd.Id, createdAd.CreatedAt)

	foundAd, adNotFoundError := findAdService.Execute(FindAdRequest{Id: createdAd.Id})
	if adNotFoundError != nil {
		fmt.Printf("Error, Ad %v not found\n", createdAd.Id)
	} else {
		fmt.Printf("Found Ad  %v\n", foundAd)
	}

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

func ProvideInMemoryRepository() AdRepository {
	return &InMemoryAdRepository{}
}
