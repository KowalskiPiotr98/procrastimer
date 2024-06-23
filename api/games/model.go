package games

import (
	"github.com/gofrs/uuid"
	"procrastimer/common/repository"
)

// Game stores information about a game the user has played or intends to play.
type Game struct {
	Id      uuid.UUID
	Name    string
	AddedOn int64
}

var _ repository.CreatableWithId = (*Game)(nil)

func (g *Game) SetId(id uuid.UUID) {
	g.Id = id
}

func NewGame(name string) *Game {
	return &Game{
		Name: name,
	}
}
