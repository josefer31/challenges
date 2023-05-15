package service

import (
	"github.com/google/uuid"
	"polaris/internal/application/domain"
)

const totalAds = 5

type AdService struct {
	AdRepository domain.AdRepository
}

type CreateAdRequest struct {
	Title       string
	Description string
	Price       uint
}

type FindAdRequest struct {
	Id string
}

type AdResponse struct {
	Id          string
	Title       string
	Description string
	CreatedAt   string
}

func (service *AdService) Create(request CreateAdRequest) AdResponse {
	ad := domain.NewAd(request.Title, request.Description, request.Price)
	savedAd := service.AdRepository.Save(ad)
	return AdResponse{
		Id:          savedAd.GetId().String(),
		Title:       savedAd.Title,
		Description: savedAd.Description,
		CreatedAt:   savedAd.GetCreatedAt().String(),
	}
}

func (service *AdService) FindAd(request FindAdRequest) (*AdResponse, error) {
	id, _ := uuid.Parse(request.Id)
	savedAd, notFoundError := service.AdRepository.FindById(id)

	if notFoundError != nil {
		return nil, notFoundError
	}

	return &AdResponse{
		Id:          savedAd.GetId().String(),
		Title:       savedAd.Title,
		Description: savedAd.Description,
		CreatedAt:   savedAd.GetCreatedAt().String(),
	}, nil
}

func (service *AdService) FindAll() []AdResponse {
	foundAds := service.AdRepository.FindAll()

	if len(foundAds) < totalAds {
		return createFindAdsResponseFrom(foundAds[:])
	}

	return createFindAdsResponseFrom(foundAds[:totalAds])
}
func createFindAdsResponseFrom(ads []domain.Ad) []AdResponse {
	adsToResponse := make([]AdResponse, 0)
	for _, ad := range ads {
		adsToResponse = append(adsToResponse, AdResponse{
			Id:          ad.GetId().String(),
			Title:       ad.Title,
			Description: ad.Description,
			CreatedAt:   ad.GetCreatedAt().String(),
		})
	}
	return adsToResponse
}
