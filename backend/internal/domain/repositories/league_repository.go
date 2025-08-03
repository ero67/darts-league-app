package repositories

import (
	"context"
	"time"
	"github.com/google/uuid"
	"darts-league-backend/internal/domain/entities"
)

type LeagueRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, league *entities.League) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.League, error)
	Update(ctx context.Context, league *entities.League) error
	Delete(ctx context.Context, id uuid.UUID) error

	// Queries
	GetAll(ctx context.Context, limit, offset int) ([]*entities.League, error)
	GetByStatus(ctx context.Context, status entities.LeagueStatus, limit, offset int) ([]*entities.League, error)
	GetBySeason(ctx context.Context, season string) ([]*entities.League, error)
	GetActive(ctx context.Context) ([]*entities.League, error)
	SearchByName(ctx context.Context, name string, limit int) ([]*entities.League, error)

	// Player management
	AddPlayer(ctx context.Context, leagueID, playerID uuid.UUID) error
	RemovePlayer(ctx context.Context, leagueID, playerID uuid.UUID) error
	IsPlayerInLeague(ctx context.Context, leagueID, playerID uuid.UUID) (bool, error)
	GetLeaguePlayerCount(ctx context.Context, leagueID uuid.UUID) (int, error)

	// Date queries
	GetLeaguesInDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entities.League, error)
	GetCurrentLeagues(ctx context.Context) ([]*entities.League, error)

	// Statistics
	GetLeagueCount(ctx context.Context) (int64, error)
	GetLeagueCountByStatus(ctx context.Context, status entities.LeagueStatus) (int64, error)
}