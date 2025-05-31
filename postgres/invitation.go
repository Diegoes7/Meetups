package postgres

import (
	"fmt"
	"log"

	// "strconv"

	"github.com/Diegoes7/meetups/models"
	"github.com/go-pg/pg/v10"
)

type InvitationRepo struct {
	DB *pg.DB // Database connection
}

type InvitationRepository interface {
	IsUserInvited(meetupID string, userID int64) (bool, error)
	InviteUser(meetupID string, userID int64) (*models.Invitation, error)
	RemoveUser(meetupID string, userID int64) error
}

func (r *InvitationRepo) IsUserInvited(meetupID string, userID string) (bool, error) {
	exists, err := r.DB.Model((*models.Invitation)(nil)).
		Where("meetup_id = ? AND user_id = ?", meetupID, userID).
		Exists()

	if err != nil {
		log.Printf("⚠️ IsUserInvited failed: %v", err)
		// Treat query failure as "not invited"
		return false, nil
	}

	return exists, nil
}

func (r *InvitationRepo) InviteUser(meetupID string, userID string) (invitation *models.Invitation, err error) {
	if r.DB == nil {
		log.Fatal("💥 r.DB is nil in InviteUser")
	}
	var id string

	//! Insert and get back the generated ID (assuming it's UUID or SERIAL stored as string)
	_, err = r.DB.QueryOne(
		pg.Scan(&id),
		`INSERT INTO meetup_invitations (meetup_id, user_id, status)
		 VALUES (?, ?, ?) RETURNING id`,
		meetupID,
		userID,
		"pending",
	)
	if err != nil {
		return nil, err
	}

	invitation = &models.Invitation{
		ID:     id,
		Status: models.InvitationStatus("pending"), // cast if needed
		// Meetup: &models.Meetup{ID: meetupID},       //* lightweight reference
		// UserID:   &models.User{ID: userID},           //* lightweight reference
		MeetupID: meetupID,
		UserID:   userID,
	}

	return invitation, nil
}

func (r *InvitationRepo) RemoveUser(input models.InviteUserInput) (*models.User, error) {
	// Optional: check if the invitation exists
	var user models.User

	// Fetch user to return later
	err := r.DB.Model(&user).
		Where("id = ?", input.UserID).
		Select()
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Delete the invitation
	_, err = r.DB.Model((*models.Invitation)(nil)).
		Where("meetup_id = ? AND user_id = ?", input.MeetupID, input.UserID).
		Delete()
	if err != nil {
		return nil, fmt.Errorf("failed to remove invitation: %w", err)
	}

	return &user, nil
}

func (r *InvitationRepo) GetInvitedUsersByMeetupID(meetupID string) ([]*models.User, error) {
	var invitations []models.Invitation
	err := r.DB.Model(&invitations).
		Where("meetup_id = ?", meetupID).
		Select()
	if err != nil {
		return nil, err
	}

	// collect user IDs
	var userIDs []string
	for _, inv := range invitations {
		userIDs = append(userIDs, inv.UserID)
	}

	if len(userIDs) == 0 {
		return []*models.User{}, nil
	}

	// load users by those IDs
	var users []*models.User
	err = r.DB.Model(&users).
		Where("id IN (?)", pg.In(userIDs)).
		Select()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *InvitationRepo) LeaveUserFromMeetup(meetupID, userID string) error {
	// Delete the invitation row
	_, err := r.DB.Model((*models.Invitation)(nil)).
		Where("meetup_id = ? AND user_id = ?", meetupID, userID).
		Delete()
	if err != nil {
		return fmt.Errorf("failed to remove user from meetup: %w", err)
	}
	return nil
}

func (d *InvitationRepo) GetMeetupsUserIsInvitedTo(userID string) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := d.DB.Model(&meetups).
		Join("JOIN invitations AS i ON i.meetup_id = meetup.id").
		Where("i.user_id = ?", userID).
		Select()

	return meetups, err
}

func (r *InvitationRepo) GetInvitations(filter *models.InvitationFilter, limit, offset *int32) ([]*models.Invitation, error) {
	var invitations []*models.Invitation

	query := r.DB.Model(&invitations).Order("id ASC")

	if filter != nil {
		if filter.UserID != "" {
			query = query.Where("user_id = ?", filter.UserID)
		}
		if filter.Status != nil && *filter.Status != "" {
			query = query.Where("status = ?", *filter.Status)
		}
	}

	if limit != nil {
		query.Limit(int(*limit))
	}

	if offset != nil {
		query.Offset(int(*offset))
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return invitations, nil
}
