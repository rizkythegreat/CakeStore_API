package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCakes(t *testing.T) {
	// Setup
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/cakes", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response []Cake
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	// Tambahkan assertions lainnya sesuai kebutuhan
}

func TestGetCake(t *testing.T) {
	// Setup
	router := setupRouter()

	// ID cake yang tersedia dalam database
	cakeID := 1

	req, _ := http.NewRequest("GET", "/cakes/"+strconv.Itoa(cakeID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response Cake
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, cakeID, response.ID)
	// Tambahkan assertions lainnya sesuai kebutuhan
}

func TestAddCake(t *testing.T) {
	// Setup
	router := setupRouter()

	// Data cake baru
	newCake := Cake{
		Title:       "New Cake",
		Description: "A delicious new cake",
		Rating:      9.5,
		Image:       "https://example.com/new_cake.jpg",
	}

	// Marshal cake menjadi JSON
	payload, _ := json.Marshal(newCake)

	req, _ := http.NewRequest("POST", "/cakes", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	// Tambahkan assertions lainnya sesuai kebutuhan
}

func TestUpdateCake(t *testing.T) {
	// Setup
	router := setupRouter()

	// ID cake yang akan diperbarui
	cakeID := 1

	// Data cake yang diperbarui
	updatedCake := Cake{
		Title:       "Updated Cake",
		Description: "An updated cake",
		Rating:      8.0,
		Image:       "https://example.com/updated_cake.jpg",
	}

	// Marshal cake menjadi JSON
	payload, _ := json.Marshal(updatedCake)

	req, _ := http.NewRequest("PUT", "/cakes/"+strconv.Itoa(cakeID), bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	// Tambahkan assertions lainnya sesuai kebutuhan
}

func TestDeleteCake(t *testing.T) {
	// Setup
	router := setupRouter()

	// ID cake yang akan dihapus
	cakeID := 1

	req, _ := http.NewRequest("DELETE", "/cakes/"+strconv.Itoa(cakeID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	// Tambahkan assertions lainnya sesuai kebutuhan
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/cakes", getCakes)
	router.GET("/cakes/:id", getCake)
	router.POST("/cakes", addCake)
	router.PUT("/cakes/:id", updateCake)
	router.DELETE("/cakes/:id", deleteCake)
	return router
}
