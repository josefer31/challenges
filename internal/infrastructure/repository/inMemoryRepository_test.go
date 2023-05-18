package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"polaris/internal/application/domain"
	"polaris/internal/test/fixtures"
	"testing"
)

var (
	adRepository = InMemoryAds{}
)

func tearDown() { adRepository.DeleteAll() }

func TestReturnEmptyAdsForEmptyRepository(t *testing.T) {
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

	actualAd := adRepository.FindById(expectedAd.GetId())

	assert.Equal(t, expectedAd, *actualAd)
}

func TestFindNonExistAdReturnNil(t *testing.T) {
	defer tearDown()
	adNotFound := adRepository.FindById(uuid.New())

	assert.Nil(t, adNotFound)
}

func TestSaveNewAd(t *testing.T) {
	defer tearDown()
	expectedAd := fixtures.RandomAd()

	adRepository.Save(expectedAd)
	actualAd := adRepository.FindById(expectedAd.GetId())

	assert.Equal(t, expectedAd, *actualAd)
}

func fillRepository(adRepository InMemoryAds, adsToSave []domain.Ad) {
	for _, ad := range adsToSave {
		adRepository.Save(ad)
	}
}

func createOneHundredRandomAds() []domain.Ad {
	createdAds := make([]domain.Ad, 100)
	for index := 0; index < 100; index++ {
		ad := fixtures.RandomAd()
		createdAds[0] = ad
	}
	return createdAds
}
