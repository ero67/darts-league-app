package postgres

import (
	"context"
	"darts-league-backend/internal/domain/repositories"

	"github.com/google/uuid"
)

type leagueStandingsRepository struct {
	db *DB
}

func NewLeagueStandingsRepository(db *DB) repositories.LeagueStandingsRepository {
	return &leagueStandingsRepository{db: db}
}

func (r *leagueStandingsRepository) Create(ctx context.Context, leagueID, playerID uuid.UUID) error {
	standing := &LeagueStanding{
		LeagueID:          leagueID,
		PlayerID:          playerID,
		TotalPoints:       0,
		TournamentsPlayed: 0,
		TournamentsWon:    0,
		FinalsReached:     0,
		SemiFinalsReached: 0,
		CurrentPosition:   0,
		PreviousPosition:  0,
	}
	return r.db.WithContext(ctx).Create(standing).Error
}

func (r *leagueStandingsRepository) GetLeagueStandings(ctx context.Context, leagueID uuid.UUID) ([]*repositories.LeagueStanding, error) {
	var models []LeagueStanding
	err := r.db.WithContext(ctx).
		Preload("Player").
		Where("league_id = ?", leagueID).
		Order("total_points DESC").
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	standings := make([]*repositories.LeagueStanding, len(models))
	for i, model := range models {
		standings[i] = ToLeagueStandingEntity(&model)
	}
	return standings, nil
}

// Add stubs for other required methods
func (r *leagueStandingsRepository) GetByLeagueAndPlayer(ctx context.Context, leagueID, playerID uuid.UUID) (*repositories.LeagueStanding, error) { return nil, nil }
func (r *leagueStandingsRepository) Update(ctx context.Context, standing *repositories.LeagueStanding) error { return nil }
func (r *leagueStandingsRepository) Delete(ctx context.Context, leagueID, playerID uuid.UUID) error { return nil }
func (r *leagueStandingsRepository) GetTopPlayers(ctx context.Context, leagueID uuid.UUID, limit int) ([]*repositories.LeagueStanding, error) { return nil, nil }
func (r *leagueStandingsRepository) GetPlayerPosition(ctx context.Context, leagueID, playerID uuid.UUID) (int, error) { return 0, nil }
func (r *leagueStandingsRepository) AddPoints(ctx context.Context, leagueID, playerID uuid.UUID, points int) error { return nil }
func (r *leagueStandingsRepository) UpdateTournamentStats(ctx context.Context, leagueID, playerID uuid.UUID, position int) error { return nil }
func (r *leagueStandingsRepository) RecalculatePositions(ctx context.Context, leagueID uuid.UUID) error { return nil }
func (r *leagueStandingsRepository) GetStandingsCount(ctx context.Context, leagueID uuid.UUID) (int64, error) { return 0, nil }
func (r *leagueStandingsRepository) GetAveragePoints(ctx context.Context, leagueID uuid.UUID) (float64, error) { return 0, nil }