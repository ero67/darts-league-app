package entities

import (
	"time"
	"github.com/google/uuid"
)

type MatchStatus string

const (
	MatchStatusPending    MatchStatus = "pending"
	MatchStatusInProgress MatchStatus = "in_progress"
	MatchStatusCompleted  MatchStatus = "completed"
)

type Match struct {
	ID           uuid.UUID   `json:"id"`
	TournamentID uuid.UUID   `json:"tournament_id"`
	Round        int         `json:"round"`
	MatchNumber  int         `json:"match_number"`
	Player1ID    *uuid.UUID  `json:"player1_id,omitempty"`
	Player2ID    *uuid.UUID  `json:"player2_id,omitempty"`
	Player1Score int         `json:"player1_score"`
	Player2Score int         `json:"player2_score"`
	WinnerID     *uuid.UUID  `json:"winner_id,omitempty"`
	Status       MatchStatus `json:"status"`
	StartedAt    *time.Time  `json:"started_at,omitempty"`
	CompletedAt  *time.Time  `json:"completed_at,omitempty"`
	CreatedAt    time.Time   `json:"created_at"`
}

// NewMatch creates a new match
func NewMatch(tournamentID uuid.UUID, round, matchNumber int) *Match {
	return &Match{
		ID:           uuid.New(),
		TournamentID: tournamentID,
		Round:        round,
		MatchNumber:  matchNumber,
		Player1Score: 0,
		Player2Score: 0,
		Status:       MatchStatusPending,
		CreatedAt:    time.Now(),
	}
}

// SetPlayers assigns players to the match
func (m *Match) SetPlayers(player1ID, player2ID uuid.UUID) error {
	if m.Status != MatchStatusPending {
		return ErrMatchAlreadyStarted
	}

	m.Player1ID = &player1ID
	m.Player2ID = &player2ID
	return nil
}

// StartMatch begins the match
func (m *Match) StartMatch() error {
	if m.Player1ID == nil || m.Player2ID == nil {
		return ErrMatchMissingPlayers
	}
	if m.Status != MatchStatusPending {
		return ErrMatchAlreadyStarted
	}

	m.Status = MatchStatusInProgress
	now := time.Now()
	m.StartedAt = &now

	return nil
}

// UpdateScore updates the match score
func (m *Match) UpdateScore(player1Score, player2Score int) error {
	if m.Status != MatchStatusInProgress {
		return ErrMatchNotInProgress
	}

	m.Player1Score = player1Score
	m.Player2Score = player2Score

	return nil
}

// CompleteMatch finishes the match and determines winner
func (m *Match) CompleteMatch(winnerID uuid.UUID) error {
	if m.Status != MatchStatusInProgress {
		return ErrMatchNotInProgress
	}

	if m.Player1ID == nil || m.Player2ID == nil {
		return ErrMatchMissingPlayers
	}

	if winnerID != *m.Player1ID && winnerID != *m.Player2ID {
		return ErrInvalidWinner
	}

	m.WinnerID = &winnerID
	m.Status = MatchStatusCompleted
	now := time.Now()
	m.CompletedAt = &now

	return nil
}

// GetOpponent returns the opponent of the given player
func (m *Match) GetOpponent(playerID uuid.UUID) (*uuid.UUID, error) {
	if m.Player1ID != nil && *m.Player1ID == playerID {
		return m.Player2ID, nil
	}
	if m.Player2ID != nil && *m.Player2ID == playerID {
		return m.Player1ID, nil
	}
	return nil, ErrPlayerNotInMatch
}