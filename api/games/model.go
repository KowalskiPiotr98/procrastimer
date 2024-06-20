package games

import (
	"github.com/gofrs/uuid"
)

type Game struct {
	Id      uuid.UUID
	Name    string
	AddedOn int64
}

func NewGame(name string) *Game {
	return &Game{
		Name: name,
	}
}
