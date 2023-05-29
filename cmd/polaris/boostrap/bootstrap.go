package boostrap

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"polaris/internal/application/domain"
	"polaris/internal/application/service"
	"polaris/internal/infrastructure/controller"
	"polaris/internal/infrastructure/repository"
)

var ads = repository.NewInMemoryAds()
var dbClient = OpenDbClient()
var postgresAds = repository.NewPostgresAds(dbClient)
var clock = domain.NewClock()
var idGenerator = domain.NewUUIDGenerator()
var createAdService = service.NewCreateAdService(postgresAds, idGenerator, clock)
var findAdService = service.NewFindAdService(postgresAds)
var findAdsService = service.NewFindAdsService(postgresAds)
var adController = controller.NewAdController(createAdService, findAdService)

func ProvideAdController() controller.AdController { return adController }

func ProvideCreateAdService() service.CreateAdService { return createAdService }
func ProvideFindAdService() service.FindAdService     { return findAdService }

func ProvideFindAdsService() service.FindAdsService { return findAdsService }

func ProvideAds() domain.Ads { return postgresAds }

func OpenDbClient() *gorm.DB {
	dsn := "host=localhost user=polaris password=123123 dbname=polaris port=5432 sslmode=disable"
	dbClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		err := dbClient.AutoMigrate(&repository.Ad{})
		if err != nil {
			log.Fatalf("Error creating table - %v", err)
		}
		return dbClient
	}
}

func CloseDbClient() {
	err := dbClient.Migrator().DropTable("ads")
	if err != nil {
		log.Fatalf("Error dropping ads - %v", err)
	}
	if db, err := dbClient.DB(); err != nil {
		log.Fatalf("Error getting db - %v", err)
	} else {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing db - %v", err)
		}
	}

}
