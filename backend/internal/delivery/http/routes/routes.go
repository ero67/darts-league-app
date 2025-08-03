package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"darts-league-backend/internal/delivery/http/handlers"
	"darts-league-backend/internal/usecases"
)

func SetupRoutes(router *gin.Engine, useCases *usecases.UseCases) {
	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	// Initialize handlers
	playerHandler := handlers.NewPlayerHandler(useCases)
	leagueHandler := handlers.NewLeagueHandler(useCases)
	tournamentHandler := handlers.NewTournamentHandler(useCases)
	matchHandler := handlers.NewMatchHandler(useCases)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"database":  "connected",
			"service":   "darts-league-api",
		})
	})

	// API v1 routes
	api := router.Group("/api")
	{
		// Player routes - FIXED: use consistent parameter names
		players := api.Group("/players")
		{
			players.GET("", playerHandler.GetPlayers)
			players.POST("", playerHandler.CreatePlayer)
			players.GET("/search", playerHandler.SearchPlayers)
			players.GET("/:id", playerHandler.GetPlayer)
			players.PUT("/:id", playerHandler.UpdatePlayer)
			players.DELETE("/:id", playerHandler.DeletePlayer)
			players.GET("/:id/matches", matchHandler.GetPlayerMatches) // Use :id instead of :player_id
		}

		// League routes - FIXED: use consistent parameter names
		leagues := api.Group("/leagues")
		{
			leagues.GET("", leagueHandler.GetLeagues)
			leagues.POST("", leagueHandler.CreateLeague)
			leagues.GET("/:id", leagueHandler.GetLeague)
			leagues.POST("/:id/players", leagueHandler.AddPlayerToLeague)
			leagues.GET("/:id/standings", leagueHandler.GetLeagueStandings)
			leagues.POST("/:id/start", leagueHandler.StartLeague)
			leagues.GET("/:id/tournaments", tournamentHandler.GetLeagueTournaments) // Use :id instead of :league_id
		}

		// Tournament routes
		tournaments := api.Group("/tournaments")
		{
			tournaments.POST("", tournamentHandler.CreateTournament)
			tournaments.GET("/:id", tournamentHandler.GetTournament)
			tournaments.POST("/:id/players", tournamentHandler.AddPlayerToTournament)
			tournaments.POST("/:id/start", tournamentHandler.StartTournament)
			tournaments.GET("/:id/matches", matchHandler.GetTournamentMatches) // Use :id instead of :tournament_id
		}

		// Match routes
		matches := api.Group("/matches")
		{
			matches.POST("", matchHandler.CreateMatch)
			matches.GET("/:id", matchHandler.GetMatch)
			matches.POST("/:id/start", matchHandler.StartMatch)
			matches.PUT("/:id/score", matchHandler.UpdateMatchScore)
			matches.POST("/:id/complete", matchHandler.CompleteMatch)
		}
	}
}