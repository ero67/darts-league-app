package repositories

// RepositoryFactory creates repository instances
type RepositoryFactory interface {
	// Individual repositories
	NewPlayerRepository() PlayerRepository
	NewLeagueRepository() LeagueRepository
	NewTournamentRepository() TournamentRepository
	NewMatchRepository() MatchRepository
	NewLeagueStandingsRepository() LeagueStandingsRepository
	NewStatisticsRepository() StatisticsRepository

	// Unit of work for transactions
	NewUnitOfWork() (UnitOfWork, error)

	// Factory lifecycle
	Close() error
}