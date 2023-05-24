package contract_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"polaris/internal/application/domain"
	mocks "polaris/internal/application/mocks/service"
	"polaris/internal/application/service"
	"polaris/internal/application/service/errorService"
	"polaris/internal/infrastructure/controller"
	"testing"
)

var bigDescription = fake.CharactersN(service.DescriptionMaxLength)
var title = fake.Title()
var price = uint(rand.Uint32())
var createAdServiceMock = new(mocks.CreateAdService)
var findAdServiceMock = new(mocks.FindAdService)

const adIdToFind = "e85d27d4-3a6d-410f-a334-fdb52452fc17"

func TestReturn400WhenMissSomeRequiredField(t *testing.T) {

	adController := controller.NewAdController(createAdServiceMock, findAdServiceMock)
	newRecorder := httptest.NewRecorder()
	ginContext := getTestGinContext(newRecorder)
	mockNotValidJsonRequest(ginContext)

	adController.HandlerCreationAd(ginContext)

	assert.Equal(t, 400, newRecorder.Code)
	createAdServiceMock.AssertNotCalled(t, "Execute")
}

func TestReturn400WhenDescriptionIsGreaterThanFifty(t *testing.T) {
	adController := controller.NewAdController(createAdServiceMock, findAdServiceMock)
	newRecorder := httptest.NewRecorder()
	ginContext := getTestGinContext(newRecorder)
	mockBigDescriptionRequest(ginContext)
	createAdServiceMock.EXPECT().Execute(mock.Anything).Return(
		nil,
		errorService.NewDescriptionLenError(bigDescription),
	)

	adController.HandlerCreationAd(ginContext)

	assert.Equal(t, 400, newRecorder.Code)
	createAdServiceMock.AssertCalled(t, "Execute", service.CreateAdRequest{
		Title:       title,
		Description: bigDescription,
		Price:       price,
	})
}

func TestReturn404WhenFindNotExistingAd(t *testing.T) {
	adController := controller.NewAdController(createAdServiceMock, findAdServiceMock)
	newRecorder := httptest.NewRecorder()
	ginContext := getTestGinContext(newRecorder)
	id, _ := uuid.Parse(adIdToFind)
	called := findAdServiceMock.EXPECT().Execute(mock.Anything).Return(nil, domain.NewAdNotFoundError(id))
	defer called.Unset()
	mockFindAdRequest(ginContext)

	adController.HandlerFindAd(ginContext)

	assert.Equal(t, 404, newRecorder.Code)
	findAdServiceMock.AssertCalled(t, "Execute", service.FindAdRequest{Id: adIdToFind})

}

func TestReturn500WhenFailsFindingAd(t *testing.T) {
	adController := controller.NewAdController(createAdServiceMock, findAdServiceMock)
	newRecorder := httptest.NewRecorder()
	ginContext := getTestGinContext(newRecorder)
	expect := findAdServiceMock.EXPECT().Execute(mock.Anything).Return(nil, errors.New("unexpected error"))
	defer expect.Unset()
	mockFindAdRequest(ginContext)

	adController.HandlerFindAd(ginContext)

	assert.Equal(t, 500, newRecorder.Code)
	findAdServiceMock.AssertCalled(t, "Execute", service.FindAdRequest{Id: adIdToFind})

}

func mockNotValidJsonRequest(c *gin.Context) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	inputData := map[string]string{
		"title":        fake.Title(),
		"description":  fake.Characters(),
		"unknownField": fake.Characters(),
	}
	jsonbytes, err := json.Marshal(inputData)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func mockBigDescriptionRequest(c *gin.Context) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(
		controller.AdDtoInput{
			Title:       title,
			Description: bigDescription,
			Price:       price,
		},
	)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func mockFindAdRequest(c *gin.Context) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	c.AddParam("adId", adIdToFind)

}

func getTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}
