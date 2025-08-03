package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	
	"darts-league-backend/internal/delivery/http"
	"darts-league-backend/internal/delivery/http/dto"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/usecases"
)

type LeagueHandler struct {
	useCases *usecases.UseCases
}

func NewLeagueHandler(useCases *usecases.UseCases) *LeagueHandler {
	return &LeagueHandler{useCases: useCases}
}

// GetLeagues godoc
// @Summary Get all leagues
// @Description Get all leagues with pagination
// @Tags leagues
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} http.PaginatedResponse
// @Router /api/leagues [get]
func (h *LeagueHandler) GetLeagues(c *gin.Context) {
	var query dto.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		http.BadRequestResponse(c, "Invalid query parameters")
		return
	}

	leagues, err := h.useCases.League.GetAllLeagues(c.Request.Context(), query.Limit, query.GetOffset())
	if err != nil {
		http.InternalErrorResponse(c, "Failed to get leagues")
		return
	}

	http.PaginatedSuccessResponse(c, leagues, query.Page, query.Limit, 0)
}

// GetLeague godoc
// @Summary Get league by ID
// @Description Get a specific league by its ID
// @Tags leagues
// @Accept json
// @Produce json
// @Param id path string true "League ID"
// @Success 200 {object} http.Response
// @Router /api/leagues/{id} [get]
func (h *LeagueHandler) GetLeague(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid league ID")
		return
	}

	league, err := h.useCases.League.GetLeague(c.Request.Context(), id)
	if err != nil {
		if err == entities.ErrLeagueNotFound {
			http.NotFoundResponse(c, "League not found")
			return
		}
		http.InternalErrorResponse(c, "Failed to get league")
		return
	}

	http.SuccessResponse(c, league)
}

// CreateLeague godoc
// @Summary Create a new league
// @Description Create a new league
// @Tags leagues
// @Accept json
// @Produce json
// @Param request body dto.CreateLeagueRequest true "League data"
// @Success 201 {object} http.Response
// @Router /api/leagues [post]
func (h *LeagueHandler) CreateLeague(c *gin.Context) {
	var req dto.CreateLeagueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	league, err := h.useCases.League.CreateLeague(c.Request.Context(), req.Name, req.Description, req.Season)
	if err != nil {
		http.InternalErrorResponse(c, "Failed to create league")
		return
	}

	http.CreatedResponse(c, league)
}

// AddPlayerToLeague godoc
// @Summary Add player to league
// @Description Add a player to a league
// @Tags leagues
// @Accept json
// @Produce json
// @Param id path string true "League ID"
// @Param request body dto.AddPlayerToLeagueRequest true "Player data"
// @Success 200 {object} http.Response
// @Router /api/leagues/{id}/players [post]
func (h *LeagueHandler) AddPlayerToLeague(c *gin.Context) {
	idStr := c.Param("id")
	leagueID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid league ID")
		return
	}

	var req dto.AddPlayerToLeagueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	err = h.useCases.League.AddPlayerToLeague(c.Request.Context(), leagueID, req.PlayerID)
	if err != nil {
		http.InternalErrorResponse(c, "Failed to add player to league")
		return
	}

	http.MessageResponse(c, "Player added to league successfully")
}

// GetLeagueStandings godoc
// @Summary Get league standings
// @Description Get current standings for a league
// @Tags leagues
// @Accept json
// @Produce json
// @Param id path string true "League ID"
// @Success 200 {object} http.Response
// @Router /api/leagues/{id}/standings [get]
func (h *LeagueHandler) GetLeagueStandings(c *gin.Context) {
	idStr := c.Param("id")
	leagueID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid league ID")
		return
	}

	standings, err := h.useCases.League.GetLeagueStandings(c.Request.Context(), leagueID)
	if err != nil {
		http.InternalErrorResponse(c, "Failed to get league standings")
		return
	}

	http.SuccessResponse(c, standings)
}

// StartLeague godoc
// @Summary Start a league
// @Description Activate a league and start the season
// @Tags leagues
// @Accept json
// @Produce json
// @Param id path string true "League ID"
// @Success 200 {object} http.Response
// @Router /api/leagues/{id}/start [post]
func (h *LeagueHandler) StartLeague(c *gin.Context) {
	idStr := c.Param("id")
	leagueID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid league ID")
		return
	}

	league, err := h.useCases.League.StartLeague(c.Request.Context(), leagueID)
	if err != nil {
		if err == entities.ErrLeagueNotFound {
			http.NotFoundResponse(c, "League not found")
			return
		}
		if err == entities.ErrLeagueAlreadyStarted {
			http.BadRequestResponse(c, "League has already been started")
			return
		}
		http.InternalErrorResponse(c, "Failed to start league")
		return
	}

	http.SuccessResponse(c, league)
}