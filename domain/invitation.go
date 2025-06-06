package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/Diegoes7/meetups/middleware"
	"github.com/Diegoes7/meetups/models"
)

func (d *Domain) InviteUserToMeetup(ctx context.Context, meetupID string, userID string) (*models.Invitation, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	meetup, err := d.MeetupRepo.GetByID(meetupID)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exist")
	}

	if !meetup.IsOwner(currentUser) {
		return nil, ErrForbidden
	}

	exists, err := d.InvitationRepo.IsUserInvited(meetupID, userID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user already invited")
	}

	return d.InvitationRepo.InviteUser(meetupID, userID)
}

func (d *Domain) RemoveUserFromMeetup(ctx context.Context, input models.InviteUserInput, loginUserID string) (*models.User, error) {
	// currentUser, err := middleware.GetCurrentUserFromCTX(ctx)

	meetup, err := d.MeetupRepo.GetByID(input.MeetupID)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exist")
	}

	if meetup.UserID == loginUserID {
		user, err := d.InvitationRepo.RemoveUser(input)
		if err != nil {
			return nil, fmt.Errorf("failed to remove user: %w", err)
		}

		return user, nil
	}

	return nil, errors.New("Do not have permission to do this.")

	// if !meetup.IsOwner(currentUser) {
	// 	return nil, ErrForbidden
	// }

	// Remove the user invitation AND get the removed user from the repo
}

func (d *Domain) GetMeetupUsersInvited(ctx context.Context, meetupID string) ([]*models.User, error) {
	// Parse meetupID string to int64 (since your DB uses BIGINT)
	// parsedID, err := strconv.ParseInt(meetupID, 10, 64)
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid meetupID")
	// }

	users, err := d.InvitationRepo.GetInvitedUsersByMeetupID(meetupID)
	if err != nil {
		return nil, fmt.Errorf("could not get invited users: %w", err)
	}

	return users, nil
}

func (d *Domain) LeaveMeetup(ctx context.Context, meetupID string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return false, ErrUnauthenticated
	}

	// Optional: check if the meetup exists
	meetup, err := d.MeetupRepo.GetByID(meetupID)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not exists")
	}

	// Ensure user is not the owner
	if meetup.UserID == currentUser.ID {
		return false, errors.New("owner cannot leave their own meetup")
	}

	// Call the invitation repo to remove the user
	err = d.InvitationRepo.LeaveUserFromMeetup(meetupID, currentUser.ID)
	if err != nil {
		return false, fmt.Errorf("failed to leave meetup: %w", err)
	}

	return true, nil
}

func (d *Domain) GetInvitations(filter *models.InvitationFilter, limit, offset *int32) ([]*models.Invitation, error) {
	return d.InvitationRepo.GetInvitations(filter, limit, offset)
}

func (d *Domain) AcceptInvitation(ctx context.Context, invitationID string) (*models.Invitation, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}
	userID := currentUser.ID

	return d.InvitationRepo.AcceptInvitation(userID, invitationID)
}

func (d *Domain) DeclineInvitation(ctx context.Context, invitationID string) (*models.Invitation, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, ErrUnauthenticated
	}
	userID := currentUser.ID

	return d.InvitationRepo.DeclineInvitation(userID, invitationID)
}
