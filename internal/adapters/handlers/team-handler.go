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

func (h *Handler) CreateTeam(c *gin.Context) {
	var team clients.TeamContract
	if err := c.ShouldBindJSON(&team); err != nil {
		log.Error().Err(err).Str("handler", "CreateTeam").Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTeam, err := h.Services.CreateTeam(&domain.Team{
		Name:        team.Name,
		Description: team.Description,
	})

	if err != nil {
		log.Error().Err(err).Str("handler", "CreateTeam").Msg("Failed to create team")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.ToTeamDTO(
		createdTeam.ID,
		createdTeam.Name,
		createdTeam.Description,
		createdTeam.CreatedAt,
		createdTeam.UpdatedAt,
	)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *Handler) GetTeams(c *gin.Context) {
	teams, err := h.Services.GetTeams()
	if err != nil {
		log.Error().Err(err).Str("handler", "GetTeams").Msg("Failed to get teams")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res []dto.Team
	for _, team := range teams {
		res = append(res, *dto.ToTeamDTO(
			team.ID,
			team.Name,
			team.Description,
			team.CreatedAt,
			team.UpdatedAt,
		))
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) GetTeamByID(c *gin.Context) {
	parsedID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Str("handler", "GetTeamByID").Msg("Failed to parse ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := h.Services.GetTeamByID(parsedID)
	if err != nil {
		log.Error().Err(err).Str("handler", "GetTeamByID").Msg("Failed to get team by ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.ToTeamDTO(
		team.ID,
		team.Name,
		team.Description,
		team.CreatedAt,
		team.UpdatedAt,
	)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) UpdateTeam(c *gin.Context) {
	parsedID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Str("handler", "UpdateTeam").Msg("Failed to parse ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var team clients.TeamContract
	if err := c.ShouldBindJSON(&team); err != nil {
		log.Error().Err(err).Str("handler", "UpdateTeam").Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTeam, err := h.Services.UpdateTeam(&domain.Team{
		ID:          parsedID,
		Name:        team.Name,
		Description: team.Description,
	})

	if err != nil {
		log.Error().Err(err).Str("handler", "UpdateTeam").Msg("Failed to update team")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.ToTeamDTO(
		updatedTeam.ID,
		updatedTeam.Name,
		updatedTeam.Description,
		updatedTeam.CreatedAt,
		updatedTeam.UpdatedAt,
	)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) DeleteTeam(c *gin.Context) {
	parsedID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error().Err(err).Str("handler", "DeleteTeam").Msg("Failed to parse ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Services.DeleteTeam(parsedID)
	if err != nil {
		log.Error().Err(err).Str("handler", "DeleteTeam").Msg("Failed to delete team")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
}
