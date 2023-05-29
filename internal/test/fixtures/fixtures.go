package fixtures

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"math/rand"
	"polaris/internal/application/domain"
	"polaris/internal/infrastructure/repository"
)

func RandomAd() *domain.Ad {
	clock := domain.NewClock()
	return domain.NewAd(
		uuid.New(),
		fake.Title(),
		fake.Characters(),
		uint(rand.Uint32()),
		clock.Now(),
	)
}

func RandomAdWithWrongDescriptionLen() *domain.Ad {
	return domain.NewAd(
		uuid.New(),
		fake.Title(),
		fake.CharactersN(60),
		uint(rand.Uint32()),
		domain.NewClock().Now(),
	)
}

func AdsInDB() []repository.Ad {
	const fiveAds = 5
	adsInDb := make([]repository.Ad, fiveAds)
	ids := [fiveAds]string{
		"16eb0606-9f26-47db-8472-d49173532088",
		"6d8b4af2-fa0a-420c-a7a2-6c99c1c2fa21",
		"cc769ac0-377e-4c0b-81ed-04db0d270b5c",
		"4102b52f-3ddc-4b3b-8797-a7cf3d61917a",
		"5beb5187-a4a0-4422-ba4b-05a5823da236",
	}

	clock := domain.NewClock()
	for index, id := range ids {
		adsInDb[index] = repository.Ad{
			Id:          id,
			Title:       fmt.Sprintf("Title%v", index),
			Description: fmt.Sprintf("Description%v", index),
			CreatedAt:   clock.Now(),
		}
	}

	return adsInDb
}
