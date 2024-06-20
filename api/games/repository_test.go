package games

import (
	"github.com/stretchr/testify/assert"
	"procrastimer/common/repository"
	"procrastimer/test"
	"testing"
)

func TestGetGame_GameFound_ReturnsValidData(t *testing.T) {
	_ = test.GetDatabaseWithCleanup(t)
	userId := test.GetRandomUuid()
	_, _ = getDatabase().Exec("insert into users (id, email) values ($1, $2)", userId, "test")
	game := NewGame("test game")
	query := "insert into games (name, user_id) values ($1, $2) returning id"
	result, _ := getDatabase().QueryRow(query, "test game", userId)
	_ = result.Scan(&game.Id)

	dbGame, err := GetGame(game.Id, userId)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, dbGame.AddedOn)
	dbGame.AddedOn = 0 //zero that so that the comparison works
	assert.Equal(t, game, dbGame)
}

func TestGetGame_GameNotFound_NotFoundReturned(t *testing.T) {
	_ = test.GetDatabaseWithCleanup(t)

	_, err := GetGame(test.GetRandomUuid(), test.GetRandomUuid())

	assert.Equal(t, repository.DataNotFoundErr, err)
}

func TestGetGame_UserNotAuthorised_ReturnsNotFound(t *testing.T) {
	_ = test.GetDatabaseWithCleanup(t)
	userId := test.GetRandomUuid()
	_, _ = getDatabase().Exec("insert into users (id, email) values ($1, $2)", userId, "test")
	game := NewGame("test game")
	query := "insert into games (name, user_id) values ($1, $2) returning id"
	result, _ := getDatabase().QueryRow(query, "test game", userId)
	_ = result.Scan(&game.Id)

	_, err := GetGame(game.Id, test.GetRandomUuid())

	assert.Equal(t, repository.DataNotFoundErr, err)
}
