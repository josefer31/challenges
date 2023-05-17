package service

import (
	. "polaris/internal/application/domain"
)

const totalAds = 5

type FindAdsResponse struct {
	Ads []FindAdResponse
}

type FindAdsService struct {
	AdRepository Ads
}

func (service FindAdsService) Execute() FindAdsResponse {

	foundAds := service.AdRepository.FindAll()

	if len(foundAds) < totalAds {
		return createFindAdsResponseFrom(foundAds[:])
	}

	return createFindAdsResponseFrom(foundAds[:totalAds])
}

func createFindAdsResponseFrom(ads []Ad) FindAdsResponse {
	adsToResponse := make([]FindAdResponse, 0)
	for _, ad := range ads {
		adsToResponse = append(adsToResponse, FindAdResponse{
			Id:          ad.GetId().String(),
			Title:       ad.Title,
			Description: ad.Description,
			CreatedAt:   ad.GetCreatedAt().String(),
		})
	}
	return FindAdsResponse{Ads: adsToResponse}
}
