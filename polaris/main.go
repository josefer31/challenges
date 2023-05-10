package main

import (
	"fmt"
	. "polaris/application/domain"
	"polaris/application/service"
	. "polaris/infrastructure/repository"
)

func main() {
	fmt.Println("Go to save new ad")

	createAdService := service.CreateAdService{
		AdRepository: ProvideInMemoryRepository(),
	}

	request := service.CreateAdRequest{
		Title:       "Car Ad",
		Description: "New Ferrary car",
		Price:       200000,
	}

	createdAd := createAdService.Execute(request)

	fmt.Printf("Your Ad %v was created at %v", createdAd.GetId(), createdAd.GetCreatedAt())
}

func ProvideInMemoryRepository() AdRepository {
	return &InMemoryAdRepository{}
}
