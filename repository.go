package sqlite

import (
	"gorm.io/gorm"
	"log"
	_ "modernc.org/sqlite"
)

type Repository struct {
	db *gorm.DB
}

func newRepository() {
	r := new(Repository)
	log.Println("Created new respository: ", r)
	addRepositoryListener(r)
	log.Println("Added repository as listener")
}

func (r *Repository) onEvent(repositoryChannel <-chan Event) {

	for evt := range repositoryChannel {
		log.Println("Execute event", evt)
		evt.Execute(r)
	}
}
