package test

import "github.com/gofrs/uuid"

func GetRandomUuid() uuid.UUID {
	id, err := uuid.NewV6()
	PanicOnErr(err)
	return id
}
