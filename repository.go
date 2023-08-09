package sqlite

import (
	"log"

	"github.com/enjekt/hexabus/bus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type repository struct {
	db *gorm.DB
}

func NewRepository() {
	bus.Get().AddRepositoryListener(new(repository))
}

func (r *repository) OnRepositoryEvent(repositoryChannel <-chan bus.RepositoryEvent) {

	for repoEvent := range repositoryChannel {

		switch evt := repoEvent.(type) {
		case Open:
			r.Open(evt)
		case Read:
			r.Read(evt)
		case Insert:
			r.Insert(evt)
		case Delete:
			r.Delete(evt)

		}
	}
}

func (r *repository) Open(evt Open) {
	log.Println("Receive open: ", evt)
	
	db, err := gorm.Open(sqlite.Open(evt.OpenConnection()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r.db = db

}

func (r *repository) Insert(evt Insert) {
	value := evt.InsertQuery()
	log.Println("Insert: ", &value)
	r.db.AutoMigrate(&value)
	tx := r.db.Create(value)
	evt.Send(bus.NewResponse(value,tx.Error))
}

func (r *repository) Read(evt Read) {
	value := evt.ReadQuery()
	tx := r.db.First(value)
	evt.Send(bus.NewResponse(value,tx.Error))

}

func (r *repository) Delete(evt Delete) {
	value := evt.DeleteQuery()
	tx := r.db.Delete(value)
	evt.Send(bus.NewResponse(value,tx.Error))

}
