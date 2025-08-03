package postgres

import (
	"context"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/domain/repositories"
)

type leagueRepository struct {
	db *DB
}

func NewLeagueRepository(db *DB) repositories.LeagueRepository {
	return &leagueRepository{db: db}
}

func (r *leagueRepository) Create(ctx context.Context, league *entities.League) error {
	model := ToLeagueModel(league)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *leagueRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.League, error) {
	var model League
	err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entities.ErrLeagueNotFound
		}
		return nil, err
	}
	return ToLeagueEntity(&model), nil
}

func (r *leagueRepository) Update(ctx context.Context, league *entities.League) error {
	model := ToLeagueModel(league)
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *leagueRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&League{}, "id = ?", id).Error
}

func (r *leagueRepository) GetAll(ctx context.Context, limit, offset int) ([]*entities.League, error) {
	var models []League
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&models).Error
	if err != nil {
		return nil, err
	}

	leagues := make([]*entities.League, len(models))
	for i, model := range models {
		leagues[i] = ToLeagueEntity(&model)
	}
	return leagues, nil
}

func (r *leagueRepository) GetByStatus(ctx context.Context, status entities.LeagueStatus, limit, offset int) ([]*entities.League, error) {
	var models []League
	err := r.db.WithContext(ctx).
		Where("status = ?", string(status)).
		Limit(limit).Offset(offset).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	leagues := make([]*entities.League, len(models))
	for i, model := range models {
		leagues[i] = ToLeagueEntity(&model)
	}
	return leagues, nil
}

func (r *leagueRepository) GetBySeason(ctx context.Context, season string) ([]*entities.League, error) {
	var models []League
	err := r.db.WithContext(ctx).Where("season = ?", season).Find(&models).Error
	if err != nil {
		return nil, err
	}

	leagues := make([]*entities.League, len(models))
	for i, model := range models {
		leagues[i] = ToLeagueEntity(&model)
	}
	return leagues, nil
}

func (r *leagueRepository) GetActive(ctx context.Context) ([]*entities.League, error) {
	return r.GetByStatus(ctx, entities.LeagueStatusActive, 100, 0)
}

func (r *leagueRepository) SearchByName(ctx context.Context, name string, limit int) ([]*entities.League, error) {
	var models []League
	err := r.db.WithContext(ctx).
		Where("name ILIKE ?", "%"+name+"%").
		Limit(limit).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	leagues := make([]*entities.League, len(models))
	for i, model := range models {
		leagues[i] = ToLeagueEntity(&model)
	}
	return leagues, nil
}

func (r *leagueRepository) AddPlayer(ctx context.Context, leagueID, playerID uuid.UUID) error {
	leaguePlayer := &LeaguePlayer{
		LeagueID: leagueID,
		PlayerID: playerID,
		IsActive: true,
	}
	return r.db.WithContext(ctx).Create(leaguePlayer).Error
}

func (r *leagueRepository) RemovePlayer(ctx context.Context, leagueID, playerID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&LeaguePlayer{}).
		Where("league_id = ? AND player_id = ?", leagueID, playerID).
		Update("is_active", false).Error
}

func (r *leagueRepository) IsPlayerInLeague(ctx context.Context, leagueID, playerID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&LeaguePlayer{}).
		Where("league_id = ? AND player_id = ? AND is_active = ?", leagueID, playerID, true).
		Count(&count).Error
	return count > 0, err
}

func (r *leagueRepository) GetLeaguePlayerCount(ctx context.Context, leagueID uuid.UUID) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&LeaguePlayer{}).
		Where("league_id = ? AND is_active = ?", leagueID, true).
		Count(&count).Error
	return int(count), err
}

func (r *leagueRepository) GetLeaguesInDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entities.League, error) {
	var models []League
	err := r.db.WithContext(ctx).
		Where("start_date >= ? AND end_date <= ?", startDate, endDate).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	leagues := make([]*entities.League, len(models))
	for i, model := range models {
		leagues[i] = ToLeagueEntity(&model)
	}
	return leagues, nil
}

func (r *leagueRepository) GetCurrentLeagues(ctx context.Context) ([]*entities.League, error) {
	now := time.Now()
	var models []League
	err := r.db.WithContext(ctx).
		Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", now, now).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	leagues := make([]*entities.League, len(models))
	for i, model := range models {
		leagues[i] = ToLeagueEntity(&model)
	}
	return leagues, nil
}

func (r *leagueRepository) GetLeagueCount(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&League{}).Count(&count).Error
	return count, err
}

func (r *leagueRepository) GetLeagueCountByStatus(ctx context.Context, status entities.LeagueStatus) (int64, error) {
 var count int64
 err := r.db.WithContext(ctx).
  Model(&League{}).
  Where("status = ?", string(status)).
  Count(&count).Error
 return count, err
}