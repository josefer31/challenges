package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	domain "polaris/internal/application/domain"
	"time"
)

type Ad struct {
	Id          string    `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	Price       uint
}
type PostgresqlRepository struct {
	dbClient *gorm.DB
}

func (p *PostgresqlRepository) Save(ad *domain.Ad) (*domain.Ad, error) {
	adToSave := Ad{
		Id:          ad.GetId().String(),
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
		CreatedAt:   ad.GetCreatedAt(),
	}

	if err := p.dbClient.Create(adToSave).Error; err != nil {
		return nil, err
	} else {
		return ad, nil
	}
}

func (p *PostgresqlRepository) FindById(id uuid.UUID) (*domain.Ad, error) {
	foundAd := Ad{Id: id.String()}

	if err := p.dbClient.First(&foundAd).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.NewAdNotFoundError(id)
		} else {
			return nil, err
		}
	} else {
		uuidId, errorParsing := uuid.Parse(foundAd.Id)
		if errorParsing != nil {
			return nil, errorParsing
		} else {

			return domain.NewAd(uuidId, foundAd.Title, foundAd.Description, foundAd.Price, foundAd.CreatedAt.UTC()), nil
		}

	}

}

func (p *PostgresqlRepository) FindAll() []domain.Ad {
	var adsInDb []Ad
	p.dbClient.Find(&adsInDb)
	domainAds := make([]domain.Ad, len(adsInDb))

	for index, adInDb := range adsInDb {
		id, _ := uuid.Parse(adInDb.Id)
		domainAds[index] = *domain.NewAd(
			id,
			adInDb.Title,
			adInDb.Description,
			adInDb.Price,
			adInDb.CreatedAt.UTC(),
		)
	}
	return domainAds
}

func NewPostgresAds(dbClient *gorm.DB) domain.Ads { return &PostgresqlRepository{dbClient: dbClient} }
