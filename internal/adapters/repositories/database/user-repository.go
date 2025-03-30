package database

import (
	dbm "WISP/internal/adapters/repositories/database/models"
	"WISP/internal/core/domain"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (d *Database) CreateUser(user *domain.User) (*domain.User, error) {
	userDB := dbm.NewUserDBM(user)
	res := d.Create(userDB)
	if res.Error != nil {
		log.Error().Err(res.Error).Str("repository", "CreateUser").Msg("Failed to create user")
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		log.Error().Str("repository", "CreateUser").Msg("User wasn't created")
		return nil, errors.New("USER WASN'T CREATED")
	}

	user = userDB.To()
	return user, nil
}

func (d *Database) GetUsers() ([]*domain.User, error) {
	var usersDBM []*dbm.User

	err := d.Find(&usersDBM).Error
	if err != nil {
		log.Error().Err(err).Str("repository", "GetUsers").Msg("Failed to get users")
		return nil, err
	}

	users := make([]*domain.User, 0)
	for _, user := range usersDBM {
		users = append(users, user.To())
	}

	return users, nil
}

func (d *Database) GetUserByEmail(email string) (*domain.User, error) {
	var userDBM dbm.User

	err := d.Where("email = ?", email).First(&userDBM).Error
	if err != nil {
		log.Error().Err(err).Str("repository", "GetUserByEmail").Msg("Failed to get user by email")
		return nil, err
	}

	user := userDBM.To()
	return user, nil
}

func (d *Database) GetUserByID(id uuid.UUID) (*domain.User, error) {
	var userDBM dbm.User

	res := d.First(&userDBM, id)
	if res.Error != nil {
		log.Error().Err(res.Error).Str("repository", "GetUserByID").Msg("Failed to get user by ID")
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("USER WASN'T FOUNDED")
	}

	user := userDBM.To()
	return user, nil
}

func (d *Database) UpdateUser(user *domain.User) (*domain.User, error) {
	var existingUser dbm.User
	if err := d.First(&existingUser, "id = ?", user.ID).Error; err != nil {
		log.Error().Err(err).Str("repository", "UpdateUser").Msg("Failed to find user")
		return nil, err
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := d.Save(&existingUser).Error; err != nil {
		log.Error().Err(err).Str("repository", "UpdateUser").Msg("Failed to update user")
		return nil, err
	}

	return existingUser.To(), nil
}

func (d *Database) DeleteUser(id uuid.UUID) error {
	res := d.Unscoped().Where("id = ?", id).Delete(&dbm.User{})
	if res.Error != nil {
		log.Error().Err(res.Error).Str("repository", "DeleteUser").Msg("Failed to delete user")
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("USER WASN'T DELETED")
	}

	return nil
}
