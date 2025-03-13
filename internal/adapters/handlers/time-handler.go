package handlers

import (
	"WISP/internal/core/service"
	"net/http"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type TimeHandler struct {
	timeService service.TimeService
}

func NewTimeHandler(timeService service.TimeService) *TimeHandler {
	return &TimeHandler{
		timeService: timeService,
	}
}

func (h *TimeHandler) GetTime(c *gin.Context) {
	time, err := h.timeService.GetCurrentTime()
	if err != nil {
		log.Error().Err(err).Msg("Ошибка получения текущего времени")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения текущего времени"})
		return
	}

	greeting, err := h.timeService.GetGreeting()
	if err != nil {
		log.Error().Err(err).Msg("Ошибка получения приветствия")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения приветствия"})
		return
	}
	
	log.Info().Msg("Запрос на получение текущего времени")

	c.JSON(http.StatusOK, gin.H{
		"time":     time,
		"greeting": greeting,
	})
}

func (h *TimeHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/time", h.GetTime)
}
