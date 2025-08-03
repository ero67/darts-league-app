package postgres

import (
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/domain/repositories"
)

// ToPlayerEntity converts GORM Player model to domain entity
func ToPlayerEntity(model *Player) *entities.Player {
	return &entities.Player{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		AvatarURL: model.AvatarURL,
		Nickname:  model.Nickname,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

// ToPlayerModel converts domain entity to GORM Player model
func ToPlayerModel(entity *entities.Player) *Player {
	return &Player{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		AvatarURL: entity.AvatarURL,
		Nickname:  entity.Nickname,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

// ToLeagueEntity converts GORM League model to domain entity
func ToLeagueEntity(model *League) *entities.League {
	return &entities.League{
		ID:                 model.ID,
		Name:               model.Name,
		Description:        model.Description,
		Season:             model.Season,
		Status:             entities.LeagueStatus(model.Status),
		PointsForWin:       model.PointsForWin,
		PointsForRunnerUp:  model.PointsForRunnerUp,
		PointsForSemiFinal: model.PointsForSemiFinal,
		MaxPlayers:         model.MaxPlayers,
		StartDate:          model.StartDate,
		EndDate:            model.EndDate,
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
	}
}

// ToLeagueModel converts domain entity to GORM League model
func ToLeagueModel(entity *entities.League) *League {
	return &League{
		ID:                 entity.ID,
		Name:               entity.Name,
		Description:        entity.Description,
		Season:             entity.Season,
		Status:             string(entity.Status),
		PointsForWin:       entity.PointsForWin,
		PointsForRunnerUp:  entity.PointsForRunnerUp,
		PointsForSemiFinal: entity.PointsForSemiFinal,
		MaxPlayers:         entity.MaxPlayers,
		StartDate:          entity.StartDate,
		EndDate:            entity.EndDate,
		CreatedAt:          entity.CreatedAt,
		UpdatedAt:          entity.UpdatedAt,
	}
}

// ToTournamentEntity converts GORM Tournament model to domain entity
func ToTournamentEntity(model *Tournament) *entities.Tournament {
	return &entities.Tournament{
		ID:               model.ID,
		LeagueID:         model.LeagueID,
		Name:             model.Name,
		Description:      model.Description,
		Type:             entities.TournamentType(model.Type),
		Status:           entities.TournamentStatus(model.Status),
		GameType:         entities.GameType(model.GameType),
		LegsPerMatch:     model.LegsPerMatch,
		SetsPerMatch:     model.SetsPerMatch,
		MaxPlayers:       model.MaxPlayers,
		EntryFee:         model.EntryFee,
		PrizePool:        model.PrizePool,
		TournamentNumber: model.TournamentNumber,
		ScheduledDate:    model.ScheduledDate,
		CreatedAt:        model.CreatedAt,
		StartedAt:        model.StartedAt,
		CompletedAt:      model.CompletedAt,
	}
}

// ToTournamentModel converts domain entity to GORM Tournament model
func ToTournamentModel(entity *entities.Tournament) *Tournament {
	return &Tournament{
		ID:               entity.ID,
		LeagueID:         entity.LeagueID,
		Name:             entity.Name,
		Description:      entity.Description,
		Type:             string(entity.Type),
		Status:           string(entity.Status),
		GameType:         string(entity.GameType),
		LegsPerMatch:     entity.LegsPerMatch,
		SetsPerMatch:     entity.SetsPerMatch,
		MaxPlayers:       entity.MaxPlayers,
		EntryFee:         entity.EntryFee,
		PrizePool:        entity.PrizePool,
		TournamentNumber: entity.TournamentNumber,
		ScheduledDate:    entity.ScheduledDate,
		CreatedAt:        entity.CreatedAt,
		StartedAt:        entity.StartedAt,
		CompletedAt:      entity.CompletedAt,
	}
}

// ToMatchEntity converts GORM Match model to domain entity
func ToMatchEntity(model *Match) *entities.Match {
	return &entities.Match{
		ID:           model.ID,
		TournamentID: model.TournamentID,
		Round:        model.Round,
		MatchNumber:  model.MatchNumber,
		Player1ID:    model.Player1ID,
		Player2ID:    model.Player2ID,
		Player1Score: model.Player1Score,
		Player2Score: model.Player2Score,
		WinnerID:     model.WinnerID,
		Status:       entities.MatchStatus(model.Status),
		StartedAt:    model.StartedAt,
		CompletedAt:  model.CompletedAt,
		CreatedAt:    model.CreatedAt,
	}
}

// ToMatchModel converts domain entity to GORM Match model
func ToMatchModel(entity *entities.Match) *Match {
	return &Match{
		ID:           entity.ID,
		TournamentID: entity.TournamentID,
		Round:        entity.Round,
		MatchNumber:  entity.MatchNumber,
		Player1ID:    entity.Player1ID,
		Player2ID:    entity.Player2ID,
		Player1Score: entity.Player1Score,
		Player2Score: entity.Player2Score,
		WinnerID:     entity.WinnerID,
		Status:       string(entity.Status),
		StartedAt:    entity.StartedAt,
		CompletedAt:  entity.CompletedAt,
		CreatedAt:    entity.CreatedAt,
	}
}

// ToLeagueStandingEntity converts GORM model to repository struct
func ToLeagueStandingEntity(model *LeagueStanding) *repositories.LeagueStanding {
	var playerName, playerNickname string
	if model.Player.Name != "" {
		playerName = model.Player.Name
	}
	if model.Player.Nickname != nil {
		playerNickname = *model.Player.Nickname
	}

	positionChange := model.CurrentPosition - model.PreviousPosition

	return &repositories.LeagueStanding{
		ID:                model.ID,
		LeagueID:          model.LeagueID,
		PlayerID:          model.PlayerID,
		PlayerName:        playerName,
		PlayerNickname:    &playerNickname,
		TotalPoints:       model.TotalPoints,
		TournamentsPlayed: model.TournamentsPlayed,
		TournamentsWon:    model.TournamentsWon,
		FinalsReached:     model.FinalsReached,
		SemiFinalsReached: model.SemiFinalsReached,
		CurrentPosition:   model.CurrentPosition,
		PreviousPosition:  model.PreviousPosition,
		PositionChange:    positionChange,
	}
}