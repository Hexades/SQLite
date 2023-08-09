package sqlite

import (
	bus "github.com/hexades/hexabus"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	r := new(repository)
	bus.Get().AddRepositoryListener(r)
	return r
}

func (r *repository) OnRepositoryEvent(repositoryChannel <-chan bus.RepositoryEvent) {

	for repoEvent := range repositoryChannel {

		switch evt := repoEvent.(type) {
		case SQLiteEvent:
			evt.Execute(r)
		}
	}
}
