package service

import (
	"github.com/google/uuid"
	. "polaris/internal/application/domain"
)

type FindAdRequest struct {
	Id string
}

type FindAdResponse struct {
	Id          string
	Title       string
	Description string
	CreatedAt   string
}

type FindAdService struct {
	adRepository Ads
}

func (service *FindAdService) Execute(request FindAdRequest) *FindAdResponse {
	id, errorParsing := uuid.Parse(request.Id)
	if errorParsing != nil {
		return nil
	}
	savedAd := service.adRepository.FindById(id)
	if savedAd == nil {
		return nil
	}
	return &FindAdResponse{
		Id:          savedAd.GetId().String(),
		Title:       savedAd.Title,
		Description: savedAd.Description,
		CreatedAt:   savedAd.GetCreatedAt().String(),
	}

}

func NewFindAdService(ads Ads) FindAdService {
	return FindAdService{ads}
}
