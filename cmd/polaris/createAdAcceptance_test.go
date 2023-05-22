package main

import (
	"bytes"
	"encoding/json"
	"github.com/icrowley/fake"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAd(t *testing.T) {
	jsonAsserter := jsonassert.New(t)
	recorder := httptest.NewRecorder()
	givenTitle := fake.Title()
	givenDescription := fake.Characters()
	givenPrice := uint(rand.Uint32())
	bodyRequest := givenBodyRequest(givenTitle, givenDescription, givenPrice)

	router := SetupRouter()
	request, _ := http.NewRequest("POST", "/ads", bytes.NewBuffer(bodyRequest))
	router.ServeHTTP(recorder, request)

	assert.Equal(t, 201, recorder.Code)
	jsonAsserter.Assertf(
		recorder.Body.String(),
		`{"title":"%s","description":"%s", "price":%d,"createdAt":"<<PRESENCE>>","id":"<<PRESENCE>>"}`,
		givenTitle,
		givenDescription,
		givenPrice,
	)
}

func givenBodyRequest(title string, description string, price uint) []byte {
	bodyRequest := MockCreateAdRequest{
		Title:       title,
		Description: description,
		Price:       price,
	}
	body, _ := json.Marshal(bodyRequest)
	return body
}

type MockCreateAdRequest struct {
	Title       string
	Description string
	Price       uint
}
