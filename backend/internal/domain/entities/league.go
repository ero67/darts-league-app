package entities

import (
	"time"
	"github.com/google/uuid"
)

type LeagueStatus string

const (
	LeagueStatusSetup     LeagueStatus = "setup"
	LeagueStatusActive    LeagueStatus = "active"
	LeagueStatusCompleted LeagueStatus = "completed"
)

type League struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description,omitempty"`
	Season      *string      `json:"season,omitempty"`
	Status      LeagueStatus `json:"status"`

	// Points system
	PointsForWin       int `json:"points_for_win"`
	PointsForRunnerUp  int `json:"points_for_runner_up"`
	PointsForSemiFinal int `json:"points_for_semi_final"`
	MaxPlayers         *int `json:"max_players,omitempty"`

	// Dates
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewLeague creates a new league with default point system
func NewLeague(name, description, season string) (*League, error) {
	if name == "" {
		return nil, ErrInvalidLeagueName
	}

	return &League{
		ID:                 uuid.New(),
		Name:               name,
		Description:        stringPtr(description),
		Season:             stringPtr(season),
		Status:             LeagueStatusSetup,
		PointsForWin:       3,
		PointsForRunnerUp:  2,
		PointsForSemiFinal: 1,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}, nil
}

// StartLeague transitions the league to active status
func (l *League) StartLeague() error {
	if l.Status != LeagueStatusSetup {
		return ErrLeagueAlreadyStarted
	}

	l.Status = LeagueStatusActive
	now := time.Now()
	l.StartDate = &now
	l.UpdatedAt = time.Now()

	return nil
}

// CompleteLeague finishes the league
func (l *League) CompleteLeague() error {
	if l.Status == LeagueStatusCompleted {
		return ErrLeagueAlreadyCompleted
	}

	l.Status = LeagueStatusCompleted
	now := time.Now()
	l.EndDate = &now
	l.UpdatedAt = time.Now()

	return nil
}

// IsActive returns true if the league is currently active
func (l *League) IsActive() bool {
	return l.Status == LeagueStatusActive
}

// CanAddTournaments returns true if tournaments can be added
func (l *League) CanAddTournaments() bool {
	return l.Status == LeagueStatusSetup || l.Status == LeagueStatusActive
}