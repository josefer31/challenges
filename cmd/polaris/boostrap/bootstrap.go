package boostrap

import (
	"polaris/internal/application/domain"
	"polaris/internal/application/service"
	"polaris/internal/infrastructure/controller"
	"polaris/internal/infrastructure/repository"
)

var ads = repository.NewInMemoryAds()
var clock = domain.NewClock()
var idGenerator = domain.NewUUIDGenerator()
var createAdService = service.NewCreateAdService(ads, idGenerator, clock)
var findAdService = service.NewFindAdService(ads)
var findAdsService = service.NewFindAdsService(ads)
var adController = controller.NewAdController(createAdService, findAdService)

func ProvideAdController() controller.AdController { return adController }

func ProvideCreateAdService() service.CreateAdService { return createAdService }
func ProvideFindAdService() service.FindAdService     { return findAdService }

func ProvideFindAdsService() service.FindAdsService { return findAdsService }

func ProvideAds() domain.Ads { return ads }
