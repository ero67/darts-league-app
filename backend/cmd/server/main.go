package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"darts-league-backend/internal/delivery/http/routes"
	"darts-league-backend/internal/infrastructure/database/postgres"
	"darts-league-backend/internal/usecases"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database connection
	dbConfig := postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := postgres.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(context.TODO()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("✅ Database connection established")

	// Initialize repositories
	factory := postgres.NewRepositoryFactory(db)
	defer factory.Close()

	playerRepo := factory.NewPlayerRepository()
	leagueRepo := factory.NewLeagueRepository()
	tournamentRepo := factory.NewTournamentRepository()
	matchRepo := factory.NewMatchRepository()
	standingsRepo := factory.NewLeagueStandingsRepository()

	// Initialize use cases
	useCases := usecases.NewUseCases(
		playerRepo,
		leagueRepo,
		tournamentRepo,
		matchRepo,
		standingsRepo,
	)

	log.Println("✅ Use cases initialized")

	// Initialize router
	router := gin.Default()

	// Setup all routes using your existing routes package
	routes.SetupRoutes(router, useCases)

	// Debug: Show registered routes
	log.Println("🔍 Registered routes:")
	for _, route := range router.Routes() {
		log.Printf("   %s %s", route.Method, route.Path)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("🚀 Server starting on port %s", port)
	// log.Printf("📖 API Documentation:")
	// log.Printf("   Health:     GET    /health")
	// log.Printf("   Players:    GET    /api/players")
	// log.Printf("   Players:    POST   /api/players") 
	// log.Printf("   Players:    GET    /api/players/{id}")
	// log.Printf("   Leagues:    GET    /api/leagues")
	// log.Printf("   Leagues:    POST   /api/leagues")
	// log.Printf("   And more...")
	
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}