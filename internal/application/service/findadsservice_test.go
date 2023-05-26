package service

import (
	"github.com/stretchr/testify/assert"
	"polaris/internal/application/domain"
	"polaris/internal/application/mocks"
	"polaris/internal/test/fixtures"
	"testing"
)

func TestFindAdsService_Execute(t *testing.T) {

	fiveAds := givenRandomAds(5)
	twelveAds := givenRandomAds(12)
	treeAds := givenRandomAds(3)

	expectedFiveAdsFromFiveAdsStored := expectedAdsResponseFor(fiveAds)
	expectedFiveAdsFromTwelveAdsStored := expectedAdsResponseFor(twelveAds)
	expectedTreeAdsFromTreeAdsStored := expectedAdsResponseFor(treeAds)
	expectedEmptyAdsForEmptyRepository := FindAdsResponse{make([]FindAdResponse, 0)}

	tests := []struct {
		name             string
		ads              mocks.Ads
		adsIndRepository []domain.Ad
		want             FindAdsResponse
	}{
		{
			name:             "Find all ads in repository when there is five ads",
			ads:              *mocks.NewAds(t),
			adsIndRepository: fiveAds,
			want:             expectedFiveAdsFromFiveAdsStored,
		},
		{
			name:             "Find five ads in repository when there is more than five ads",
			ads:              *mocks.NewAds(t),
			adsIndRepository: twelveAds,
			want:             expectedFiveAdsFromTwelveAdsStored,
		},
		{
			name:             "Find all ads in repository when there is less than five ads",
			ads:              *mocks.NewAds(t),
			adsIndRepository: treeAds,
			want:             expectedTreeAdsFromTreeAdsStored,
		},
		{
			name:             "Find empty list when there is not ads in the repository",
			ads:              *mocks.NewAds(t),
			adsIndRepository: make([]domain.Ad, 0),
			want:             expectedEmptyAdsForEmptyRepository,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := FindAdsService{
				&tt.ads,
			}
			tt.ads.EXPECT().FindAll().Return(tt.adsIndRepository)

			got := service.Execute()

			assert.Equalf(t, len(tt.want.Ads), len(got.Ads), "Total found ads")
			assert.Equalf(t, tt.want, got, "Comparition list")
			tt.ads.AssertCalled(t, "FindAll")
		})
	}
}

func expectedAdsResponseFor(ads []domain.Ad) FindAdsResponse {
	adsResponseSize := len(ads)
	if adsResponseSize > 5 {
		adsResponseSize = 5
	}

	foundAdsResponse := make([]FindAdResponse, adsResponseSize)
	for index, ad := range ads[:adsResponseSize] {
		foundAdsResponse[index] = FindAdResponse{
			Id:          ad.GetId().String(),
			Title:       ad.Title,
			Description: ad.Description,
			CreatedAt:   ad.GetCreatedAt().String(),
		}
	}
	return FindAdsResponse{foundAdsResponse}
}

func givenRandomAds(totalAds int) []domain.Ad {
	listOfAds := make([]domain.Ad, totalAds)
	for index := 0; index < totalAds; index++ {
		listOfAds[index] = *fixtures.RandomAd()
	}
	return listOfAds
}
