package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "polaris/internal/application/domain"
	"polaris/internal/application/mocks"
	"testing"
	"time"
)

func TestCreateAd(t *testing.T) {
	adRepository := new(mocks.AdRepository)
	clock := new(mocks.Clock)
	idGenerator := new(mocks.IdGenerator)
	service := CreateAdService{adRepository, idGenerator, clock}
	ad := givenAd()
	stubMocks(adRepository, ad, clock, idGenerator)
	request := givenAdToCreate()
	expected := givenExpectedResponse(ad)

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

func stubMocks(adRepository *mocks.AdRepository, ad Ad, clock *mocks.Clock, idGenerator *mocks.IdGenerator) {
	adRepository.EXPECT().Save(ad).Return(ad).Times(1)
	clock.EXPECT().Now().Return(ad.GetCreatedAt()).Times(1)
	idGenerator.EXPECT().Next().Return(ad.GetId()).Times(1)
}

func givenAdToCreate() CreateAdRequest {
	return CreateAdRequest{
		Title:       "Laptop",
		Description: "New apple laptop",
		Price:       12,
	}
}

func givenAd() Ad {
	id := uuid.New()
	now := time.Now()
	return NewAd(
		id,
		"Laptop",
		"New apple laptop",
		12,
		now,
	)
}
