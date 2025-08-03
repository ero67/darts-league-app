package repositories

import (
	"context"
	"database/sql"
)

// BaseRepository defines common database operations
type BaseRepository interface {
	// Transaction management
	BeginTx(ctx context.Context) (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
	RollbackTx(tx *sql.Tx) error

	// Health check
	Ping(ctx context.Context) error

	// Connection management
	Close() error
}

// UnitOfWork represents a collection of repository operations that should be executed atomically
type UnitOfWork interface {
	// Repository access within transaction
	Players() PlayerRepository
	Leagues() LeagueRepository
	Tournaments() TournamentRepository
	Matches() MatchRepository
	Standings() LeagueStandingsRepository
	Statistics() StatisticsRepository

	// Transaction control
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}