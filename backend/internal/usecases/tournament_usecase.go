package usecases

import (
	"context"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/domain/repositories"

	"github.com/google/uuid"
)

type TournamentUseCase struct {
	tournamentRepo repositories.TournamentRepository
	leagueRepo     repositories.LeagueRepository
	matchRepo      repositories.MatchRepository
}

func NewTournamentUseCase(
	tournamentRepo repositories.TournamentRepository,
	leagueRepo repositories.LeagueRepository,
	matchRepo repositories.MatchRepository,
) *TournamentUseCase {
	return &TournamentUseCase{
		tournamentRepo: tournamentRepo,
		leagueRepo:     leagueRepo,
		matchRepo:      matchRepo,
	}
}

// CreateTournament creates a new tournament in a league
func (uc *TournamentUseCase) CreateTournament(ctx context.Context, leagueID uuid.UUID, name string, tournamentType entities.TournamentType) (*entities.Tournament, error) {
	// Check if league exists and can add tournaments
	league, err := uc.leagueRepo.GetByID(ctx, leagueID)
	if err != nil {
		return nil, err
	}

	if !league.CanAddTournaments() {
		return nil, entities.ErrLeagueAlreadyCompleted
	}

	// Get next tournament number
	tournamentNumber, err := uc.tournamentRepo.GetNextTournamentNumber(ctx, leagueID)
	if err != nil {
		return nil, err
	}

	// Create tournament entity
	tournament, err := entities.NewTournament(leagueID, name, tournamentType, tournamentNumber)
	if err != nil {
		return nil, err
	}

	// Save to database
	err = uc.tournamentRepo.Create(ctx, tournament)
	if err != nil {
		return nil, err
	}

	return tournament, nil
}

// GetTournament retrieves a tournament by ID
func (uc *TournamentUseCase) GetTournament(ctx context.Context, id uuid.UUID) (*entities.Tournament, error) {
	return uc.tournamentRepo.GetByID(ctx, id)
}

// GetLeagueTournaments retrieves all tournaments for a league
func (uc *TournamentUseCase) GetLeagueTournaments(ctx context.Context, leagueID uuid.UUID) ([]*entities.Tournament, error) {
	return uc.tournamentRepo.GetByLeagueID(ctx, leagueID)
}

// AddPlayerToTournament adds a player to a tournament
func (uc *TournamentUseCase) AddPlayerToTournament(ctx context.Context, tournamentID, playerID uuid.UUID) error {
	// Get tournament
	tournament, err := uc.tournamentRepo.GetByID(ctx, tournamentID)
	if err != nil {
		return err
	}

	// Check if players can be added
	if !tournament.CanAddPlayers() {
		return entities.ErrTournamentAlreadyStarted
	}

	// Check if player is already in tournament
	exists, err := uc.tournamentRepo.IsPlayerInTournament(ctx, tournamentID, playerID)
	if err != nil {
		return err
	}
	if exists {
		return nil // Already in tournament
	}

	// Add player
	return uc.tournamentRepo.AddPlayer(ctx, tournamentID, playerID, nil)
}

// StartTournament begins a tournament and generates bracket
func (uc *TournamentUseCase) StartTournament(ctx context.Context, id uuid.UUID) (*entities.Tournament, error) {
	// Get tournament
	tournament, err := uc.tournamentRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Start tournament (includes business rules)
	err = tournament.StartTournament()
	if err != nil {
		return nil, err
	}

	// Save tournament changes
	err = uc.tournamentRepo.Update(ctx, tournament)
	if err != nil {
		return nil, err
	}

	// Generate bracket (simplified)
	err = uc.generateBracket(ctx, tournament)
	if err != nil {
		return nil, err
	}

	return tournament, nil
}

// generateBracket creates initial matches for the tournament (simplified)
func (uc *TournamentUseCase) generateBracket(ctx context.Context, tournament *entities.Tournament) error {
	// This is a simplified bracket generation
	// In a real implementation, you'd have more complex logic for different tournament types
	
	// For now, just create placeholder matches
	match1 := entities.NewMatch(tournament.ID, 1, 1)
	match2 := entities.NewMatch(tournament.ID, 1, 2)
	
	// Save matches
	err := uc.matchRepo.Create(ctx, match1)
	if err != nil {
		return err
	}
	
	return uc.matchRepo.Create(ctx, match2)
}