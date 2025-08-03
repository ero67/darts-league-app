package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"darts-league-backend/internal/delivery/http"
	"darts-league-backend/internal/delivery/http/dto"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/usecases"
)

type MatchHandler struct {
	useCases *usecases.UseCases
}

func NewMatchHandler(useCases *usecases.UseCases) *MatchHandler {
	return &MatchHandler{useCases: useCases}
}

// CreateMatch godoc
// @Summary Create a new match
// @Description Create a new match between two players
// @Tags matches
// @Accept json
// @Produce json
// @Param request body dto.CreateMatchRequest true "Match creation data"
// @Success 201 {object} http.Response
// @Router /api/matches [post]
func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var req dto.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	// Create match with optional tournament assignment
	var match *entities.Match
	if req.TournamentID != nil {
		match = entities.NewMatch(*req.TournamentID, req.Round, req.MatchNumber)
	} else {
		// For standalone matches, use zero UUID and round/match number 1
		match = entities.NewMatch(uuid.Nil, 1, 1)
	}

	// Set players if provided
	if req.Player1ID != nil && req.Player2ID != nil {
		if err := match.SetPlayers(*req.Player1ID, *req.Player2ID); err != nil {
			http.BadRequestResponse(c, "Failed to set players")
			return
		}
	}

	// Save the match
	if err := h.useCases.Match.CreateMatch(c.Request.Context(), match); err != nil {
		http.InternalErrorResponse(c, "Failed to create match")
		return
	}

	http.CreatedResponse(c, match)
}

// GetMatch godoc
// @Summary Get match by ID
// @Description Get a specific match by its ID
// @Tags matches
// @Accept json
// @Produce json
// @Param id path string true "Match ID"
// @Success 200 {object} http.Response
// @Router /api/matches/{id} [get]
func (h *MatchHandler) GetMatch(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid match ID")
		return
	}

	match, err := h.useCases.Match.GetMatch(c.Request.Context(), id)
	if err != nil {
		if err == entities.ErrMatchNotFound {
			http.NotFoundResponse(c, "Match not found")
			return
		}
		http.InternalErrorResponse(c, "Failed to get match")
		return
	}

	http.SuccessResponse(c, match)
}

// GetTournamentMatches godoc
// @Summary Get matches for a tournament
// @Description Get all matches in a specific tournament
// @Tags matches
// @Accept json
// @Produce json
// @Param tournament_id path string true "Tournament ID"
// @Success 200 {object} http.Response
// @Router /api/tournaments/{tournament_id}/matches [get]
func (h *MatchHandler) GetTournamentMatches(c *gin.Context) {
	idStr := c.Param("tournament_id")
	tournamentID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid tournament ID")
		return
	}

	matches, err := h.useCases.Match.GetTournamentMatches(c.Request.Context(), tournamentID)
	if err != nil {
		http.InternalErrorResponse(c, "Failed to get matches")
		return
	}

	http.SuccessResponse(c, matches)
}

// StartMatch godoc
// @Summary Start a match
// @Description Start a match with two players
// @Tags matches
// @Accept json
// @Produce json
// @Param id path string true "Match ID"
// @Param request body dto.StartMatchRequest true "Match start data"
// @Success 200 {object} http.Response
// @Router /api/matches/{id}/start [post]
func (h *MatchHandler) StartMatch(c *gin.Context) {
	idStr := c.Param("id")
	matchID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid match ID")
		return
	}

	var req dto.StartMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	match, err := h.useCases.Match.StartMatch(c.Request.Context(), matchID, req.Player1ID, req.Player2ID)
	if err != nil {
		if err == entities.ErrMatchNotFound {
			http.NotFoundResponse(c, "Match not found")
			return
		}
		if err == entities.ErrMatchAlreadyStarted {
			http.BadRequestResponse(c, "Match has already been started")
			return
		}
		http.InternalErrorResponse(c, "Failed to start match")
		return
	}

	http.SuccessResponse(c, match)
}

// UpdateMatchScore godoc
// @Summary Update match score
// @Description Update the score of an ongoing match
// @Tags matches
// @Accept json
// @Produce json
// @Param id path string true "Match ID"
// @Param request body dto.UpdateMatchScoreRequest true "Score update data"
// @Success 200 {object} http.Response
// @Router /api/matches/{id}/score [put]
func (h *MatchHandler) UpdateMatchScore(c *gin.Context) {
	idStr := c.Param("id")
	matchID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid match ID")
		return
	}

	var req dto.UpdateMatchScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	match, err := h.useCases.Match.UpdateMatchScore(c.Request.Context(), matchID, req.Player1Score, req.Player2Score)
	if err != nil {
		if err == entities.ErrMatchNotFound {
			http.NotFoundResponse(c, "Match not found")
			return
		}
		if err == entities.ErrMatchNotInProgress {
			http.BadRequestResponse(c, "Match is not in progress")
			return
		}
		http.InternalErrorResponse(c, "Failed to update match score")
		return
	}

	http.SuccessResponse(c, match)
}

// CompleteMatch godoc
// @Summary Complete a match
// @Description Complete a match and declare the winner
// @Tags matches
// @Accept json
// @Produce json
// @Param id path string true "Match ID"
// @Param request body dto.CompleteMatchRequest true "Match completion data"
// @Success 200 {object} http.Response
// @Router /api/matches/{id}/complete [post]
func (h *MatchHandler) CompleteMatch(c *gin.Context) {
	idStr := c.Param("id")
	matchID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid match ID")
		return
	}

	var req dto.CompleteMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	match, err := h.useCases.Match.CompleteMatch(c.Request.Context(), matchID, req.WinnerID)
	if err != nil {
		if err == entities.ErrMatchNotFound {
			http.NotFoundResponse(c, "Match not found")
			return
		}
		if err == entities.ErrMatchNotInProgress {
			http.BadRequestResponse(c, "Match is not in progress")
			return
		}
		if err == entities.ErrInvalidWinner {
			http.BadRequestResponse(c, "Invalid winner - must be one of the match participants")
			return
		}
		http.InternalErrorResponse(c, "Failed to complete match")
		return
	}

	http.SuccessResponse(c, match)
}

// GetPlayerMatches godoc
// @Summary Get matches for a player
// @Description Get all matches for a specific player
// @Tags matches
// @Accept json
// @Produce json
// @Param player_id path string true "Player ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} http.PaginatedResponse
// @Router /api/players/{player_id}/matches [get]
func (h *MatchHandler) GetPlayerMatches(c *gin.Context) {
	idStr := c.Param("id")
	playerID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid player ID")
		return
	}

	var query dto.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		http.BadRequestResponse(c, "Invalid query parameters")
		return
	}

	matches, err := h.useCases.Match.GetPlayerMatches(c.Request.Context(), playerID, query.Limit, query.GetOffset())
	if err != nil {
		http.InternalErrorResponse(c, "Failed to get player matches")
		return
	}

	http.PaginatedSuccessResponse(c, matches, query.Page, query.Limit, 0)
}