package fixtures

import (
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"math/rand"
	"polaris/internal/application/domain"
	"time"
)

func RandomAd() domain.Ad {
	return domain.NewAd(
		uuid.New(),
		fake.Title(),
		fake.Characters(),
		uint(rand.Uint32()),
		time.Now(),
	)
}
