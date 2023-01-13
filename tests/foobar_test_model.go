package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
)

var SharedTestID int

func GetIdFromHttpResponse(response *httptest.ResponseRecorder) int {
	type ResponseBodyID struct {
		ID int `json:"id"`
	}
	responseJSON, _ := ioutil.ReadAll(response.Body)
	var responseBodyId ResponseBodyID
	json.Unmarshal(responseJSON, &responseBodyId)
	return responseBodyId.ID
}
