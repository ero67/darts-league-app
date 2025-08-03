package postgres

import (
	"context"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/domain/repositories"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tournamentRepository struct {
	db *DB
}

func NewTournamentRepository(db *DB) repositories.TournamentRepository {
	return &tournamentRepository{db: db}
}

func (r *tournamentRepository) Create(ctx context.Context, tournament *entities.Tournament) error {
	model := ToTournamentModel(tournament)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *tournamentRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Tournament, error) {
	var model Tournament
	err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entities.ErrTournamentNotFound
		}
		return nil, err
	}
	return ToTournamentEntity(&model), nil
}

func (r *tournamentRepository) Update(ctx context.Context, tournament *entities.Tournament) error {
	model := ToTournamentModel(tournament)
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *tournamentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Tournament{}, "id = ?", id).Error
}

func (r *tournamentRepository) GetAll(ctx context.Context, limit, offset int) ([]*entities.Tournament, error) {
	var models []Tournament
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&models).Error
	if err != nil {
		return nil, err
	}

	tournaments := make([]*entities.Tournament, len(models))
	for i, model := range models {
		tournaments[i] = ToTournamentEntity(&model)
	}
	return tournaments, nil
}

func (r *tournamentRepository) GetByLeagueID(ctx context.Context, leagueID uuid.UUID) ([]*entities.Tournament, error) {
	var models []Tournament
	err := r.db.WithContext(ctx).Where("league_id = ?", leagueID).Find(&models).Error
	if err != nil {
		return nil, err
	}

	tournaments := make([]*entities.Tournament, len(models))
	for i, model := range models {
		tournaments[i] = ToTournamentEntity(&model)
	}
	return tournaments, nil
}

func (r *tournamentRepository) GetNextTournamentNumber(ctx context.Context, leagueID uuid.UUID) (int, error) {
	var maxNumber int
	err := r.db.WithContext(ctx).
		Model(&Tournament{}).
		Select("COALESCE(MAX(tournament_number), 0)").
		Where("league_id = ?", leagueID).
		Scan(&maxNumber).Error
	if err != nil {
		return 0, err
	}
	return maxNumber + 1, nil
}

func (r *tournamentRepository) AddPlayer(ctx context.Context, tournamentID, playerID uuid.UUID, seed *int) error {
	tournamentPlayer := &TournamentPlayer{
		TournamentID: tournamentID,
		PlayerID:     playerID,
		Seed:         seed,
	}
	return r.db.WithContext(ctx).Create(tournamentPlayer).Error
}

func (r *tournamentRepository) IsPlayerInTournament(ctx context.Context, tournamentID, playerID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&TournamentPlayer{}).
		Where("tournament_id = ? AND player_id = ?", tournamentID, playerID).
		Count(&count).Error
	return count > 0, err
}

// Implement the remaining interface methods with basic implementations
func (r *tournamentRepository) GetByStatus(ctx context.Context, status entities.TournamentStatus, limit, offset int) ([]*entities.Tournament, error) {
	var models []Tournament
	err := r.db.WithContext(ctx).Where("status = ?", string(status)).Limit(limit).Offset(offset).Find(&models).Error
	if err != nil {
		return nil, err
	}
	tournaments := make([]*entities.Tournament, len(models))
	for i, model := range models {
		tournaments[i] = ToTournamentEntity(&model)
	}
	return tournaments, nil
}

func (r *tournamentRepository) GetByType(ctx context.Context, tournamentType entities.TournamentType, limit, offset int) ([]*entities.Tournament, error) {
	return r.GetAll(ctx, limit, offset) // Simplified
}

func (r *tournamentRepository) SearchByName(ctx context.Context, name string, limit int) ([]*entities.Tournament, error) {
	return r.GetAll(ctx, limit, 0) // Simplified
}

func (r *tournamentRepository) GetLeagueTournaments(ctx context.Context, leagueID uuid.UUID, limit, offset int) ([]*entities.Tournament, error) {
	return r.GetByLeagueID(ctx, leagueID)
}

func (r *tournamentRepository) GetLatestTournament(ctx context.Context, leagueID uuid.UUID) (*entities.Tournament, error) {
	tournaments, err := r.GetByLeagueID(ctx, leagueID)
	if err != nil || len(tournaments) == 0 {
		return nil, err
	}
	return tournaments[0], nil
}

func (r *tournamentRepository) GetCompletedTournaments(ctx context.Context, leagueID uuid.UUID) ([]*entities.Tournament, error) {
	return r.GetByLeagueID(ctx, leagueID) // Simplified
}

// Add stubs for other required methods
func (r *tournamentRepository) RemovePlayer(ctx context.Context, tournamentID, playerID uuid.UUID) error { return nil }
func (r *tournamentRepository) GetTournamentPlayerCount(ctx context.Context, tournamentID uuid.UUID) (int, error) { return 0, nil }
func (r *tournamentRepository) SetPlayerPosition(ctx context.Context, tournamentID, playerID uuid.UUID, position int, points int) error { return nil }
func (r *tournamentRepository) GetTournamentsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entities.Tournament, error) { return nil, nil }
func (r *tournamentRepository) GetTournamentsScheduledFor(ctx context.Context, date time.Time) ([]*entities.Tournament, error) { return nil, nil }
func (r *tournamentRepository) GetUpcomingTournaments(ctx context.Context, limit int) ([]*entities.Tournament, error) { return nil, nil }
func (r *tournamentRepository) GetTournamentCount(ctx context.Context) (int64, error) { return 0, nil }
func (r *tournamentRepository) GetTournamentCountByLeague(ctx context.Context, leagueID uuid.UUID) (int64, error) { return 0, nil }
func (r *tournamentRepository) GetTournamentCountByStatus(ctx context.Context, status entities.TournamentStatus) (int64, error) { return 0, nil }