package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"polaris/internal/application/domain"
	"polaris/internal/application/mocks"
	"polaris/internal/test/fixtures"
	"testing"
)

func TestFindExistingAd(t *testing.T) {
	ads := new(mocks.Ads)
	ad := fixtures.RandomAd()
	findAdService := NewFindAdService(ads)
	ads.EXPECT().FindById(mock.Anything).Return(&ad)

	expectedAdResponse := givenExpectedAdResponse(ad)
	actualResponse := findAdService.Execute(FindAdRequest{ad.GetId().String()})

	assert.Equal(t, *actualResponse, expectedAdResponse)
	ads.AssertCalled(t, "FindById", ad.GetId())
}

func TestFindNonExistingAdReturnNil(t *testing.T) {
	ads := new(mocks.Ads)

	findAdService := NewFindAdService(ads)
	ads.EXPECT().FindById(mock.Anything).Return(nil)

	randomId := uuid.New()
	actualResponse := findAdService.Execute(FindAdRequest{randomId.String()})

	assert.Nil(t, actualResponse)
	ads.AssertCalled(t, "FindById", randomId)
}

func TestFindUsingInvalidUuidReturnNil(t *testing.T) {
	ads := new(mocks.Ads)

	findAdService := NewFindAdService(ads)
	ads.EXPECT().FindById(mock.Anything).Return(nil)

	actualResponse := findAdService.Execute(FindAdRequest{"INVALID UUID"})

	assert.Nil(t, actualResponse)
	ads.AssertNotCalled(t, "FindById", "INVALID UUID")
}

func givenExpectedAdResponse(ad domain.Ad) FindAdResponse {
	return FindAdResponse{
		Id:          ad.GetId().String(),
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
		CreatedAt:   ad.GetCreatedAt().String(),
	}
}
