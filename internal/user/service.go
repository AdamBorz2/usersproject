package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type repository interface {
	Insert(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	SelectAll(ctx context.Context) ([]User, error)
	Select(ctx context.Context, userID uuid.UUID) (*User, error)
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

type Service struct {
	repo repository
}

func (s *Service) AddUser(ctx context.Context, user *User) (*User, error) {
	addedUser, err := s.repo.Insert(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("error in method Service.AddUser: %w", err)
	}

	return addedUser, nil

}
func (s *Service) UpdateUser(ctx context.Context, user *User) (*User, error) {
	updatedUser, err := s.repo.Update(ctx, user)

	if err != nil {
		return nil, fmt.Errorf("error in method Service.Update: %w", err)
	}

	return updatedUser, nil
}

func (s *Service) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	err := s.repo.Delete(ctx, userID)

	if err != nil {
		return fmt.Errorf("error in method Service.DeleteUser: %w", err)
	}

	return nil
}

func (s *Service) GetAllUsers(ctx context.Context) ([]User, error) {
	users, err := s.repo.SelectAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in method Service.GetAllUsers: %w", err)
	}

	return users, nil
}

func (s *Service) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	user, err := s.repo.Select(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error in method Service.GetUser: %w", err)
	}

	return user, nil
}
