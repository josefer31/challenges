package acceptanceTest

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"polaris/cmd/polaris/boostrap"
	"polaris/internal/application/domain"
	"polaris/pkg/server"
	"testing"
	"time"
)

const idInDB = "e85d27d4-3a6d-410f-a334-fdb52452fc17"
const titleInDb = "title"
const descriptionInDb = "description"
const priceInDb = 12

func TestMain(m *testing.M) {
	ads := boostrap.ProvideAds()
	id, _ := uuid.Parse(idInDB)
	_, err := ads.Save(domain.NewAd(id, titleInDb, descriptionInDb, priceInDb, time.Now()))
	if err != nil {
		log.Printf("Error trying to creating an ad - %v", err)
		return
	}
	exitVal := m.Run()
	defer func() { os.Exit(exitVal) }()
	defer boostrap.CloseDbClient()
}

func TestFindAd(t *testing.T) {
	jsonAsserter := jsonassert.New(t)
	recorder := httptest.NewRecorder()
	router := server.SetupRouter()
	url := fmt.Sprintf("/ads/%v", idInDB)
	request, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	jsonAsserter.Assertf(
		recorder.Body.String(),
		`{"title":"%s","description":"%s", "price":%d,"createdAt":"<<PRESENCE>>","id":"%s"}`,
		titleInDb,
		descriptionInDb,
		priceInDb,
		idInDB,
	)
}
