package database

import (
	dbm "WISP/internal/adapters/repositories/database/models"
	"WISP/internal/core/domain"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (d *Database) CreateTeam(team *domain.Team) (*domain.Team, error) {
	teamDB := dbm.NewTeamDBM(team)
	res := d.Create(teamDB)
	if res.Error != nil {
		log.Error().Err(res.Error).Str("repository", "CreateTeam").Msg("Failed to create team")
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("TEAM WASN'T CREATED")
	}

	team = teamDB.To()
	return team, nil
}

func (d *Database) GetTeams() ([]*domain.Team, error) {
	var teamsDBM []*dbm.Team

	err := d.Find(&teamsDBM).Error
	if err != nil {
		log.Error().Err(err).Str("repository", "GetTeams").Msg("Failed to get teams")
		return nil, err
	}

	teams := make([]*domain.Team, 0)
	for _, team := range teamsDBM {
		teams = append(teams, team.To())
	}

	return teams, nil
}

func (d *Database) GetTeamByID(id uuid.UUID) (*domain.Team, error) {
	var teamDBM dbm.Team

	res := d.First(&teamDBM, id)
	if res.Error != nil {
		log.Error().Err(res.Error).Str("repository", "GetTeamByID").Msg("Failed to get team by ID")
		return nil, res.Error
	}

	team := teamDBM.To()
	return team, nil
}

func (d *Database) UpdateTeam(team *domain.Team) (*domain.Team, error) {
	var existingTeam dbm.Team
	if err := d.First(&existingTeam, "id = ?", team.ID).Error; err != nil {
		log.Error().Err(err).Str("repository", "UpdateTeam").Msg("Failed to find team")
		return nil, err
	}

	existingTeam.Name = team.Name
	existingTeam.Description = team.Description

	if err := d.Save(&existingTeam).Error; err != nil {
		log.Error().Err(err).Str("repository", "UpdateTeam").Msg("Failed to update team")
		return nil, err
	}

	return existingTeam.To(), nil
}

func (d *Database) DeleteTeam(id uuid.UUID) error {
	res := d.Unscoped().Where("id = ?", id).Delete(&dbm.Team{})
	if res.Error != nil {
		log.Error().Err(res.Error).Str("repository", "DeleteTeam").Msg("Failed to delete team")
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("TEAM WASN'T DELETED")
	}

	return nil
}
