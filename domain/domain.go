package domain

import (
	"errors"

	"github.com/Diegoes7/meetups/models"
	"github.com/Diegoes7/meetups/postgres"
)

var (
	ErrBadCredential   = errors.New("email/password doesn't match(work)")
	ErrUnauthenticated = errors.New("No authenticated user")
	ErrForbidden       = errors.New("unauthorized, not a owner of the meetup")
)

type Domain struct {
	UserRepo       postgres.UserRepo
	MeetupRepo     postgres.MeetupRepo
	InvitationRepo postgres.InvitationRepo
	MessageRepo    postgres.MessageRepo
}

func NewDomain(userRepo postgres.UserRepo, meetupRepo postgres.MeetupRepo, invitationRepo postgres.InvitationRepo, messageRepo postgres.MessageRepo) *Domain {
	return &Domain{
		UserRepo:       userRepo,
		MeetupRepo:     meetupRepo,
		InvitationRepo: invitationRepo,
		MessageRepo:    messageRepo,
	}
}

type Ownable interface {
	IsOwner(user *models.User) bool
}

func checkOwnership(owner Ownable, user *models.User) bool {
	return owner.IsOwner(user)
}
