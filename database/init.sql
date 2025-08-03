-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Players table
CREATE TABLE players (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE,
    avatar_url VARCHAR(500),
    nickname VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Leagues table (main competition like "Winter 2025 League")
CREATE TABLE leagues (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    season VARCHAR(100), -- e.g., "2025 Winter", "2025 Spring"
    status VARCHAR(50) DEFAULT 'setup', -- 'setup', 'active', 'completed'
    
    -- League settings
    points_for_win INTEGER DEFAULT 3,
    points_for_runner_up INTEGER DEFAULT 2,
    points_for_semi_final INTEGER DEFAULT 1,
    max_players INTEGER,
    
    -- Dates
    start_date DATE,
    end_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- League participants (players registered for the entire league)
CREATE TABLE league_players (
    league_id UUID REFERENCES leagues(id) ON DELETE CASCADE,
    player_id UUID REFERENCES players(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    PRIMARY KEY (league_id, player_id)
);

-- League standings (overall points table)
CREATE TABLE league_standings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    league_id UUID REFERENCES leagues(id) ON DELETE CASCADE,
    player_id UUID REFERENCES players(id) ON DELETE CASCADE,
    
    -- Points and performance
    total_points INTEGER DEFAULT 0,
    tournaments_played INTEGER DEFAULT 0,
    tournaments_won INTEGER DEFAULT 0,
    finals_reached INTEGER DEFAULT 0,
    semi_finals_reached INTEGER DEFAULT 0,
    
    -- Position tracking
    current_position INTEGER,
    previous_position INTEGER,
    
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(league_id, player_id)
);

-- Tournaments table (individual events within a league)
CREATE TABLE tournaments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    league_id UUID REFERENCES leagues(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    
    -- Tournament format
    type VARCHAR(50) NOT NULL, -- 'single_elimination', 'double_elimination', 'round_robin'
    status VARCHAR(50) DEFAULT 'setup', -- 'setup', 'in_progress', 'completed'
    
    -- Game settings
    game_type VARCHAR(50) DEFAULT '501', -- '501', '301', 'cricket'
    legs_per_match INTEGER DEFAULT 3,
    sets_per_match INTEGER DEFAULT 1,
    
    -- Tournament details
    max_players INTEGER,
    entry_fee DECIMAL(10,2),
    prize_pool DECIMAL(10,2),
    tournament_number INTEGER, -- 1st tournament, 2nd tournament, etc.
    
    -- Dates
    scheduled_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMP,
    completed_at TIMESTAMP
);

-- Tournament participants (subset of league players)
CREATE TABLE tournament_players (
    tournament_id UUID REFERENCES tournaments(id) ON DELETE CASCADE,
    player_id UUID REFERENCES players(id) ON DELETE CASCADE,
    seed INTEGER,
    final_position INTEGER, -- 1st, 2nd, 3rd, etc.
    points_earned INTEGER DEFAULT 0,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (tournament_id, player_id)
);

-- Matches table
CREATE TABLE matches (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tournament_id UUID REFERENCES tournaments(id) ON DELETE CASCADE,
    round INTEGER NOT NULL,
    match_number INTEGER NOT NULL,
    player1_id UUID REFERENCES players(id),
    player2_id UUID REFERENCES players(id),
    player1_score INTEGER DEFAULT 0, -- sets/legs won
    player2_score INTEGER DEFAULT 0,
    winner_id UUID REFERENCES players(id),
    status VARCHAR(50) DEFAULT 'pending', -- 'pending', 'in_progress', 'completed'
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Games table (individual legs within a match)
CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    match_id UUID REFERENCES matches(id) ON DELETE CASCADE,
    leg_number INTEGER NOT NULL,
    set_number INTEGER NOT NULL DEFAULT 1,
    player1_score INTEGER DEFAULT 501, -- remaining score
    player2_score INTEGER DEFAULT 501,
    winner_id UUID REFERENCES players(id),
    status VARCHAR(50) DEFAULT 'in_progress', -- 'in_progress', 'completed'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP
);

-- Individual throws/turns
CREATE TABLE throws (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    game_id UUID REFERENCES games(id) ON DELETE CASCADE,
    player_id UUID REFERENCES players(id),
    throw_number INTEGER NOT NULL, -- 1, 2, 3 within the turn
    turn_number INTEGER NOT NULL, -- which turn in the leg
    score INTEGER NOT NULL, -- points scored in this throw
    multiplier INTEGER DEFAULT 1, -- 1=single, 2=double, 3=triple
    is_bust BOOLEAN DEFAULT FALSE,
    remaining_score INTEGER, -- score after this throw
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tournament statistics (per tournament)
CREATE TABLE tournament_stats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tournament_id UUID REFERENCES tournaments(id) ON DELETE CASCADE,
    player_id UUID REFERENCES players(id) ON DELETE CASCADE,
    
    -- Tournament performance
    final_position INTEGER,
    matches_played INTEGER DEFAULT 0,
    matches_won INTEGER DEFAULT 0,
    legs_played INTEGER DEFAULT 0,
    legs_won INTEGER DEFAULT 0,
    
    -- Throwing stats for this tournament
    total_throws INTEGER DEFAULT 0,
    total_score INTEGER DEFAULT 0,
    average_score DECIMAL(5,2) DEFAULT 0,
    best_finish INTEGER,
    checkout_percentage DECIMAL(5,2) DEFAULT 0,
    first_9_average DECIMAL(5,2) DEFAULT 0,
    
    -- Accuracy stats
    singles_hit INTEGER DEFAULT 0,
    doubles_hit INTEGER DEFAULT 0,
    triples_hit INTEGER DEFAULT 0,
    
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(tournament_id, player_id)
);

-- League statistics (aggregated across all tournaments in league)
CREATE TABLE league_stats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    league_id UUID REFERENCES leagues(id) ON DELETE CASCADE,
    player_id UUID REFERENCES players(id) ON DELETE CASCADE,
    
    -- Overall league performance
    total_matches_played INTEGER DEFAULT 0,
    total_matches_won INTEGER DEFAULT 0,
    total_legs_played INTEGER DEFAULT 0,
    total_legs_won INTEGER DEFAULT 0,
    
    -- Aggregated throwing stats
    total_throws INTEGER DEFAULT 0,
    total_score INTEGER DEFAULT 0,
    overall_average DECIMAL(5,2) DEFAULT 0,
    best_finish INTEGER,
    overall_checkout_percentage DECIMAL(5,2) DEFAULT 0,
    
    -- League specific achievements
    tournament_wins INTEGER DEFAULT 0,
    podium_finishes INTEGER DEFAULT 0, -- top 3 finishes
    
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(league_id, player_id)
);

-- Indexes for performance
CREATE INDEX idx_leagues_status ON leagues(status);
CREATE INDEX idx_leagues_season ON leagues(season);
CREATE INDEX idx_league_standings_league ON league_standings(league_id);
CREATE INDEX idx_league_standings_position ON league_standings(league_id, current_position);
CREATE INDEX idx_tournaments_league ON tournaments(league_id);
CREATE INDEX idx_tournaments_status ON tournaments(status);
CREATE INDEX idx_tournament_players_tournament ON tournament_players(tournament_id);
CREATE INDEX idx_tournament_players_player ON tournament_players(player_id);
CREATE INDEX idx_matches_tournament ON matches(tournament_id);
CREATE INDEX idx_matches_status ON matches(status);
CREATE INDEX idx_games_match ON games(match_id);
CREATE INDEX idx_throws_game ON throws(game_id);
CREATE INDEX idx_throws_player ON throws(player_id);
CREATE INDEX idx_tournament_stats_tournament ON tournament_stats(tournament_id);
CREATE INDEX idx_tournament_stats_player ON tournament_stats(player_id);
CREATE INDEX idx_league_stats_league ON league_stats(league_id);
CREATE INDEX idx_league_stats_player ON league_stats(player_id);

-- Insert sample data for testing
INSERT INTO players (name, email, nickname) VALUES 
    ('John "The Power" Smith', 'john@example.com', 'The Power'),
    ('Sarah "Bullseye" Johnson', 'sarah@example.com', 'Bullseye'),
    ('Mike "Lightning" Wilson', 'mike@example.com', 'Lightning'),
    ('Emma "Dart Queen" Davis', 'emma@example.com', 'Dart Queen'),
    ('Tom "The Machine" Brown', 'tom@example.com', 'The Machine'),
    ('Lisa "Sharp Shooter" White', 'lisa@example.com', 'Sharp Shooter'),
    ('Chris "Thunder" Miller', 'chris@example.com', 'Thunder'),
    ('Anna "Precision" Taylor', 'anna@example.com', 'Precision');

-- Create a sample league
INSERT INTO leagues (name, description, season, status, start_date, end_date) VALUES 
    ('Winter Championship 2025', 'Premier winter darts league with weekly tournaments', '2025 Winter', 'active', '2025-01-01', '2025-03-31');

-- Add players to the league
INSERT INTO league_players (league_id, player_id)
SELECT l.id, p.id 
FROM leagues l, players p 
WHERE l.name = 'Winter Championship 2025';

-- Initialize league standings
INSERT INTO league_standings (league_id, player_id, current_position, previous_position)
SELECT lp.league_id, lp.player_id, 
       ROW_NUMBER() OVER (ORDER BY RANDOM()) as current_position,
       ROW_NUMBER() OVER (ORDER BY RANDOM()) as previous_position
FROM league_players lp
JOIN leagues l ON lp.league_id = l.id
WHERE l.name = 'Winter Championship 2025';

-- Create sample tournaments
INSERT INTO tournaments (league_id, name, description, type, status, tournament_number, scheduled_date)
SELECT l.id, 
       'Tournament ' || generate_series(1, 3),
       'Week ' || generate_series(1, 3) || ' tournament',
       'single_elimination',
       CASE generate_series(1, 3)
           WHEN 1 THEN 'completed'
           WHEN 2 THEN 'in_progress'
           ELSE 'setup'
       END,
       generate_series(1, 3),
       CURRENT_DATE + (generate_series(1, 3) - 1) * INTERVAL '7 days'
FROM leagues l
WHERE l.name = 'Winter Championship 2025';