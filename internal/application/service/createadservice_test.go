package service

import (
	"github.com/stretchr/testify/assert"
	. "polaris/internal/application/domain"
	"polaris/internal/application/mocks"
	"polaris/internal/test/fixtures"
	"testing"
)

func TestCreateAd(t *testing.T) {
	adRepository := new(mocks.Ads)
	clock := new(mocks.Clock)
	idGenerator := new(mocks.IdGenerator)
	service := CreateAdService{adRepository, idGenerator, clock}
	randomAd := fixtures.RandomAd()
	stubMocks(adRepository, randomAd, clock, idGenerator)
	request := adToCreateAdRequest(randomAd)
	expected := givenExpectedResponse(randomAd)

	actual := service.Execute(request)

	assert.Equal(t, actual, expected)
}

func givenExpectedResponse(ad Ad) CreateAdResponse {
	return CreateAdResponse{
		Id:          ad.GetId().String(),
		Title:       ad.Title,
		Description: ad.Description,
		CreatedAt:   ad.GetCreatedAt().String(),
	}
}

func stubMocks(adRepository *mocks.Ads, ad Ad, clock *mocks.Clock, idGenerator *mocks.IdGenerator) {
	adRepository.EXPECT().Save(ad).Return(ad).Times(1)
	clock.EXPECT().Now().Return(ad.GetCreatedAt()).Times(1)
	idGenerator.EXPECT().Next().Return(ad.GetId()).Times(1)
}

func adToCreateAdRequest(ad Ad) CreateAdRequest {
	return CreateAdRequest{
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
	}
}
