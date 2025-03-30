package database

import (
	dbm "WISP/internal/adapters/repositories/database/models"
	"WISP/internal/core/domain"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (d *Database) CreateTeamMember(userID uuid.UUID, teamID uuid.UUID) (*domain.TeamMember, error) {
	teamMember := &dbm.TeamMember{
		UserID: userID,
		TeamID: teamID,
	}

	res := d.Create(teamMember)
	if res.Error != nil {
		log.Error().Err(res.Error).Str("repository", "CreateTeamMember").Msg("Failed to create team member")
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("TEAM MEMBER WASN'T CREATED")
	}

	return teamMember.To(), nil
}

func (d *Database) GetTeamMembers() ([]*domain.TeamMember, error) {
	var teamMembersDBM []*dbm.TeamMember

	err := d.Find(&teamMembersDBM).Error
	if err != nil {
		log.Error().Err(err).Str("repository", "GetTeamMembers").Msg("Failed to get team members")
		return nil, err
	}

	teamMembers := make([]*domain.TeamMember, 0)
	for _, teamMember := range teamMembersDBM {
		teamMembers = append(teamMembers, teamMember.To())
	}

	return teamMembers, nil
}

func (d *Database) RemoveTeamMember(userID uuid.UUID, teamID uuid.UUID) error {
	res := d.Unscoped().Where("user_id = ? AND team_id = ?", userID, teamID).Delete(&dbm.TeamMember{})
	if res.Error != nil {
		log.Error().Err(res.Error).Str("repository", "RemoveTeamMember").Msg("Failed to remove team member")
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("TEAM MEMBER WASN'T DELETED")
	}

	return nil
}
