package handlers

import (
	"WISP/internal/core/domain"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTime(c *gin.Context) {
	time, err := h.Services.GetCurrentTime()
	if err != nil {
		log.Error().Err(err).Msg("Ошибка получения текущего времени")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения текущего времени"})
		return
	}

	greeting, err := h.Services.GetGreeting()
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

func (h *Handler) GetTimeWithPathFormat(c *gin.Context) {
	format := c.Param("format")

	time, err := h.Services.GetTimeWithFormat(format)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка получения текущего времени")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения текущего времени"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"time":        time,
		"used_format": format,
		"params":      "path",
	})
}

func (h *Handler) GetTimeWithQueryFormat(c *gin.Context) {
	format := c.DefaultQuery("format", "2006-01-02 15:04:05")

	time, err := h.Services.GetTimeWithFormat(format)
	if err != nil {
		log.Error().Err(err).Str("format", format).Msg("Ошибка получения времени с форматом")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения времени с форматом"})
		return
	}

	withGreeting := c.Query("with_greeting") == "true"

	response := gin.H{
		"time":        time,
		"used_format": format,
		"param_type":  "query",
	}

	if withGreeting {
		greeting, err := h.Services.GetGreeting()
		if err == nil {
			response["greeting"] = greeting
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CalculateTimeDifference(c *gin.Context) {
	var request domain.TimeDifferenceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("Ошибка разбора JSON запроса")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Некорректный JSON запрос",
			"details": err.Error(),
		})
		return
	}

	difference, err := h.Services.CalculateTimeDifference(request.FromTime, request.ToTime)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при расчете разницы времени")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Ошибка при расчете разницы времени",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"from":       request.FromTime,
		"to":         request.ToTime,
		"difference": difference,
	})
}

func (h *Handler) GetTimeV2(c *gin.Context) {
	format := c.DefaultQuery("format", "2006-01-02 15:04:05")

	time, err := h.Services.GetTimeWithFormat(format)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка получения времени")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения времени"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"time":        time,
		"used_format": format,
		"api_version": "v2",
	})
}