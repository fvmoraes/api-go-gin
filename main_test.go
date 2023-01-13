package main

import (
	"api-go-gin/controllers"
	"api-go-gin/initializers"
	"api-go-gin/tests"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEndpointCreateFoobar(t *testing.T) {
	initializers.StartDatabaseConnect()
	r := tests.RouteSetupTesting()
	r.POST("/foobar", controllers.CreateFoobar)
	contents := []byte(tests.MyFoobarContents)
	request, _ := http.NewRequest("POST", "/foobar", bytes.NewBuffer(contents))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	tests.SharedTestID = tests.GetIdFromHttpResponse(response)
	assert.Equal(t, http.StatusOK, response.Code, "Not expected status code")
}

func TestEndpointFoobarByParamId(t *testing.T) {
	initializers.StartDatabaseConnect()
	r := tests.RouteSetupTesting()
	r.GET("/foobar/:id", controllers.ShownFoobarByParamId)
	endpointPath := "/foobar/" + strconv.Itoa(tests.SharedTestID)
	request, _ := http.NewRequest("GET", endpointPath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Not expected status code")
}

//DELETE
//EDIT
//BY_PARAM_REG

func TestEndpointFoobarAll(t *testing.T) {
	initializers.StartDatabaseConnect()
	r := tests.RouteSetupTesting()
	r.GET("/foobar", controllers.ShownFoobar)
	request, _ := http.NewRequest("GET", "/foobar", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	// fmt.Println(response.Body)
	assert.Equal(t, http.StatusOK, response.Code, "Not expected status code")
}

func TestEndpointFoobarMock(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := tests.RouteSetupTesting()
	r.GET("/foobar/mock", controllers.ShownFoobarMockToLearnTests)

	request, _ := http.NewRequest("GET", "/foobar/mock", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Not expected status code")
}

func TestEndpointFoobarMockContent(t *testing.T) {
	r := tests.RouteSetupTesting()
	r.GET("/foobar/mock", controllers.ShownFoobarMockToLearnTests)
	request, _ := http.NewRequest("GET", "/foobar/mock", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	mock := tests.MyFoobarString
	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))
	// fmt.Println(mock)
	assert.Equal(t, string(mock), string(body), "Not expected contents")
}
