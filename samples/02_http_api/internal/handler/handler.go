package handler

import (
	"net/http"

	"reservation_api/internal/domain"
	"reservation_api/internal/service"

	"github.com/gin-gonic/gin"
)

// ReservationHandler は予約に関するHTTPハンドラです
type ReservationHandler struct {
	service *service.ReservationService
}

// NewReservationHandler は新しいReservationHandlerを作成します
func NewReservationHandler(svc *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{service: svc}
}

// CreateReservation は予約を作成します
func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var req domain.CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation, err := h.service.CreateReservation(&req)
	if err != nil {
		if err.Error() == "時間が重複している予約が既に存在します" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reservation)
}

// ListReservations は予約一覧を取得します
func (h *ReservationHandler) ListReservations(c *gin.Context) {
	reservations, err := h.service.ListReservations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

// CancelReservation は予約をキャンセルします
func (h *ReservationHandler) CancelReservation(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.CancelReservation(id); err != nil {
		if err.Error() == "予約が見つかりません: "+id {
			c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

