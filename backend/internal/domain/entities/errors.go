package entities

import "errors"

// Player errors
var (
	ErrInvalidPlayerName = errors.New("player name cannot be empty")
	ErrPlayerNotFound    = errors.New("player not found")
)

// League errors
var (
	ErrInvalidLeagueName       = errors.New("league name cannot be empty")
	ErrLeagueNotFound          = errors.New("league not found")
	ErrLeagueAlreadyStarted    = errors.New("league has already started")
	ErrLeagueAlreadyCompleted  = errors.New("league is already completed")
)

// Tournament errors
var (
	ErrInvalidTournamentName      = errors.New("tournament name cannot be empty")
	ErrTournamentNotFound         = errors.New("tournament not found")
	ErrTournamentAlreadyStarted   = errors.New("tournament has already started")
	ErrTournamentAlreadyCompleted = errors.New("tournament is already completed")
)

// Match errors
var (
	ErrMatchNotFound        = errors.New("match not found")
	ErrMatchAlreadyStarted  = errors.New("match has already started")
	ErrMatchNotInProgress   = errors.New("match is not in progress")
	ErrMatchMissingPlayers  = errors.New("match requires both players to be set")
	ErrPlayerNotInMatch     = errors.New("player is not participating in this match")
	ErrInvalidWinner        = errors.New("winner must be one of the match participants")
)