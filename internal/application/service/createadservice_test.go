package service

import (
	"github.com/stretchr/testify/assert"
	"polaris/internal/application/domain"
	"polaris/internal/application/mocks"
	"polaris/internal/test/fixtures"
	"testing"
)

func TestCreateAd(t *testing.T) {
	adRepository := new(mocks.Ads)
	clock := new(mocks.Clock)
	idGenerator := new(mocks.IdGenerator)
	service := NewCreateAdService(adRepository, idGenerator, clock)
	randomAd := fixtures.RandomAd()
	stubMocks(adRepository, randomAd, clock, idGenerator)
	request := adToCreateAdRequest(randomAd)
	expected := givenExpectedResponse(randomAd)

	actual, _ := service.Execute(request)

	assert.Equal(t, actual, expected)
	adRepository.AssertCalled(t, "Save", randomAd)
}

func TestReturnErrorWhenDescriptionGreaterThanFifty(t *testing.T) {
	adRepository := new(mocks.Ads)
	clock := new(mocks.Clock)
	idGenerator := new(mocks.IdGenerator)
	service := NewCreateAdService(adRepository, idGenerator, clock)
	randomAd := fixtures.RandomAdWithWrongDescriptionLen()
	stubMocks(adRepository, randomAd, clock, idGenerator)
	request := adToCreateAdRequest(randomAd)

	_, err := service.Execute(request)
	t.Run("Description greater than fifty must give an error", func(t *testing.T) {
		assert.Error(t, err)
		adRepository.AssertNotCalled(t, "Save", randomAd)

	})

}

func givenExpectedResponse(ad *domain.Ad) *CreateAdResponse {
	return &CreateAdResponse{
		Id:          ad.GetId().String(),
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
		CreatedAt:   ad.GetCreatedAt().String(),
	}
}

func stubMocks(adRepository *mocks.Ads, ad *domain.Ad, clock *mocks.Clock, idGenerator *mocks.IdGenerator) {
	adRepository.EXPECT().Save(ad).Return(ad, nil)
	clock.EXPECT().Now().Return(ad.GetCreatedAt())
	idGenerator.EXPECT().Next().Return(ad.GetId())
}

func adToCreateAdRequest(ad *domain.Ad) CreateAdRequest {
	return CreateAdRequest{
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
	}
}
