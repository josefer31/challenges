package main

import (
	"fmt"
	. "polaris/internal/application/domain"
	. "polaris/internal/application/service"
	. "polaris/internal/infrastructure/repository"
)

var (
	adRepository = ProvideInMemoryRepository()
	adService    = ProvideAdService(adRepository)
)

func main() {

	fmt.Println("Go to save new ad")
	fmt.Println("--------------------------------")

	createdAd := createAd("titulo1", "descripcion 1", 12)
	fmt.Printf("Your new Ad %v was created at %v\n", createdAd.Id, createdAd.CreatedAt)
	fmt.Println("--------------------------------")

	foundAd, adNotFoundError := adService.FindAd(FindAdRequest{Id: createdAd.Id})
	if adNotFoundError != nil {
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
	foundAdResponse := adService.FindAll()

	fmt.Println("--------------------------------")
	fmt.Printf("Found Ads  %v\n", foundAdResponse)

}

func createAd(
	title string,
	description string,
	price uint,
) AdResponse {
	request := CreateAdRequest{
		Title:       title,
		Description: description,
		Price:       price,
	}

	createdAd := adService.Create(request)
	return createdAd
}

func ProvideAdService(adRepository AdRepository) AdService {
	return AdService{
		AdRepository: adRepository,
	}
}

func ProvideInMemoryRepository() AdRepository {
	return &InMemoryAdRepository{}
}
