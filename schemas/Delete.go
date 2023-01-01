package schemas

import (
	"api/database"
)

type Remover struct {
	Name string
}

func Remove(name string) error {
	s := Remover{Name: name}

	return database.Delete("Catalog", s)
}
