package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/Diegoes7/meetups/middleware"
	"github.com/Diegoes7/meetups/models"
)

func (d *Domain) GetMessagesByMeetup(ctx context.Context, meetupID string, limit *int32, offset *int32) ([]*models.Message, error) {
	meetup, err := d.MeetupRepo.GetByID(meetupID)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exist")
	}
	messages, err := d.MessageRepo.GetMessagesByMeetup(meetupID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}

	return messages, nil
}

func (d *Domain) SendMessage(ctx context.Context, input models.NewMessageInput) (*models.Message, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	loginUserID := currentUser.ID
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	// Step 1: Validate meetup exists
	meetup, err := d.MeetupRepo.GetByID(input.MeetupID)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exist")
	}

	// Step 2: Check if user is allowed to send message (e.g., is invited or is the owner)
	allowed, err := d.InvitationRepo.IsUserInvited(input.MeetupID, loginUserID)
	if err != nil {
		return nil, fmt.Errorf("failed to check invitation: %w", err)
	}
	if !allowed && meetup.UserID != loginUserID {
		return nil, errors.New("user is not authorized to send messages to this meetup")
	}

	// Step 3: Create message in DB
	newMessage, err := d.MessageRepo.CreateMessage(&input, loginUserID)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	// Step 4: Return created message
	return newMessage, nil

}

// EditMessage is the resolver for the editMessage field.
func (d *Domain) EditMessage(ctx context.Context, input models.UpdateMessageInput) (*models.Message, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	loginUserID := currentUser.ID
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	// Step 1: Retrieve the message by ID
	message, err := d.MessageRepo.GetMessageByID(input.MessageID)
	if err != nil || message == nil {
		return nil, errors.New("message not found")
	}

	// Step 2: Verify ownership — only the sender can edit their message
	if message.SenderID != loginUserID {
		return nil, errors.New("not authorized to edit this message")
	}

	// Step 3: Update message content
	updatedMessage, err := d.MessageRepo.UpdateContent(input.MessageID, input.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to update message: %w", err)
	}

	// Step 4: Return updated message
	return updatedMessage, nil
}

// DeleteMessage is the resolver for the deleteMessage field.
func (d *Domain) DeleteMessage(ctx context.Context, messageID string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	loginUserID := currentUser.ID
	if err != nil {
		return false, fmt.Errorf("failed to get current user: %w", err)
	}

	// Step 1: Retrieve the message
	message, err := d.MessageRepo.GetMessageByID(messageID)
	if err != nil || message == nil {
		return false, errors.New("message not found")
	}

	// Step 2: Verify ownership — only sender or meetup owner can delete
	meetup, err := d.MeetupRepo.GetByID(message.MeetupID)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not found")
	}

	if message.SenderID != loginUserID && meetup.UserID != loginUserID {
		return false, errors.New("not authorized to delete this message")
	}

	// Step 3: Delete message from DB
	err = d.MessageRepo.Delete(message)
	if err != nil {
		return false, fmt.Errorf("failed to delete message: %w", err)
	}

	// Step 4: Return success
	return true, nil
}
