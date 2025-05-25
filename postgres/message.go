package postgres

import (
	"fmt"
	"log"

	"github.com/Diegoes7/meetups/models"
	"github.com/go-pg/pg/v10"
)

type MessageRepo struct {
	DB *pg.DB
}

type MessageRepository interface {
	CreateMessage(input *models.NewMessageInput, senderID string) (*models.Message, error)
	GetMessageByID(messageID string) (*models.Message, error)
	UpdateContent(messageID string, newContent string) (*models.Message, error)
	GetMessagesByMeetup(meetupID string, limit *int32, offset *int32) ([]*models.Message, error)
	Delete(message *models.Message) error
}

// func (r *MessageRepo) GetMessagesByMeetup(meetupID string, limit *int32, offset *int32) ([]*models.Message, error) {
// 	var messages []*models.Message

// 	query := r.DB.Model(&messages).
// 		Where("meetup_id = ?", meetupID).
// 		Order("timestamp ASC") // or DESC depending on your needs

// 	if limit != nil {
// 		query = query.Limit(int(*limit))
// 	}
// 	if offset != nil {
// 		query = query.Offset(int(*offset))
// 	}

// 	err := query.Select()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return messages, nil
// }

func (r *MessageRepo) GetMessagesByMeetup(meetupID string, limit *int32, offset *int32) ([]*models.Message, error) {
	var messages []*models.Message

	query := r.DB.Model(&messages).
		Relation("Sender"). // join the Sender user table
		Where("meetup_id = ?", meetupID).
		Order("timestamp ASC")

	if limit != nil {
		query = query.Limit(int(*limit))
	}
	if offset != nil {
		query = query.Offset(int(*offset))
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// CreateMessage inserts a new message and returns the created message with ID.
func (r *MessageRepo) CreateMessage(input *models.NewMessageInput, senderID string) (*models.Message, error) {
	if r.DB == nil {
		log.Fatal("r.DB is nil in CreateMessage")
	}

	var id string
	// Assuming your messages table has columns: id (PK), content, sender_id, created_at, updated_at
	_, err := r.DB.QueryOne(
		pg.Scan(&id),
		`INSERT INTO messages (content, sender_id, meetup_id) VALUES (?, ?, ?) RETURNING id`,
		input.Content,
		senderID,
		input.MeetupID, // <- Pass this fields
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert message: %w", err)
	}

	// Return the created message minimally or load the full entity if needed
	message := &models.Message{
		ID:       id,
		Content:  input.Content,
		SenderID: senderID,
		MeetupID: input.MeetupID,
	}

	return message, nil
}

// GetMessageByID fetches a message by its ID.
func (r *MessageRepo) GetMessageByID(messageID string) (*models.Message, error) {
	var message models.Message
	err := r.DB.Model(&message).
		Where("id = ?", messageID).
		Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil // or a custom NotFound error
		}
		return nil, fmt.Errorf("failed to get message by id: %w", err)
	}
	return &message, nil
}

// UpdateContent updates the content of a message by its ID and returns the updated message.
func (r *MessageRepo) UpdateContent(messageID string, newContent string) (*models.Message, error) {
	var message models.Message
	_, err := r.DB.Model(&message).
		Column("content").
		Where("id = ?", messageID).
		Update()
	if err != nil {
		return nil, fmt.Errorf("failed to update message content: %w", err)
	}

	// After update, fetch updated message to return
	messagePtr, err := r.GetMessageByID(messageID)
	if err != nil {
		return nil, err
	}
	return messagePtr, nil
}

func (r *MessageRepo) Delete(message *models.Message) error {
	_, err := r.DB.Model(message).
		Where("id = ?", message.ID).
		Delete()
	return err
}
