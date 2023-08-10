package sqlite

import (
	"log"

	bus "github.com/hexades/hexabus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type OpenFunction = func(event *Open, repo *repository) bus.Response
type SQLiteFunction = func(data *model, repo *repository) bus.Response

var BasicOpenFunc = func(event *Open, repo *repository) bus.Response {
	log.Println("Received open: ", event)
	db, err := gorm.Open(sqlite.Open(event.connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo.db = db
	return bus.NewResponse(repo.db, err)
}
var ReadFirstFunc = func(data *model, repo *repository) bus.Response {
	value := data.value
	tx := repo.db.First(value)
	return bus.NewResponse(value, tx.Error)
}

var BasicInsertFunc = func(data *model, repo *repository) bus.Response {
	value := data.value
	//TODO Remove this after initial development
	repo.db.AutoMigrate(&value)
	tx := repo.db.Create(value)
	return bus.NewResponse(value, tx.Error)
}

var BasicUpdateFunc = func(data *model, repo *repository) bus.Response {
	value := data.value
	tx := repo.db.Updates(value)
	return bus.NewResponse(value, tx.Error)
}

var BasicDeleteFunc = func(data *model, repo *repository) bus.Response {
	value := data.value
	tx := repo.db.Delete(value)
	return bus.NewResponse(value, tx.Error)
}
