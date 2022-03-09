package main_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mlenngren/iban-validator/ibanService/handlers"
	"github.com/stretchr/testify/assert"
)

func TestValidateIban(t *testing.T) {

	validateIbanRequest := handlers.ValidateIbanRequest{
		Iban:      "AL35202111 0900000000 01234567",
		Validated: 1,
	}

	reqBody, err := json.Marshal(validateIbanRequest)
	a := assert.New(t)

	if err != nil {
		a.Error(err)
	}

	req, w := setValidateRouter(bytes.NewBuffer(reqBody))

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := handlers.ValidateIbanRequest{}

	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := validateIbanRequest
	a.Equal(expected, actual)

}

func setValidateRouter(body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	r.POST("/validate", handlers.ValidateIban)
	req, err := http.NewRequest(http.MethodPost, "/validate", body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w
}
