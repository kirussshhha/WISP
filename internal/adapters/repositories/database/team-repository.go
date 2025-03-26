package database

import (
	dbm "WISP/internal/adapters/repositories/database/models"
	"WISP/internal/core/domain"
	"errors"

	"github.com/google/uuid"
)

func (d *Database) CreateTeam(team *domain.Team) (*domain.Team, error) {
	teamDB := dbm.NewTeamDBM(team)
	res := d.Create(teamDB)
	if res.Error != nil {
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
		return nil, err
	}

	teams := make([]*domain.Team, 0)
	for _, team := range teamsDBM {
		teams = append(teams, team.To())
	}

	return teams, nil
}

func (d *Database) UpdateTeam(team *domain.Team) (*domain.Team, error) {
	var existingTeam dbm.Team
    if err := d.First(&existingTeam, "id = ?", team.ID).Error; err != nil {
        return nil, err
    }

	existingTeam.Name = team.Name
    existingTeam.Description = team.Description

    if err := d.Save(&existingTeam).Error; err != nil {
        return nil, err
    }

	return existingTeam.To(), nil
}

func (d *Database) DeleteTeam(id uuid.UUID) error {
	res := d.Unscoped().Where("id = ?", id).Delete(&dbm.Team{})
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("TEAM WASN'T DELETED")
	}

	return nil
}