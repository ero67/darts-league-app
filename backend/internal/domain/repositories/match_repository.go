package repositories

import (
	"context"
	"github.com/google/uuid"
	"darts-league-backend/internal/domain/entities"
)

type MatchRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, match *entities.Match) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Match, error)
	Update(ctx context.Context, match *entities.Match) error
	Delete(ctx context.Context, id uuid.UUID) error

	// Queries
	GetAll(ctx context.Context, limit, offset int) ([]*entities.Match, error)
	GetByTournamentID(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Match, error)
	GetByPlayerID(ctx context.Context, playerID uuid.UUID, limit, offset int) ([]*entities.Match, error)
	GetByStatus(ctx context.Context, status entities.MatchStatus, limit, offset int) ([]*entities.Match, error)

	// Tournament-specific queries
	GetTournamentMatches(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Match, error)
	GetMatchesByRound(ctx context.Context, tournamentID uuid.UUID, round int) ([]*entities.Match, error)
	GetCurrentMatches(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Match, error)
	GetCompletedMatches(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Match, error)

	// Player-specific queries
	GetPlayerMatches(ctx context.Context, playerID uuid.UUID, tournamentID *uuid.UUID) ([]*entities.Match, error)
	GetPlayerMatchesInLeague(ctx context.Context, playerID, leagueID uuid.UUID) ([]*entities.Match, error)
	GetLiveMatchesForPlayer(ctx context.Context, playerID uuid.UUID) ([]*entities.Match, error)

	// Bracket generation helpers
	CreateBracketMatches(ctx context.Context, matches []*entities.Match) error
	GetNextMatch(ctx context.Context, tournamentID uuid.UUID, round int) (*entities.Match, error)
	GetMaxRound(ctx context.Context, tournamentID uuid.UUID) (int, error)

	// Statistics
	GetMatchCount(ctx context.Context) (int64, error)
	GetMatchCountByTournament(ctx context.Context, tournamentID uuid.UUID) (int64, error)
	GetMatchCountByPlayer(ctx context.Context, playerID uuid.UUID) (int64, error)
	GetMatchCountByStatus(ctx context.Context, status entities.MatchStatus) (int64, error)
}