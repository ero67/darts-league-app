package entities

import (
	"time"
	"github.com/google/uuid"
)

type TournamentType string
type TournamentStatus string
type GameType string

const (
	TournamentTypeSingleElimination TournamentType = "single_elimination"
	TournamentTypeDoubleElimination TournamentType = "double_elimination"
	TournamentTypeRoundRobin        TournamentType = "round_robin"

	TournamentStatusSetup      TournamentStatus = "setup"
	TournamentStatusInProgress TournamentStatus = "in_progress"
	TournamentStatusCompleted  TournamentStatus = "completed"

	GameType501     GameType = "501"
	GameType301     GameType = "301"
	GameTypeCricket GameType = "cricket"
)

type Tournament struct {
	ID          uuid.UUID        `json:"id"`
	LeagueID    uuid.UUID        `json:"league_id"`
	Name        string           `json:"name"`
	Description *string          `json:"description,omitempty"`
	Type        TournamentType   `json:"type"`
	Status      TournamentStatus `json:"status"`

	// Game settings
	GameType        GameType `json:"game_type"`
	LegsPerMatch    int      `json:"legs_per_match"`
	SetsPerMatch    int      `json:"sets_per_match"`
	MaxPlayers      *int     `json:"max_players,omitempty"`
	TournamentNumber int     `json:"tournament_number"`

	// Financial
	EntryFee  *float64 `json:"entry_fee,omitempty"`
	PrizePool *float64 `json:"prize_pool,omitempty"`

	// Dates
	ScheduledDate *time.Time `json:"scheduled_date,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	StartedAt     *time.Time `json:"started_at,omitempty"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
}

// NewTournament creates a new tournament
func NewTournament(leagueID uuid.UUID, name, description string, tournamentType TournamentType, tournamentNumber int) (*Tournament, error) {
	if name == "" {
		return nil, ErrInvalidTournamentName
	}

	return &Tournament{
		ID:               uuid.New(),
		LeagueID:         leagueID,
		Name:             name,
		Description:      stringPtr(description),
		Type:             tournamentType,
		Status:           TournamentStatusSetup,
		GameType:         GameType501,
		LegsPerMatch:     3,
		SetsPerMatch:     1,
		TournamentNumber: tournamentNumber,
		CreatedAt:        time.Now(),
	}, nil
}

// StartTournament begins the tournament
func (t *Tournament) StartTournament() error {
	if t.Status != TournamentStatusSetup {
		return ErrTournamentAlreadyStarted
	}

	t.Status = TournamentStatusInProgress
	now := time.Now()
	t.StartedAt = &now

	return nil
}

// CompleteTournament finishes the tournament
func (t *Tournament) CompleteTournament() error {
	if t.Status == TournamentStatusCompleted {
		return ErrTournamentAlreadyCompleted
	}

	t.Status = TournamentStatusCompleted
	now := time.Now()
	t.CompletedAt = &now

	return nil
}

// CanAddPlayers returns true if players can still be added
func (t *Tournament) CanAddPlayers() bool {
	return t.Status == TournamentStatusSetup
}

// IsInProgress returns true if tournament is currently running
func (t *Tournament) IsInProgress() bool {
	return t.Status == TournamentStatusInProgress
}