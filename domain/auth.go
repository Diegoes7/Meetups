package domain

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Diegoes7/meetups/models"
)

// Login implements MutationResolver.
func (d *Domain) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := d.UserRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrBadCredential
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredential
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (d *Domain) Logout(ctx context.Context, userID string) (*models.User, error) {
	user, err := d.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	// No token blacklist or session invalidation here
	return user, nil
}

// Register is the resolver for the register field.
func (d *Domain) Register(ctx context.Context, input *models.RegisterArgs) (*models.AuthResponse, error) {
	_, err := d.UserRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	_, err = d.UserRepo.GetUserByUserName(input.Username)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	//! TODO: create a verification code

	tx, err := d.UserRepo.DB.Begin()
	if err != nil {
		log.Printf("error creating transaction: %v", err)
		return nil, errors.New("something went wrong")
	}
	defer tx.Rollback()

	if _, err := d.UserRepo.CreateUser(tx, user); err != nil {
		log.Printf("error creating user: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error while commiting: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	authResponse := &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}

	return authResponse, nil
}
