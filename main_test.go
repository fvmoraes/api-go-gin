package main

import (
	"api-go-gin/controllers"
	"api-go-gin/helpers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func RouteSetupTesting() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestEndpointFoobarMock(t *testing.T) {
	r := RouteSetupTesting()
	r.GET("/foobar/mock", controllers.ShownFoobarMockToLearnTests)

	request, _ := http.NewRequest("GET", "/foobar/mock", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Not expected status code")
}

func TestEndpointFoobarMockContent(t *testing.T) {
	r := RouteSetupTesting()
	r.GET("/foobar/mock", controllers.ShownFoobarMockToLearnTests)
	request, _ := http.NewRequest("GET", "/foobar/mock", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	mock := helpers.MyFoobarString
	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))
	// fmt.Println(mock)
	assert.Equal(t, string(mock), string(body), "Not expected content")
}
