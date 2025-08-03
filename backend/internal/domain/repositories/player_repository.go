package repositories

import (
	"context"
	"github.com/google/uuid"
	"darts-league-backend/internal/domain/entities"
)

type PlayerRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, player *entities.Player) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Player, error)
	GetByEmail(ctx context.Context, email string) (*entities.Player, error)
	Update(ctx context.Context, player *entities.Player) error
	Delete(ctx context.Context, id uuid.UUID) error

	// Queries
	GetAll(ctx context.Context, limit, offset int) ([]*entities.Player, error)
	SearchByName(ctx context.Context, name string, limit int) ([]*entities.Player, error)
	GetByIDs(ctx context.Context, ids []uuid.UUID) ([]*entities.Player, error)

	// League-specific queries
	GetLeaguePlayers(ctx context.Context, leagueID uuid.UUID) ([]*entities.Player, error)
	GetTournamentPlayers(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Player, error)

	// Statistics
	GetPlayerCount(ctx context.Context) (int64, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}