package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"reservation_api/internal/domain"
	"reservation_api/internal/repository"
	"reservation_api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	repo := repository.NewMemoryRepository()
	svc := service.NewReservationService(repo)
	h := NewReservationHandler(svc)

	r := gin.New()
	api := r.Group("/api/v1")
	{
		api.POST("/reservations", h.CreateReservation)
		api.GET("/reservations", h.ListReservations)
		api.DELETE("/reservations/:id", h.CancelReservation)
	}

	return r
}

func TestCreateReservation(t *testing.T) {
	router := setupRouter()

	reqBody := domain.CreateReservationRequest{
		UserName:     "Alice",
		ResourceName: "Room-A",
		StartsAt:     "2025-01-02T09:00:00Z",
		EndsAt:       "2025-01-02T10:00:00Z",
	}

	jsonData, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/api/v1/reservations", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestListReservations(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/api/v1/reservations", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

