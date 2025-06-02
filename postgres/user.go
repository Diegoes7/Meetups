package postgres

import (
	"fmt"

	"github.com/Diegoes7/meetups/models"
	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	DB *pg.DB
}

// ! use a generic to query the single fieldin the repo
func (u *UserRepo) GetUserByField(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

func (u *UserRepo) GetUserByID(id string) (*models.User, error) {
	return u.GetUserByField("id", id)
}

func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	return u.GetUserByField("email", email)
}

func (u *UserRepo) GetUserByUserName(username string) (*models.User, error) {
	return u.GetUserByField("username", username)
}

func (u *UserRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}

func (m *UserRepo) GetUsers() ([]*models.User, error) {
	var users []*models.User

	query := m.DB.Model(&users).Order("id")

	err := query.Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}
