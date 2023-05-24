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
	Price       uint
	CreatedAt   string
}
type FindAdService interface {
	Execute(request FindAdRequest) (*FindAdResponse, error)
}
type FindAdServiceImpl struct {
	adRepository Ads
}

func (service *FindAdServiceImpl) Execute(request FindAdRequest) (*FindAdResponse, error) {
	id, errorParsing := uuid.Parse(request.Id)
	if errorParsing != nil {
		return nil, errorParsing
	}

	if foundAd, err := service.adRepository.FindById(id); err != nil {
		return nil, err
	} else {
		return &FindAdResponse{
			Id:          foundAd.GetId().String(),
			Title:       foundAd.Title,
			Description: foundAd.Description,
			Price:       foundAd.Price,
			CreatedAt:   foundAd.GetCreatedAt().String(),
		}, nil
	}

}

func NewFindAdService(ads Ads) FindAdService {
	return &FindAdServiceImpl{ads}
}
