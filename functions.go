package sqlite

import (
	"log"

	bus "github.com/hexades/hexabus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type OpenFunction = func(event *Open, repo *repository)
type ReadFunction = func(event *Read, repo *repository)
type InsertFunction = func(event *Insert, repo *repository)
type UpdateFunction = func(event *Update, repo *repository)
type DeleteFunction = func(event *Delete, repo *repository)

var BasicOpenFunc = func(event *Open, repo *repository) {
	log.Println("Received open: ",event)
	db, err := gorm.Open(sqlite.Open(event.connection), &gorm.Config{})
	if err!=nil { panic(err)}
	repo.db = db
	event.Send(bus.NewResponse(repo.db, err))
}
var ReadFirstFunc = func(event *Read, repo *repository) {
	value := event.queryValue
	tx := repo.db.First(value)
	event.Send(bus.NewResponse(value, tx.Error))
}

var BasicInsertFunc = func(event *Insert, repo *repository) {
	value := event.value
	//TODO Remove this after initial development
	repo.db.AutoMigrate(&value)
	tx := repo.db.Create(value)
	event.Send(bus.NewResponse(value, tx.Error))
}

var BasicUpdateFunc = func(event *Update, repo *repository) {
	value := event.value
	tx := repo.db.Updates(value)
	event.Send(bus.NewResponse(value, tx.Error))
}

var BasicDeleteFunc = func(event *Update, repo *repository) {
	value := event.value
	tx := repo.db.Delete(value)
	event.Send(bus.NewResponse(value, tx.Error))
}