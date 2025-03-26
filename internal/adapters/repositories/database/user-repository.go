package database

import (
	dbm "WISP/internal/adapters/repositories/database/models"
	"WISP/internal/core/domain"
	"errors"

	"github.com/google/uuid"
)

func (d *Database) CreateUser(user *domain.User) (*domain.User, error) {
	userDB := dbm.NewUserDBM(user)
	res := d.Create(userDB)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("USER WASN'T CREATED")
	}

	user = userDB.To()
	return user, nil
}

func (d *Database) GetUserByEmail(email string) (*domain.User, error) {
	var userDBM dbm.User

	err := d.Where("email = ?", email).First(&userDBM).Error
	if err != nil {
		return nil, err
	}

	user := userDBM.To()
	return user, nil
}

func (d *Database) GetUserByID(id uuid.UUID) (*domain.User, error) {
	var userDBM dbm.User

	res := d.First(&userDBM, id)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("USER WASN'T FOUNDED")
	}

	user := userDBM.To()
	return user, nil
}

func (d *Database) DeleteUser(id uuid.UUID) error {
	res := d.Unscoped().Where("id = ?", id).Delete(&dbm.User{})
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("USER WASN'T DELETED")
	}

	return nil
}
