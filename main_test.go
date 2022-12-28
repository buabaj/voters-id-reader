package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/buabaj/voters-id-reader/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestRootHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome to the voters ID reader API","status":"success"}`
	r := setUpRouter()
	r.GET("/", controllers.Root)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}

// func TestReadIDHandler(t *testing.T) {
// 	r := setUpRouter()
// 	r.POST("/read", controllers.ReadID)
// 	w := httptest.NewRecorder()
// 	//make a request to the server with a file v_qr.jpg
// 	req, _ := http.NewRequest("POST", "/read", nil)
// 	req.Header.Set("Content-Type", "multipart/form-data")
// 	req.Header.Set("Content-Disposition", `form-data; name="file"; filename="v.jpg"`)
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }
