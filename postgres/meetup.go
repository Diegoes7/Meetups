package postgres

import (
	"fmt"

	"github.com/Diegoes7/meetups/models"
	"github.com/go-pg/pg/v10"
)

type MeetupRepo struct {
	DB *pg.DB
}

func (m MeetupRepo) GetMeetupsForUser(user *models.User) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := m.DB.Model(&meetups).Where("user_id = ?", user.ID).Order("id").Select()
	return meetups, err
}

func (m *MeetupRepo) GetMeetups(filter *models.MeetupsFilter, limit, offset *int32) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	query := m.DB.Model(&meetups).Order("id")

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name ILIKE ? ", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	if limit != nil {
		query.Limit(int(*limit))
	}

	if limit != nil {
		query.Offset(int(*offset))
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetupRepo) GetMeetup(meetupID string) (*models.Meetup, error) {
	var meetup models.Meetup

	err := m.DB.Model(&meetup).
		Where("id = ?", meetupID).
		Limit(1).
		Select()

	if err != nil {
		return nil, err
	}

	return &meetup, nil
}

func (m *MeetupRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()

	return meetup, err
}

func (m *MeetupRepo) GetByID(id string) (*models.Meetup, error) {
	var meetup models.Meetup
	err := m.DB.Model(&meetup).Where("id = ?", id).First()
	return &meetup, err
}

func (m *MeetupRepo) Update(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Update()
	return meetup, err
}

func (m MeetupRepo) Delete(meetup *models.Meetup) error {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Delete()
	return err
}
