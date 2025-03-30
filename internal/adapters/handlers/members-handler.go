package handlers

import (
	"WISP/internal/core/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (h *Handler) CreateTeamMember(c *gin.Context) {
	teamID, err := uuid.Parse(c.Param("teamId"))
	if err != nil {
		log.Error().Err(err).Str("handler", "CreateTeamMember").Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		log.Error().Err(err).Str("handler", "CreateTeamMember").Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTeamMember, err := h.Services.CreateTeamMember(userID, teamID)
	if err != nil {
		log.Error().Err(err).Str("handler", "CreateTeamMember").Msg("Failed to create team member")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.ToTeamMemberDTO(
		createdTeamMember.UserID,
		createdTeamMember.TeamID,
		createdTeamMember.CreatedAt,
		createdTeamMember.UpdatedAt,
	)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *Handler) GetTeamMembers(c *gin.Context) {
	teamMembers, err := h.Services.GetTeamMembers()
	if err != nil {
		log.Error().Err(err).Str("handler", "GetTeamMembers").Msg("Failed to get team members")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res []dto.TeamMember
	for _, teamMember := range teamMembers {
		res = append(res, *dto.ToTeamMemberDTO(
			teamMember.UserID,
			teamMember.TeamID,
			teamMember.CreatedAt,
			teamMember.UpdatedAt,
		))
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) RemoveTeamMember(c *gin.Context) {
	teamID, err := uuid.Parse(c.Param("teamId"))
	if err != nil {
		log.Error().Err(err).Str("handler", "RemoveTeamMember").Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		log.Error().Err(err).Str("handler", "RemoveTeamMember").Msg("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Services.RemoveTeamMember(userID, teamID)
	if err != nil {
		log.Error().Err(err).Str("handler", "RemoveTeamMember").Msg("Failed to remove team member")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team member removed"})
}
