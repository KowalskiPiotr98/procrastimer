package games

import (
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"procrastimer/common/repository"
	"procrastimer/common/slice"
	"procrastimer/test"
	"testing"
)

func makeTestUser() uuid.UUID {
	userId := test.GetRandomUuid()
	_, _ = getDatabase().Exec("insert into users (id, email) values ($1, $2)", userId, "test")
	return userId
}

func insertTestGame(game *Game, userId uuid.UUID) {
	query := "insert into games (name, user_id) values ($1, $2) returning id"
	result, _ := getDatabase().QueryRow(query, game.Name, userId)
	_ = result.Scan(&game.Id)
}

func TestGetGame_GameFound_ReturnsValidData(t *testing.T) {
	_ = test.GetDatabaseWithCleanup(t)
	userId := makeTestUser()
	game := NewGame("test game")
	insertTestGame(game, userId)

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
	userId := makeTestUser()
	game := NewGame("test game")
	insertTestGame(game, userId)

	_, err := GetGame(game.Id, test.GetRandomUuid())

	assert.Equal(t, repository.DataNotFoundErr, err)
}

func TestGetGames_ReturnsOnlyAuthorised(t *testing.T) {
	_ = test.GetDatabaseWithCleanup(t)
	userId := makeTestUser()
	game1 := NewGame("test game 1")
	game2 := NewGame("test game 2")
	game3 := NewGame("test game 3")
	game4 := NewGame("test game 4")
	insertTestGame(game1, userId)
	insertTestGame(game2, userId)
	insertTestGame(game3, test.GetRandomUuid())
	insertTestGame(game4, test.GetRandomUuid())

	games, err := GetGames(userId, 0, 100)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(games))
	ids := slice.Map(games, func(game *Game) uuid.UUID { return game.Id })
	assert.Contains(t, ids, game1.Id)
	assert.Contains(t, ids, game2.Id)
}

func TestCreateGame_NewGame_AddedToDatabase(t *testing.T) {
	_ = test.GetDatabaseWithCleanup(t)
	userId := makeTestUser()
	game := NewGame("test game")

	err := CreateGame(game, userId)

	assert.Nil(t, err)
	dbGame, err := GetGame(game.Id, userId)
	test.PanicOnErr(err)
	assert.NotEqual(t, 0, dbGame.AddedOn)
	dbGame.AddedOn = 0
	assert.Equal(t, game, dbGame)
}

func TestCreateGame_DuplicateGame_ReturnsError(t *testing.T) {
	_ = test.GetDatabaseWithCleanup(t)
	userId := makeTestUser()
	test.PanicOnErr(CreateGame(NewGame("test game"), userId))
	game := NewGame("test game")

	err := CreateGame(game, userId)

	assert.Equal(t, repository.AlreadyExistsErr, err)
	dbGames, err := GetGames(userId, 0, 100)
	test.PanicOnErr(err)
	assert.Equal(t, 1, len(dbGames))
}
