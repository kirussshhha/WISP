package handlers

import (
	"WISP/internal/contracts/clients"
	"WISP/internal/core/domain"
	"WISP/internal/core/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user clients.UserContract
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error().Err(err).Str("handler", "CreateUser").Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.Services.CreateUser(&domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		log.Error().Err(err).Str("handler", "CreateUser").Msg("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.ToUserDTO(
		createdUser.ID,
		createdUser.Username,
		createdUser.Email,
		createdUser.CreatedAt,
		createdUser.UpdatedAt,
	)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *Handler) GetUsers(c *gin.Context) {
	email := c.Query("email")
	if email != "" {
		user, err := h.Services.GetUserByEmail(email)
		if err != nil {
			log.Error().Err(err).Str("handler", "GetUserByEmail").Msg("Failed to get user by email")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res := dto.ToUserDTO(
			user.ID,
			user.Username,
			user.Email,
			user.CreatedAt,
			user.UpdatedAt,
		)
		c.JSON(http.StatusOK, gin.H{"data": res})
		return
	}

	users, err := h.Services.GetUsers()
	if err != nil {
		log.Error().Err(err).Str("handler", "GetUsers").Msg("Failed to get users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res []dto.User
	for _, user := range users {
		res = append(res, *dto.ToUserDTO(
			user.ID,
			user.Username,
			user.Email,
			user.CreatedAt,
			user.UpdatedAt,
		))
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) GetUserByID(c *gin.Context) {
	parsedID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Str("handler", "GetUserByID").Msg("Failed to parse ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Services.GetUserByID(parsedID)
	if err != nil {
		log.Error().Err(err).Str("handler", "GetUserByID").Msg("Failed to get user by ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.ToUserDTO(
		user.ID,
		user.Username,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	parsedID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Str("handler", "UpdateUser").Msg("Failed to parse ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user clients.UserContract
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error().Err(err).Str("handler", "UpdateUser").Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.Services.UpdateUser(&domain.User{
		ID:       parsedID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		log.Error().Err(err).Str("handler", "UpdateUser").Msg("Failed to update user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.ToUserDTO(
		updatedUser.ID,
		updatedUser.Username,
		updatedUser.Email,
		updatedUser.CreatedAt,
		updatedUser.UpdatedAt,
	)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	parsedID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Str("handler", "DeleteUser").Msg("Failed to parse ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Services.DeleteUser(parsedID)
	if err != nil {
		log.Error().Err(err).Str("handler", "DeleteUser").Msg("Failed to delete user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
