package postgres

import (
	"time"

	"github.com/google/uuid"
	// "gorm.io/gorm"
)

// Player GORM model
type Player struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string     `gorm:"size:100;not null"`
	Email     *string    `gorm:"size:255;uniqueIndex"`
	AvatarURL *string    `gorm:"size:500"`
	Nickname  *string    `gorm:"size:50"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

func (Player) TableName() string {
	return "players"
}

// League GORM model
type League struct {
	ID                 uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name               string     `gorm:"size:255;not null"`
	Description        *string    `gorm:"type:text"`
	Season             *string    `gorm:"size:100"`
	Status             string     `gorm:"size:50;default:'setup'"`
	PointsForWin       int        `gorm:"default:3"`
	PointsForRunnerUp  int        `gorm:"default:2"`
	PointsForSemiFinal int        `gorm:"default:1"`
	MaxPlayers         *int
	StartDate          *time.Time `gorm:"type:date"`
	EndDate            *time.Time `gorm:"type:date"`
	CreatedAt          time.Time  `gorm:"autoCreateTime"`
	UpdatedAt          time.Time  `gorm:"autoUpdateTime"`
}

func (League) TableName() string {
	return "leagues"
}

// Tournament GORM model
type Tournament struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	LeagueID         uuid.UUID  `gorm:"type:uuid;not null;index"`
	Name             string     `gorm:"size:255;not null"`
	Description      *string    `gorm:"type:text"`
	Type             string     `gorm:"size:50;not null"`
	Status           string     `gorm:"size:50;default:'setup'"`
	GameType         string     `gorm:"size:50;default:'501'"`
	LegsPerMatch     int        `gorm:"default:3"`
	SetsPerMatch     int        `gorm:"default:1"`
	MaxPlayers       *int
	EntryFee         *float64   `gorm:"type:decimal(10,2)"`
	PrizePool        *float64   `gorm:"type:decimal(10,2)"`
	TournamentNumber int        `gorm:"not null"`
	ScheduledDate    *time.Time `gorm:"type:date"`
	CreatedAt        time.Time  `gorm:"autoCreateTime"`
	StartedAt        *time.Time
	CompletedAt      *time.Time

	// Foreign key relationship
	League League `gorm:"foreignKey:LeagueID"`
}

func (Tournament) TableName() string {
	return "tournaments"
}

// Match GORM model
type Match struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TournamentID uuid.UUID  `gorm:"type:uuid;not null;index"`
	Round        int        `gorm:"not null"`
	MatchNumber  int        `gorm:"not null"`
	Player1ID    *uuid.UUID `gorm:"type:uuid;index"`
	Player2ID    *uuid.UUID `gorm:"type:uuid;index"`
	Player1Score int        `gorm:"default:0"`
	Player2Score int        `gorm:"default:0"`
	WinnerID     *uuid.UUID `gorm:"type:uuid;index"`
	Status       string     `gorm:"size:50;default:'pending'"`
	StartedAt    *time.Time
	CompletedAt  *time.Time
	CreatedAt    time.Time  `gorm:"autoCreateTime"`

	// Foreign key relationships
	Tournament Tournament `gorm:"foreignKey:TournamentID"`
	Player1    *Player    `gorm:"foreignKey:Player1ID"`
	Player2    *Player    `gorm:"foreignKey:Player2ID"`
	Winner     *Player    `gorm:"foreignKey:WinnerID"`
}

func (Match) TableName() string {
	return "matches"
}

// LeagueStanding GORM model
type LeagueStanding struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	LeagueID          uuid.UUID `gorm:"type:uuid;not null;index"`
	PlayerID          uuid.UUID `gorm:"type:uuid;not null;index"`
	TotalPoints       int       `gorm:"default:0"`
	TournamentsPlayed int       `gorm:"default:0"`
	TournamentsWon    int       `gorm:"default:0"`
	FinalsReached     int       `gorm:"default:0"`
	SemiFinalsReached int       `gorm:"default:0"`
	CurrentPosition   int
	PreviousPosition  int
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`

	// Foreign key relationships
	League League `gorm:"foreignKey:LeagueID"`
	Player Player `gorm:"foreignKey:PlayerID"`
}

func (LeagueStanding) TableName() string {
	return "league_standings"
}

// LeaguePlayer GORM model (junction table)
type LeaguePlayer struct {
	LeagueID uuid.UUID `gorm:"type:uuid;primaryKey"`
	PlayerID uuid.UUID `gorm:"type:uuid;primaryKey"`
	JoinedAt time.Time `gorm:"autoCreateTime"`
	IsActive bool      `gorm:"default:true"`

	// Foreign key relationships
	League League `gorm:"foreignKey:LeagueID"`
	Player Player `gorm:"foreignKey:PlayerID"`
}

func (LeaguePlayer) TableName() string {
	return "league_players"
}

// TournamentPlayer GORM model (junction table)
type TournamentPlayer struct {
	TournamentID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	PlayerID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Seed          *int
	FinalPosition *int
	PointsEarned  int       `gorm:"default:0"`
	JoinedAt      time.Time `gorm:"autoCreateTime"`

	// Foreign key relationships
	Tournament Tournament `gorm:"foreignKey:TournamentID"`
	Player     Player     `gorm:"foreignKey:PlayerID"`
}

func (TournamentPlayer) TableName() string {
	return "tournament_players"
}