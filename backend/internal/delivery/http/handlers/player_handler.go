package handlers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	
	"darts-league-backend/internal/delivery/http"
	"darts-league-backend/internal/delivery/http/dto"
	"darts-league-backend/internal/domain/entities"
	"darts-league-backend/internal/usecases"
)

type PlayerHandler struct {
	useCases *usecases.UseCases
}

func NewPlayerHandler(useCases *usecases.UseCases) *PlayerHandler {
	return &PlayerHandler{useCases: useCases}
}

// GetPlayers godoc
// @Summary Get all players
// @Description Get all players with pagination
// @Tags players
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} http.PaginatedResponse
// @Router /api/players [get]
func (h *PlayerHandler) GetPlayers(c *gin.Context) {
	var query dto.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		http.BadRequestResponse(c, "Invalid query parameters")
		return
	}

	players, err := h.useCases.Player.GetAllPlayers(c.Request.Context(), query.Limit, query.GetOffset())
	if err != nil {
		http.InternalErrorResponse(c, "Failed to get players")
		return
	}

	// For now, we'll skip total count for simplicity
	http.PaginatedSuccessResponse(c, players, query.Page, query.Limit, 0)
}

// GetPlayer godoc
// @Summary Get player by ID
// @Description Get a specific player by their ID
// @Tags players
// @Accept json
// @Produce json
// @Param id path string true "Player ID"
// @Success 200 {object} http.Response
// @Router /api/players/{id} [get]
func (h *PlayerHandler) GetPlayer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid player ID")
		return
	}

	player, err := h.useCases.Player.GetPlayer(c.Request.Context(), id)
	if err != nil {
		if err == entities.ErrPlayerNotFound {
			http.NotFoundResponse(c, "Player not found")
			return
		}
		http.InternalErrorResponse(c, "Failed to get player")
		return
	}

	http.SuccessResponse(c, player)
}

// CreatePlayer godoc
// @Summary Create a new player
// @Description Create a new player
// @Tags players
// @Accept json
// @Produce json
// @Param request body dto.CreatePlayerRequest true "Player data"
// @Success 201 {object} http.Response
// @Router /api/players [post]
func (h *PlayerHandler) CreatePlayer(c *gin.Context) {
	var req dto.CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	player, err := h.useCases.Player.CreatePlayer(c.Request.Context(), req.Name, req.Email, req.Nickname)
	if err != nil {
		http.InternalErrorResponse(c, "Failed to create player")
		return
	}

	http.CreatedResponse(c, player)
}

// UpdatePlayer godoc
// @Summary Update a player
// @Description Update an existing player
// @Tags players
// @Accept json
// @Produce json
// @Param id path string true "Player ID"
// @Param request body dto.UpdatePlayerRequest true "Updated player data"
// @Success 200 {object} http.Response
// @Router /api/players/{id} [put]
func (h *PlayerHandler) UpdatePlayer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid player ID")
		return
	}

	var req dto.UpdatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http.BadRequestResponse(c, "Invalid request data")
		return
	}

	player, err := h.useCases.Player.UpdatePlayer(c.Request.Context(), id, req.Name, req.Email, req.Nickname, req.AvatarURL)
	if err != nil {
		if err == entities.ErrPlayerNotFound {
			http.NotFoundResponse(c, "Player not found")
			return
		}
		http.InternalErrorResponse(c, "Failed to update player")
		return
	}

	http.SuccessResponse(c, player)
}

// DeletePlayer godoc
// @Summary Delete a player
// @Description Delete an existing player
// @Tags players
// @Accept json
// @Produce json
// @Param id path string true "Player ID"
// @Success 200 {object} http.Response
// @Router /api/players/{id} [delete]
func (h *PlayerHandler) DeletePlayer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.BadRequestResponse(c, "Invalid player ID")
		return
	}

	err = h.useCases.Player.DeletePlayer(c.Request.Context(), id)
	if err != nil {
		if err == entities.ErrPlayerNotFound {
			http.NotFoundResponse(c, "Player not found")
			return
		}
		http.InternalErrorResponse(c, "Failed to delete player")
		return
	}

	http.MessageResponse(c, "Player deleted successfully")
}

// SearchPlayers godoc
// @Summary Search players by name
// @Description Search for players by name
// @Tags players
// @Accept json
// @Produce json
// @Param q query string true "Search query"
// @Param limit query int false "Items limit" default(10)
// @Success 200 {object} http.Response
// @Router /api/players/search [get]
func (h *PlayerHandler) SearchPlayers(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		http.BadRequestResponse(c, "Search query is required")
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	// For now, we'll use the GetAllPlayers method
	// In a real implementation, you'd implement search in the repository
	players, err := h.useCases.Player.GetAllPlayers(c.Request.Context(), limit, 0)
	if err != nil {
		http.InternalErrorResponse(c, "Failed to search players")
		return
	}

	http.SuccessResponse(c, players)
}