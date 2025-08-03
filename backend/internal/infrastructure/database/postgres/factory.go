package postgres

import (
	"darts-league-backend/internal/domain/repositories"
)

type repositoryFactory struct {
	db *DB
}

func NewRepositoryFactory(db *DB) repositories.RepositoryFactory {
	return &repositoryFactory{db: db}
}

func (f *repositoryFactory) NewPlayerRepository() repositories.PlayerRepository {
	return NewPlayerRepository(f.db)
}

func (f *repositoryFactory) NewLeagueRepository() repositories.LeagueRepository {
	return NewLeagueRepository(f.db)
}

func (f *repositoryFactory) NewTournamentRepository() repositories.TournamentRepository {
	return NewTournamentRepository(f.db)
}

func (f *repositoryFactory) NewMatchRepository() repositories.MatchRepository {
	return NewMatchRepository(f.db)
}

func (f *repositoryFactory) NewLeagueStandingsRepository() repositories.LeagueStandingsRepository {
	return NewLeagueStandingsRepository(f.db)
}

func (f *repositoryFactory) NewStatisticsRepository() repositories.StatisticsRepository {
	// TODO: Implement statistics repository when needed
	return nil
}

func (f *repositoryFactory) NewUnitOfWork() (repositories.UnitOfWork, error) {
	// TODO: Implement unit of work when needed for complex transactions
	return nil, nil
}

func (f *repositoryFactory) Close() error {
	return f.db.Close()
}