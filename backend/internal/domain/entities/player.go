package entities

import (
	"time"
	"github.com/google/uuid"
)

type Player struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     *string    `json:"email,omitempty"`
	AvatarURL *string    `json:"avatar_url,omitempty"`
	Nickname  *string    `json:"nickname,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewPlayer creates a new player with validation
func NewPlayer(name string, email, nickname *string) (*Player, error) {
	if name == "" {
		return nil, ErrInvalidPlayerName
	}

	return &Player{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Nickname:  nickname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// UpdateProfile updates player profile information
func (p *Player) UpdateProfile(name string, email, nickname, avatarURL *string) error {
	if name == "" {
		return ErrInvalidPlayerName
	}

	p.Name = name
	p.Email = email
	p.Nickname = nickname
	p.AvatarURL = avatarURL
	p.UpdatedAt = time.Now()

	return nil
}

// DisplayName returns the preferred display name for the player
func (p *Player) DisplayName() string {
	if p.Nickname != nil && *p.Nickname != "" {
		return *p.Nickname
	}
	return p.Name
}