package contract_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	mocks "polaris/internal/application/mocks/service"
	"polaris/internal/application/service"
	"polaris/internal/application/service/errorService"
	"polaris/internal/infrastructure/controller"
	"testing"
)

var bigDescription = fake.CharactersN(service.DescriptionMaxLength)
var title = fake.Title()
var price = uint(rand.Uint32())

func TestReturnInvalidResponseWhenMissSomeRequiredField(t *testing.T) {
	createAdServiceMock := mocks.NewCreateAdService(t)
	adController := controller.NewAdController(createAdServiceMock)
	newRecorder := httptest.NewRecorder()
	ginContext := getTestGinContext(newRecorder)
	mockNotValidJsonRequest(ginContext)

	adController.HandlerCreationAd(ginContext)

	assert.Equal(t, 400, newRecorder.Code)
	createAdServiceMock.AssertNotCalled(t, "Execute")
}

func TestReturnInvalidResponseWhenDescriptionIsGreaterThanFifty(t *testing.T) {
	createAdServiceMock := mocks.NewCreateAdService(t)
	adController := controller.NewAdController(createAdServiceMock)
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

func getTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}
