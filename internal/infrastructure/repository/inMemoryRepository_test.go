package repository

import (
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"polaris/internal/application/domain"
	"testing"
	"time"
)

var (
	adRepository = InMemoryAds{}
)

func tearDown() { adRepository.DeleteAll() }

func TestReturnEmptyAds(t *testing.T) {
	defer tearDown()
	actual := adRepository.FindAll()

	assert.Empty(t, actual)
}

func TestReturnAllAds(t *testing.T) {
	defer tearDown()
	expectedAds := createOneHundredRandomAds()
	fillRepository(adRepository, expectedAds)

	actualAds := adRepository.FindAll()

	assert.Equal(t, expectedAds, actualAds)
}

func TestReturnSomeAd(t *testing.T) {
	defer tearDown()
	listOfAds := createOneHundredRandomAds()
	fillRepository(adRepository, listOfAds)
	expectedAd := listOfAds[0]

	actualAd, _ := adRepository.FindById(expectedAd.GetId())

	assert.Equal(t, &expectedAd, actualAd)
}

func TestFindNonExistAdReturnError(t *testing.T) {
	defer tearDown()
	_, adNotFound := adRepository.FindById(uuid.New())

	assert.Error(t, adNotFound)
}

func TestSaveNewAd(t *testing.T) {
	defer tearDown()
	expectedAd := givenSomeAd()

	adRepository.Save(expectedAd)
	actualAd, _ := adRepository.FindById(expectedAd.GetId())

	assert.Equal(t, &expectedAd, actualAd)
}

func fillRepository(adRepository InMemoryAds, adsToSave []domain.Ad) {
	for _, ad := range adsToSave {
		adRepository.Save(ad)
	}
}

func createOneHundredRandomAds() []domain.Ad {
	createdAds := make([]domain.Ad, 100)
	for index := 0; index < 100; index++ {
		ad := givenSomeAd()
		createdAds[0] = ad
	}
	return createdAds
}

func givenSomeAd() domain.Ad {
	return domain.NewAd(
		uuid.New(),
		fake.Title(),
		fake.Characters(),
		uint(rand.Uint32()),
		time.Now(),
	)
}
