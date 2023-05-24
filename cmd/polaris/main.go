package main

import (
	"fmt"
	"polaris/cmd/polaris/boostrap"
	. "polaris/internal/application/service"
	"polaris/pkg/server"
)

func main() {
	createAdService := boostrap.ProvideCreateAdService()
	findAdService := boostrap.ProvideFindAdService()
	findAdsService := boostrap.ProvideFindAdsService()

	fmt.Println("Go to save new ad")
	fmt.Println("--------------------------------")

	createdAd := createAd(createAdService, "titulo1", "descripcion 1", 12)
	fmt.Printf("Your new Ad %v was created at %v\n", createdAd.Id, createdAd.CreatedAt)
	fmt.Println("--------------------------------")

	foundAd, err := findAdService.Execute(FindAdRequest{Id: createdAd.Id})
	if err != nil {
		fmt.Printf("Error, Ad %v not found\n", createdAd.Id)
	} else {
		fmt.Printf("Found Ad  %v\n", foundAd)
	}
	fmt.Println("--------------------------------")

	createAd(createAdService, "titulo2", "descripcion 2", 12)
	createAd(createAdService, "titulo3", "descripcion 3", 12)
	createAd(createAdService, "titulo4", "descripcion 4", 12)
	createAd(createAdService, "titulo5", "descripcion 5", 12)
	createAd(createAdService, "titulo6", "descripcion 6", 12)
	createAd(createAdService, "titulo7", "descripcion 7", 12)
	foundAdResponse := findAdsService.Execute()

	fmt.Println("--------------------------------")
	fmt.Printf("Found Ads  %v\n", foundAdResponse.Ads)

	router := server.SetupRouter()
	router.Run(":8080")

}

func createAd(service CreateAdService, title string, description string, price uint) *CreateAdResponse {
	request := CreateAdRequest{
		Title:       title,
		Description: description,
		Price:       price,
	}

	createdAd, _ := service.Execute(request)
	return createdAd
}
