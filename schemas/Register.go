package schemas

import (
	"api/database"
)

type Register struct {
	Name        string `json:"name" bjson:"name"`
	Description string `json:"description" bjson:"description"`
	Genre       string `json:"genre" bjson:"genre"`
}

func Create(name string, description string, genre string) error {
	s := Register{Name: name, Description: description, Genre: genre}

	return database.Insert("Catalog", s)
}
