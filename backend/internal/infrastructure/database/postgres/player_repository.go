package postgres

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/domain/repositories"
)

type playerRepository struct {
	db *DB
}

func NewPlayerRepository(db *DB) repositories.PlayerRepository {
	return &playerRepository{db: db}
}

func (r *playerRepository) Create(ctx context.Context, player *entities.Player) error {
	model := ToPlayerModel(player)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *playerRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Player, error) {
	var model Player
	err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entities.ErrPlayerNotFound
		}
		return nil, err
	}
	return ToPlayerEntity(&model), nil
}

func (r *playerRepository) GetByEmail(ctx context.Context, email string) (*entities.Player, error) {
	var model Player
	err := r.db.WithContext(ctx).First(&model, "email = ?", email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entities.ErrPlayerNotFound
		}
		return nil, err
	}
	return ToPlayerEntity(&model), nil
}

func (r *playerRepository) Update(ctx context.Context, player *entities.Player) error {
	model := ToPlayerModel(player)
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *playerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Player{}, "id = ?", id).Error
}

func (r *playerRepository) GetAll(ctx context.Context, limit, offset int) ([]*entities.Player, error) {
	var models []Player
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&models).Error
	if err != nil {
		return nil, err
	}

	players := make([]*entities.Player, len(models))
	for i, model := range models {
		players[i] = ToPlayerEntity(&model)
	}
	return players, nil
}

func (r *playerRepository) SearchByName(ctx context.Context, name string, limit int) ([]*entities.Player, error) {
	var models []Player
	err := r.db.WithContext(ctx).
		Where("name ILIKE ?", "%"+name+"%").
		Limit(limit).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	players := make([]*entities.Player, len(models))
	for i, model := range models {
		players[i] = ToPlayerEntity(&model)
	}
	return players, nil
}

func (r *playerRepository) GetByIDs(ctx context.Context, ids []uuid.UUID) ([]*entities.Player, error) {
	var models []Player
	err := r.db.WithContext(ctx).Find(&models, "id IN ?", ids).Error
	if err != nil {
		return nil, err
	}

	players := make([]*entities.Player, len(models))
	for i, model := range models {
		players[i] = ToPlayerEntity(&model)
	}
	return players, nil
}

func (r *playerRepository) GetLeaguePlayers(ctx context.Context, leagueID uuid.UUID) ([]*entities.Player, error) {
	var models []Player
	err := r.db.WithContext(ctx).
		Joins("JOIN league_players ON players.id = league_players.player_id").
		Where("league_players.league_id = ? AND league_players.is_active = ?", leagueID, true).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	players := make([]*entities.Player, len(models))
	for i, model := range models {
		players[i] = ToPlayerEntity(&model)
	}
	return players, nil
}

func (r *playerRepository) GetTournamentPlayers(ctx context.Context, tournamentID uuid.UUID) ([]*entities.Player, error) {
	var models []Player
	err := r.db.WithContext(ctx).
		Joins("JOIN tournament_players ON players.id = tournament_players.player_id").
		Where("tournament_players.tournament_id = ?", tournamentID).
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	players := make([]*entities.Player, len(models))
	for i, model := range models {
		players[i] = ToPlayerEntity(&model)
	}
	return players, nil
}

func (r *playerRepository) GetPlayerCount(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&Player{}).Count(&count).Error
	return count, err
}

func (r *playerRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&Player{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}