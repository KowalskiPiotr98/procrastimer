package games

import (
	"github.com/gofrs/uuid"
	"procrastimer/common/repository"
)

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
