package usecases

import (
	"context"
	"github.com/google/uuid"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/domain/repositories"
)

type MatchUseCase struct {
	matchRepo     repositories.MatchRepository
	standingsRepo repositories.LeagueStandingsRepository
}

func NewMatchUseCase(matchRepo repositories.MatchRepository, standingsRepo repositories.LeagueStandingsRepository) *MatchUseCase {
	return &MatchUseCase{
		matchRepo:     matchRepo,
		standingsRepo: standingsRepo,
	}
}

// CreateMatch creates a new match
func (uc *MatchUseCase) CreateMatch(ctx context.Context, match *entities.Match) error {
	return uc.matchRepo.Create(ctx, match)
}

// GetMatch retrieves a match by ID
func (uc *MatchUseCase) GetMatch(ctx context.Context, id uuid.UUID) (*entities.Match, error) {
	return uc.matchRepo.GetByID(ctx, id)
}

// GetTournamentMatches retrieves all matches for a tournament
func (uc *MatchUseCase) GetTournamentMatches(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Match, error) {
	return uc.matchRepo.GetByTournamentID(ctx, tournamentID)
}

// StartMatch begins a match
func (uc *MatchUseCase) StartMatch(ctx context.Context, matchID uuid.UUID, player1ID, player2ID uuid.UUID) (*entities.Match, error) {
	// Get match
	match, err := uc.matchRepo.GetByID(ctx, matchID)
	if err != nil {
		return nil, err
	}

	// Set players
	err = match.SetPlayers(player1ID, player2ID)
	if err != nil {
		return nil, err
	}

	// Start match
	err = match.StartMatch()
	if err != nil {
		return nil, err
	}

	// Save changes
	err = uc.matchRepo.Update(ctx, match)
	if err != nil {
		return nil, err
	}

	return match, nil
}

// UpdateMatchScore updates the score of an ongoing match
func (uc *MatchUseCase) UpdateMatchScore(ctx context.Context, matchID uuid.UUID, player1Score, player2Score int) (*entities.Match, error) {
	// Get match
	match, err := uc.matchRepo.GetByID(ctx, matchID)
	if err != nil {
		return nil, err
	}

	// Update score
	err = match.UpdateScore(player1Score, player2Score)
	if err != nil {
		return nil, err
	}

	// Save changes
	err = uc.matchRepo.Update(ctx, match)
	if err != nil {
		return nil, err
	}

	return match, nil
}

// CompleteMatch finishes a match and updates standings
func (uc *MatchUseCase) CompleteMatch(ctx context.Context, matchID uuid.UUID, winnerID uuid.UUID) (*entities.Match, error) {
	// Get match
	match, err := uc.matchRepo.GetByID(ctx, matchID)
	if err != nil {
		return nil, err
	}

	// Complete match
	err = match.CompleteMatch(winnerID)
	if err != nil {
		return nil, err
	}

	// Save changes
	err = uc.matchRepo.Update(ctx, match)
	if err != nil {
		return nil, err
	}

	// TODO: Update tournament progression and league standings
	// This would be more complex in a real implementation

	return match, nil
}

// GetPlayerMatches retrieves matches for a specific player
func (uc *MatchUseCase) GetPlayerMatches(ctx context.Context, playerID uuid.UUID, limit, offset int) ([]*entities.Match, error) {
	return uc.matchRepo.GetByPlayerID(ctx, playerID, limit, offset)
}