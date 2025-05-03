package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/Diegoes7/meetups/middleware"
	"github.com/Diegoes7/meetups/models"
)

func (d *Domain) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	if len(input.Name) < 5 {
		return nil, errors.New("name is too short, need to be at least 5 signs!")
	}
	if len(input.Description) < 7 {
		return nil, errors.New("description need to be longer")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}
	return d.MeetupRepo.CreateMeetup(meetup)
}

func (d *Domain) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	meetup, err := d.MeetupRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup not exists")
	}

	//$ object oriented
	if !meetup.IsOwner(currentUser) {
		return nil, ErrForbidden
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("Name is not long enough.")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}

	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("Description is not long enough.")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, fmt.Errorf("no update done")
	}

	meetup, err = d.MeetupRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}
	return meetup, nil
}

// DeleteMeetup implements MutationResolver.
func (d *Domain) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return false, ErrUnauthenticated
	}

	meetup, err := d.MeetupRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not exists")
	}

	//! more functional
	if !checkOwnership(meetup, currentUser) {
		return false, ErrForbidden
	}

	err = d.MeetupRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v ", err)
	}
	return true, nil
}
