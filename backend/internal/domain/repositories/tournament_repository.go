package repositories

import (
	"context"
	"time"
	"github.com/google/uuid"
	"darts-league-backend/internal/domain/entities"
)

type TournamentRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, tournament *entities.Tournament) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Tournament, error)
	Update(ctx context.Context, tournament *entities.Tournament) error
	Delete(ctx context.Context, id uuid.UUID) error

	// Queries
	GetAll(ctx context.Context, limit, offset int) ([]*entities.Tournament, error)
	GetByLeagueID(ctx context.Context, leagueID uuid.UUID) ([]*entities.Tournament, error)
	GetByStatus(ctx context.Context, status entities.TournamentStatus, limit, offset int) ([]*entities.Tournament, error)
	GetByType(ctx context.Context, tournamentType entities.TournamentType, limit, offset int) ([]*entities.Tournament, error)
	SearchByName(ctx context.Context, name string, limit int) ([]*entities.Tournament, error)

	// League-specific queries
	GetLeagueTournaments(ctx context.Context, leagueID uuid.UUID, limit, offset int) ([]*entities.Tournament, error)
	GetNextTournamentNumber(ctx context.Context, leagueID uuid.UUID) (int, error)
	GetLatestTournament(ctx context.Context, leagueID uuid.UUID) (*entities.Tournament, error)
	GetCompletedTournaments(ctx context.Context, leagueID uuid.UUID) ([]*entities.Tournament, error)

	// Player management
	AddPlayer(ctx context.Context, tournamentID, playerID uuid.UUID, seed *int) error
	RemovePlayer(ctx context.Context, tournamentID, playerID uuid.UUID) error
	IsPlayerInTournament(ctx context.Context, tournamentID, playerID uuid.UUID) (bool, error)
	GetTournamentPlayerCount(ctx context.Context, tournamentID uuid.UUID) (int, error)
	SetPlayerPosition(ctx context.Context, tournamentID, playerID uuid.UUID, position int, points int) error

	// Date queries
	GetTournamentsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entities.Tournament, error)
	GetTournamentsScheduledFor(ctx context.Context, date time.Time) ([]*entities.Tournament, error)
	GetUpcomingTournaments(ctx context.Context, limit int) ([]*entities.Tournament, error)

	// Statistics
	GetTournamentCount(ctx context.Context) (int64, error)
	GetTournamentCountByLeague(ctx context.Context, leagueID uuid.UUID) (int64, error)
	GetTournamentCountByStatus(ctx context.Context, status entities.TournamentStatus) (int64, error)
}