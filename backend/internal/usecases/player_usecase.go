package usecases

import (
	"context"
	"github.com/google/uuid"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/domain/repositories"
)

type PlayerUseCase struct {
	playerRepo repositories.PlayerRepository
}

func NewPlayerUseCase(playerRepo repositories.PlayerRepository) *PlayerUseCase {
	return &PlayerUseCase{
		playerRepo: playerRepo,
	}
}

// CreatePlayer creates a new player
func (uc *PlayerUseCase) CreatePlayer(ctx context.Context, name string, email, nickname *string) (*entities.Player, error) {
	// Create player entity (includes validation)
	player, err := entities.NewPlayer(name, email, nickname)
	if err != nil {
		return nil, err
	}

	// Check if email already exists
	if email != nil && *email != "" {
		exists, err := uc.playerRepo.ExistsByEmail(ctx, *email)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, entities.ErrPlayerNotFound // reuse error for simplicity
		}
	}

	// Save to database
	err = uc.playerRepo.Create(ctx, player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// GetPlayer retrieves a player by ID
func (uc *PlayerUseCase) GetPlayer(ctx context.Context, id uuid.UUID) (*entities.Player, error) {
	return uc.playerRepo.GetByID(ctx, id)
}

// GetAllPlayers retrieves all players with pagination
func (uc *PlayerUseCase) GetAllPlayers(ctx context.Context, limit, offset int) ([]*entities.Player, error) {
	return uc.playerRepo.GetAll(ctx, limit, offset)
}

// UpdatePlayer updates player information
func (uc *PlayerUseCase) UpdatePlayer(ctx context.Context, id uuid.UUID, name string, email, nickname, avatarURL *string) (*entities.Player, error) {
	// Get existing player
	player, err := uc.playerRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update player (includes validation)
	err = player.UpdateProfile(name, email, nickname, avatarURL)
	if err != nil {
		return nil, err
	}

	// Save changes
	err = uc.playerRepo.Update(ctx, player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

// DeletePlayer removes a player
func (uc *PlayerUseCase) DeletePlayer(ctx context.Context, id uuid.UUID) error {
	return uc.playerRepo.Delete(ctx, id)
}