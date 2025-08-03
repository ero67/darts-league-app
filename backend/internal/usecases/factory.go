package usecases

import "darts-league-backend/internal/domain/repositories"

// UseCases holds all use case instances
type UseCases struct {
	Player     *PlayerUseCase
	League     *LeagueUseCase
	Tournament *TournamentUseCase
	Match      *MatchUseCase
}

// NewUseCases creates all use case instances
func NewUseCases(
	playerRepo repositories.PlayerRepository,
	leagueRepo repositories.LeagueRepository,
	tournamentRepo repositories.TournamentRepository,
	matchRepo repositories.MatchRepository,
	standingsRepo repositories.LeagueStandingsRepository,
) *UseCases {
	return &UseCases{
		Player:     NewPlayerUseCase(playerRepo),
		League:     NewLeagueUseCase(leagueRepo, standingsRepo),
		Tournament: NewTournamentUseCase(tournamentRepo, leagueRepo, matchRepo),
		Match:      NewMatchUseCase(matchRepo, standingsRepo),
	}
}