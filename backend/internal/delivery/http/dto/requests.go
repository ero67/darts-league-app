package dto

import "github.com/google/uuid"

// Player DTOs
type CreatePlayerRequest struct {
	Name     string  `json:"name" binding:"required,min=1,max=100"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Nickname *string `json:"nickname,omitempty" binding:"omitempty,max=50"`
}

type UpdatePlayerRequest struct {
	Name      string  `json:"name" binding:"required,min=1,max=100"`
	Email     *string `json:"email,omitempty" binding:"omitempty,email"`
	Nickname  *string `json:"nickname,omitempty" binding:"omitempty,max=50"`
	AvatarURL *string `json:"avatar_url,omitempty" binding:"omitempty,url"`
}

// League DTOs
type CreateLeagueRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=255"`
	Description string `json:"description,omitempty"`
	Season      string `json:"season,omitempty" binding:"omitempty,max=100"`
}

type AddPlayerToLeagueRequest struct {
	PlayerID uuid.UUID `json:"player_id" binding:"required"`
}

// Tournament DTOs
type CreateTournamentRequest struct {
	LeagueID    uuid.UUID `json:"league_id" binding:"required"`
	Name        string    `json:"name" binding:"required,min=1,max=255"`
	Type        string    `json:"type" binding:"required,oneof=single_elimination double_elimination round_robin"`
}

type AddPlayerToTournamentRequest struct {
	PlayerID uuid.UUID `json:"player_id" binding:"required"`
}

// Match DTOs
type CreateMatchRequest struct {
	TournamentID *uuid.UUID `json:"tournament_id,omitempty"`
	Player1ID    *uuid.UUID `json:"player1_id,omitempty"`
	Player2ID    *uuid.UUID `json:"player2_id,omitempty"`
	Round        int        `json:"round,omitempty" binding:"min=0"`
	MatchNumber  int        `json:"match_number,omitempty" binding:"min=0"`
}

type StartMatchRequest struct {
	Player1ID uuid.UUID `json:"player1_id" binding:"required"`
	Player2ID uuid.UUID `json:"player2_id" binding:"required"`
}

type UpdateMatchScoreRequest struct {
	Player1Score int `json:"player1_score" binding:"min=0"`
	Player2Score int `json:"player2_score" binding:"min=0"`
}

type CompleteMatchRequest struct {
	WinnerID uuid.UUID `json:"winner_id" binding:"required"`
}

// Common DTOs
type PaginationQuery struct {
	Page  int `form:"page,default=1" binding:"min=1"`
	Limit int `form:"limit,default=10" binding:"min=1,max=100"`
}

func (p *PaginationQuery) GetOffset() int {
	return (p.Page - 1) * p.Limit
}