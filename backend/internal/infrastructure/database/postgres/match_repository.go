package postgres

import (
	"context"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type matchRepository struct {
	db *DB
}

func NewMatchRepository(db *DB) repositories.MatchRepository {
	return &matchRepository{db: db}
}

func (r *matchRepository) Create(ctx context.Context, match *entities.Match) error {
	model := ToMatchModel(match)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *matchRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Match, error) {
	var model Match
	err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entities.ErrMatchNotFound
		}
		return nil, err
	}
	return ToMatchEntity(&model), nil
}

func (r *matchRepository) Update(ctx context.Context, match *entities.Match) error {
	model := ToMatchModel(match)
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *matchRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Match{}, "id = ?", id).Error
}

func (r *matchRepository) GetByTournamentID(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Match, error) {
	var models []Match
	err := r.db.WithContext(ctx).Where("tournament_id = ?", tournamentID).Find(&models).Error
	if err != nil {
		return nil, err
	}

	matches := make([]*entities.Match, len(models))
	for i, model := range models {
		matches[i] = ToMatchEntity(&model)
	}
	return matches, nil
}

func (r *matchRepository) GetByPlayerID(ctx context.Context, playerID uuid.UUID, limit, offset int) ([]*entities.Match, error) {
	var models []Match
	err := r.db.WithContext(ctx).
		Where("player1_id = ? OR player2_id = ?", playerID, playerID).
		Limit(limit).Offset(offset).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	matches := make([]*entities.Match, len(models))
	for i, model := range models {
		matches[i] = ToMatchEntity(&model)
	}
	return matches, nil
}

// Add stubs for other required methods
func (r *matchRepository) GetAll(ctx context.Context, limit, offset int) ([]*entities.Match, error) { return nil, nil }
func (r *matchRepository) GetByStatus(ctx context.Context, status entities.MatchStatus, limit, offset int) ([]*entities.Match, error) { return nil, nil }
func (r *matchRepository) GetTournamentMatches(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Match, error) { return r.GetByTournamentID(ctx, tournamentID) }
func (r *matchRepository) GetMatchesByRound(ctx context.Context, tournamentID uuid.UUID, round int) ([]*entities.Match, error) { return nil, nil }
func (r *matchRepository) GetCurrentMatches(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Match, error) { return nil, nil }
func (r *matchRepository) GetCompletedMatches(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Match, error) { return nil, nil }
func (r *matchRepository) GetPlayerMatches(ctx context.Context, playerID uuid.UUID, tournamentID *uuid.UUID) ([]*entities.Match, error) { return nil, nil }
func (r *matchRepository) GetPlayerMatchesInLeague(ctx context.Context, playerID, leagueID uuid.UUID) ([]*entities.Match, error) { return nil, nil }
func (r *matchRepository) GetLiveMatchesForPlayer(ctx context.Context, playerID uuid.UUID) ([]*entities.Match, error) { return nil, nil }
func (r *matchRepository) CreateBracketMatches(ctx context.Context, matches []*entities.Match) error { return nil }
func (r *matchRepository) GetNextMatch(ctx context.Context, tournamentID uuid.UUID, round int) (*entities.Match, error) { return nil, nil }
func (r *matchRepository) GetMaxRound(ctx context.Context, tournamentID uuid.UUID) (int, error) { return 0, nil }
func (r *matchRepository) GetMatchCount(ctx context.Context) (int64, error) { return 0, nil }
func (r *matchRepository) GetMatchCountByTournament(ctx context.Context, tournamentID uuid.UUID) (int64, error) { return 0, nil }
func (r *matchRepository) GetMatchCountByPlayer(ctx context.Context, playerID uuid.UUID) (int64, error) { return 0, nil }
func (r *matchRepository) GetMatchCountByStatus(ctx context.Context, status entities.MatchStatus) (int64, error) { return 0, nil }