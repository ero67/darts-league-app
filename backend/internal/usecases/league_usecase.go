package usecases

import (
	"context"
	"github.com/google/uuid"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/domain/repositories"
)

type LeagueUseCase struct {
	leagueRepo   repositories.LeagueRepository
	standingsRepo repositories.LeagueStandingsRepository
}

func NewLeagueUseCase(leagueRepo repositories.LeagueRepository, standingsRepo repositories.LeagueStandingsRepository) *LeagueUseCase {
	return &LeagueUseCase{
		leagueRepo:   leagueRepo,
		standingsRepo: standingsRepo,
	}
}

// CreateLeague creates a new league
func (uc *LeagueUseCase) CreateLeague(ctx context.Context, name, description, season string) (*entities.League, error) {
	// Create league entity (includes validation)
	league, err := entities.NewLeague(name, description, season)
	if err != nil {
		return nil, err
	}

	// Save to database
	err = uc.leagueRepo.Create(ctx, league)
	if err != nil {
		return nil, err
	}

	return league, nil
}

// GetLeague retrieves a league by ID
func (uc *LeagueUseCase) GetLeague(ctx context.Context, id uuid.UUID) (*entities.League, error) {
	return uc.leagueRepo.GetByID(ctx, id)
}

// GetAllLeagues retrieves all leagues
func (uc *LeagueUseCase) GetAllLeagues(ctx context.Context, limit, offset int) ([]*entities.League, error) {
	return uc.leagueRepo.GetAll(ctx, limit, offset)
}

// AddPlayerToLeague adds a player to a league
func (uc *LeagueUseCase) AddPlayerToLeague(ctx context.Context, leagueID, playerID uuid.UUID) error {
	// Check if player is already in league
	exists, err := uc.leagueRepo.IsPlayerInLeague(ctx, leagueID, playerID)
	if err != nil {
		return err
	}
	if exists {
		return nil // Already in league, no error
	}

	// Add player to league
	err = uc.leagueRepo.AddPlayer(ctx, leagueID, playerID)
	if err != nil {
		return err
	}

	// Create league standings entry
	return uc.standingsRepo.Create(ctx, leagueID, playerID)
}

// GetLeagueStandings retrieves the current league standings
func (uc *LeagueUseCase) GetLeagueStandings(ctx context.Context, leagueID uuid.UUID) ([]*repositories.LeagueStanding, error) {
	return uc.standingsRepo.GetLeagueStandings(ctx, leagueID)
}

// StartLeague activates a league
func (uc *LeagueUseCase) StartLeague(ctx context.Context, id uuid.UUID) (*entities.League, error) {
	// Get league
	league, err := uc.leagueRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Start league (includes business rules)
	err = league.StartLeague()
	if err != nil {
		return nil, err
	}

	// Save changes
	err = uc.leagueRepo.Update(ctx, league)
	if err != nil {
		return nil, err
	}

	return league, nil
}