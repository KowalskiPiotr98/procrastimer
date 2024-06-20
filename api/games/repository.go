package games

import (
	"github.com/KowalskiPiotr98/gotabase"
	"github.com/gofrs/uuid"
	"procrastimer/common/repository"
)

var (
	getDatabase      = func() gotabase.Connector { return gotabase.GetConnection() }
	beginTransaction = func() (*gotabase.Transaction, error) { return gotabase.BeginTransaction() }
)

// GetGame returns a Game pointer with given id, or error if not found.
func GetGame(id uuid.UUID, userId uuid.UUID) (*Game, error) {
	query := "select id, name, added_on from games where id = $1 and user_id = $2"
	result, err := getDatabase().QueryRow(query, id, userId)
	if err != nil {
		return nil, err
	}
	return scanGame(result)
}

func scanGame(row gotabase.Row) (*Game, error) {
	game := &Game{}
	if err := row.Scan(&game.Id, &game.Name, &game.AddedOn); err != nil {
		if repository.IsDataNotFoundErr(err) {
			return nil, repository.DataNotFoundErr
		}
		return nil, err
	}
	return game, nil
}
