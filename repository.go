package sqlite

import (
	"gorm.io/gorm"
	"log"
)


type Repository struct {
	db *gorm.DB
}

func NewRepository() {
	r := new(Repository)
	log.Println("Created new respository: ", r)
	AddListener(r)
	log.Println("Added repository as listener")
}

func (r *Repository) OnEvent(repositoryChannel <-chan Event) {

	for evt := range repositoryChannel {
		log.Println("Execute event", evt)
		evt.Execute(r)
	}
}
