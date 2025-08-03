package repositories

import (
	"context"
	"github.com/google/uuid"
)

// LeagueStanding represents a player's position in a league
type LeagueStanding struct {
	ID                  uuid.UUID `json:"id"`
	LeagueID            uuid.UUID `json:"league_id"`
	PlayerID            uuid.UUID `json:"player_id"`
	PlayerName          string    `json:"player_name"`
	PlayerNickname      *string   `json:"player_nickname,omitempty"`
	TotalPoints         int       `json:"total_points"`
	TournamentsPlayed   int       `json:"tournaments_played"`
	TournamentsWon      int       `json:"tournaments_won"`
	FinalsReached       int       `json:"finals_reached"`
	SemiFinalsReached   int       `json:"semi_finals_reached"`
	CurrentPosition     int       `json:"current_position"`
	PreviousPosition    int       `json:"previous_position"`
	PositionChange      int       `json:"position_change"` // calculated field
}

type LeagueStandingsRepository interface {
	// Basic operations
	Create(ctx context.Context, leagueID, playerID uuid.UUID) error
	GetByLeagueAndPlayer(ctx context.Context, leagueID, playerID uuid.UUID) (*LeagueStanding, error)
	Update(ctx context.Context, standing *LeagueStanding) error
	Delete(ctx context.Context, leagueID, playerID uuid.UUID) error

	// League standings queries
	GetLeagueStandings(ctx context.Context, leagueID uuid.UUID) ([]*LeagueStanding, error)
	GetTopPlayers(ctx context.Context, leagueID uuid.UUID, limit int) ([]*LeagueStanding, error)
	GetPlayerPosition(ctx context.Context, leagueID, playerID uuid.UUID) (int, error)

	// Points management
	AddPoints(ctx context.Context, leagueID, playerID uuid.UUID, points int) error
	UpdateTournamentStats(ctx context.Context, leagueID, playerID uuid.UUID, position int) error
	RecalculatePositions(ctx context.Context, leagueID uuid.UUID) error

	// Statistics
	GetStandingsCount(ctx context.Context, leagueID uuid.UUID) (int64, error)
	GetAveragePoints(ctx context.Context, leagueID uuid.UUID) (float64, error)
}