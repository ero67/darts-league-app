package repositories

import (
	"context"
	"time"
	"github.com/google/uuid"
)

// TournamentStats represents player statistics for a specific tournament
type TournamentStats struct {
	ID                    uuid.UUID `json:"id"`
	TournamentID          uuid.UUID `json:"tournament_id"`
	PlayerID              uuid.UUID `json:"player_id"`
	FinalPosition         *int      `json:"final_position,omitempty"`
	MatchesPlayed         int       `json:"matches_played"`
	MatchesWon            int       `json:"matches_won"`
	LegsPlayed            int       `json:"legs_played"`
	LegsWon               int       `json:"legs_won"`
	TotalThrows           int       `json:"total_throws"`
	TotalScore            int       `json:"total_score"`
	AverageScore          float64   `json:"average_score"`
	BestFinish            *int      `json:"best_finish,omitempty"`
	CheckoutPercentage    float64   `json:"checkout_percentage"`
	First9Average         float64   `json:"first_9_average"`
	SinglesHit            int       `json:"singles_hit"`
	DoublesHit            int       `json:"doubles_hit"`
	TriplesHit            int       `json:"triples_hit"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// LeagueStats represents aggregated player statistics across all tournaments in a league
type LeagueStats struct {
	ID                        uuid.UUID `json:"id"`
	LeagueID                  uuid.UUID `json:"league_id"`
	PlayerID                  uuid.UUID `json:"player_id"`
	TotalMatchesPlayed        int       `json:"total_matches_played"`
	TotalMatchesWon           int       `json:"total_matches_won"`
	TotalLegsPlayed           int       `json:"total_legs_played"`
	TotalLegsWon              int       `json:"total_legs_won"`
	TotalThrows               int       `json:"total_throws"`
	TotalScore                int       `json:"total_score"`
	OverallAverage            float64   `json:"overall_average"`
	BestFinish                *int      `json:"best_finish,omitempty"`
	OverallCheckoutPercentage float64   `json:"overall_checkout_percentage"`
	TournamentWins            int       `json:"tournament_wins"`
	PodiumFinishes            int       `json:"podium_finishes"`
	UpdatedAt                 time.Time `json:"updated_at"`
}

type StatisticsRepository interface {
	// Tournament Statistics
	CreateTournamentStats(ctx context.Context, stats *TournamentStats) error
	GetTournamentStats(ctx context.Context, tournamentID, playerID uuid.UUID) (*TournamentStats, error)
	UpdateTournamentStats(ctx context.Context, stats *TournamentStats) error
	GetAllTournamentStats(ctx context.Context, tournamentID uuid.UUID) ([]*TournamentStats, error)
	GetPlayerTournamentStats(ctx context.Context, playerID uuid.UUID, limit, offset int) ([]*TournamentStats, error)

	// League Statistics
	CreateLeagueStats(ctx context.Context, stats *LeagueStats) error
	GetLeagueStats(ctx context.Context, leagueID, playerID uuid.UUID) (*LeagueStats, error)
	UpdateLeagueStats(ctx context.Context, stats *LeagueStats) error
	GetAllLeagueStats(ctx context.Context, leagueID uuid.UUID) ([]*LeagueStats, error)
	GetPlayerLeagueStats(ctx context.Context, playerID uuid.UUID) ([]*LeagueStats, error)

	// Aggregation and calculations
	RecalculateLeagueStats(ctx context.Context, leagueID, playerID uuid.UUID) error
	GetTopPerformers(ctx context.Context, leagueID uuid.UUID, metric string, limit int) ([]*LeagueStats, error)
	GetLeagueAverages(ctx context.Context, leagueID uuid.UUID) (*LeagueStats, error)

	// Comparative statistics
	ComparePlayerStats(ctx context.Context, player1ID, player2ID, leagueID uuid.UUID) (*LeagueStats, *LeagueStats, error)
	GetPlayerRanking(ctx context.Context, leagueID, playerID uuid.UUID, metric string) (int, error)
}