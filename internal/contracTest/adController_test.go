package contract_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	mocks "polaris/internal/application/mocks/service"
	"polaris/internal/infrastructure/controller"
	"testing"
)

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

func mockNotValidJsonRequest(c *gin.Context) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(`{"title":"any","description":"12","missingPriceField":"any"}`)
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
