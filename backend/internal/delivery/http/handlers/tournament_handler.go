package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"darts-league-backend/internal/delivery/http"
	"darts-league-backend/internal/delivery/http/dto"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/usecases"
)

type TournamentHandler struct {
	useCases *usecases.UseCases
}

func NewTournamentHandler(useCases *usecases.UseCases) *TournamentHandler {
	return &TournamentHandler{useCases: useCases}
}

// GetTournament godoc
// @Summary Get tournament by ID
// @Description Get a specific tournament by its ID
// @Tags tournaments
// @Accept json
// @Produce json
// @Param id path string true "Tournament ID"
// @Success 200 {object} http.Response
// @Router /api/tournaments/{id} [get]
func (h *TournamentHandler) GetTournament(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid tournament ID")
		return
	}

	tournament, err := h.useCases.Tournament.GetTournament(c.Request.Context(), id)
	if err != nil {
		if err == entities.ErrTournamentNotFound {
			http.NotFoundResponse(c, "Tournament not found")
			return
		}
		http.InternalErrorResponse(c, "Failed to get tournament")
		return
	}

	http.SuccessResponse(c, tournament)
}

// CreateTournament godoc
// @Summary Create a new tournament
// @Description Create a new tournament in a league
// @Tags tournaments
// @Accept json
// @Produce json
// @Param request body dto.CreateTournamentRequest true "Tournament data"
// @Success 201 {object} http.Response
// @Router /api/tournaments [post]
func (h *TournamentHandler) CreateTournament(c *gin.Context) {
	var req dto.CreateTournamentRequest
	log.Println("Creating tournament with request:", c.Request.Body)
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	tournamentType := entities.TournamentType(req.Type)
	log.Println(req.LeagueID, req.Name, tournamentType)
	tournament, err := h.useCases.Tournament.CreateTournament(c.Request.Context(), req.LeagueID, req.Name, tournamentType)
	if err != nil {
		http.InternalErrorResponse(c, "Failed to create tournament")
		return
	}

	http.CreatedResponse(c, tournament)
}

// GetLeagueTournaments godoc
// @Summary Get tournaments for a league
// @Description Get all tournaments in a specific league
// @Tags tournaments
// @Accept json
// @Produce json
// @Param league_id path string true "League ID"
// @Success 200 {object} http.Response
// @Router /api/leagues/{league_id}/tournaments [get]
func (h *TournamentHandler) GetLeagueTournaments(c *gin.Context) {
	idStr := c.Param("id")
	leagueID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid league ID")
		return
	}

	tournaments, err := h.useCases.Tournament.GetLeagueTournaments(c.Request.Context(), leagueID)
	if err != nil {
		http.InternalErrorResponse(c, "Failed to get tournaments")
		return
	}

	http.SuccessResponse(c, tournaments)
}

// AddPlayerToTournament godoc
// @Summary Add player to tournament
// @Description Add a player to a tournament
// @Tags tournaments
// @Accept json
// @Produce json
// @Param id path string true "Tournament ID"
// @Param request body dto.AddPlayerToTournamentRequest true "Player data"
// @Success 200 {object} http.Response
// @Router /api/tournaments/{id}/players [post]
func (h *TournamentHandler) AddPlayerToTournament(c *gin.Context) {
	idStr := c.Param("id")
	tournamentID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid tournament ID")
		return
	}

	var req dto.AddPlayerToTournamentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	err = h.useCases.Tournament.AddPlayerToTournament(c.Request.Context(), tournamentID, req.PlayerID)
	if err != nil {
		if err == entities.ErrTournamentAlreadyStarted {
			http.BadRequestResponse(c, "Cannot add players to a tournament that has already started")
			return
		}
		http.InternalErrorResponse(c, "Failed to add player to tournament")
		return
	}

	http.MessageResponse(c, "Player added to tournament successfully")
}

// StartTournament godoc
// @Summary Start a tournament
// @Description Start a tournament and generate the bracket
// @Tags tournaments
// @Accept json
// @Produce json
// @Param id path string true "Tournament ID"
// @Success 200 {object} http.Response
// @Router /api/tournaments/{id}/start [post]
func (h *TournamentHandler) StartTournament(c *gin.Context) {
	idStr := c.Param("id")
	tournamentID, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid tournament ID")
		return
	}

	tournament, err := h.useCases.Tournament.StartTournament(c.Request.Context(), tournamentID)
	if err != nil {
		if err == entities.ErrTournamentNotFound {
			http.NotFoundResponse(c, "Tournament not found")
			return
		}
		if err == entities.ErrTournamentAlreadyStarted {
			http.BadRequestResponse(c, "Tournament has already been started")
			return
		}
		http.InternalErrorResponse(c, "Failed to start tournament")
		return
	}

	http.SuccessResponse(c, tournament)
}