package integrationTest

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"polaris/internal/application/domain"
	"polaris/internal/infrastructure/repository"
	"polaris/internal/test/fixtures"
	"testing"
)

var postgresAds domain.Ads
var connection *gorm.DB

func TestMain(m *testing.M) {
	connection = dbConnection()
	initDb(connection)
	postgresAds = repository.NewPostgresAds(connection)
	exitVal := m.Run()
	defer func() { os.Exit(exitVal) }()
	defer clearDB()
}

func clearDB() {
	err := connection.Migrator().DropTable("ads")
	if err != nil {
		log.Fatalf("Error cleaning table %v", err)
	}
	db, err := connection.DB()
	if err != nil {
		return
	}
	err = db.Close()
	if err != nil {
		return
	}
}

func initDb(connection *gorm.DB) {
	err := connection.AutoMigrate(&repository.Ad{})
	if err != nil {
		log.Fatalf("Error inserting an ad %v", err)
	}

	for _, ad := range fixtures.AdsInDB() {
		connection.Create(&ad)
	}

}

func TestPostgresqlRepository_Save(t *testing.T) {
	adToSave := fixtures.RandomAd()
	tests := []struct {
		name     string
		ads      domain.Ads
		givenAd  *domain.Ad
		expected *domain.Ad
	}{
		{
			name:     "Repository should save any Ad",
			ads:      postgresAds,
			givenAd:  adToSave,
			expected: adToSave,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			savedAd, _ := tt.ads.Save(tt.givenAd)

			actualAdInDb, _ := tt.ads.FindById(savedAd.GetId())

			assert.Equalf(t, tt.expected, actualAdInDb, "Save(%v)", tt.givenAd)
		})
	}
}

func dbConnection() *gorm.DB {
	dsn := "host=localhost user=polaris password=123123 dbname=polaris port=5432 sslmode=disable"
	dbClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		return dbClient
	}
	return nil
}
