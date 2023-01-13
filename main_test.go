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
func TestEndpointFoobarAll(t *testing.T) {
	initializers.StartDatabaseConnect()
	r := tests.RouteSetupTesting()
	r.GET("/foobar", controllers.ShownFoobar)
	request, _ := http.NewRequest("GET", "/foobar", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
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
func TestEndpointFoobarByParamReg(t *testing.T) {
	initializers.StartDatabaseConnect()
	r := tests.RouteSetupTesting()
	r.GET("/foobar/param/:reg", controllers.ShownFoobarByParamReg)
	request, _ := http.NewRequest("GET", "/foobar/param/1010", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Not expected status code")
}
func TestEndpointEditFoobarByParamId(t *testing.T) {
	initializers.StartDatabaseConnect()
	r := tests.RouteSetupTesting()
	r.PATCH("/foobar/:id", controllers.EditFoobarByParamId)
	contents := []byte(tests.MyFoobarEdited)
	endpointPath := "/foobar/" + strconv.Itoa(tests.SharedTestID)
	request, _ := http.NewRequest("PATCH", endpointPath, bytes.NewBuffer(contents))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	tests.SharedTestID = tests.GetIdFromHttpResponse(response)
	assert.Equal(t, http.StatusOK, response.Code, "Not expected status code")
}
func TestEndpointDeleteFoobarByParamId(t *testing.T) {
	initializers.StartDatabaseConnect()
	r := tests.RouteSetupTesting()
	r.DELETE("/foobar/:id", controllers.DeleteFoobarByParamId)
	contents := []byte(tests.MyFoobarEdited)
	endpointPath := "/foobar/" + strconv.Itoa(tests.SharedTestID)
	request, _ := http.NewRequest("DELETE", endpointPath, bytes.NewBuffer(contents))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	tests.SharedTestID = tests.GetIdFromHttpResponse(response)
	assert.Equal(t, http.StatusAccepted, response.Code, "Not expected status code")
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
	assert.Equal(t, string(mock), string(body), "Not expected contents")
}
