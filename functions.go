package hsqlite

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type Executable = func(data *Model, repo *Repository) Response

var BasicOpenFunc = func(data *Model, repo *Repository) Response {
	log.Println("Received open: ", data)
	db, err := gorm.Open(sqlite.Open(data.String()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repo.db = db
	return NewResponse(repo.db, err)
}
var ReadFirstFunc = func(data *Model, repo *Repository) Response {
	value := data.Value
	tx := repo.db.First(value)
	return NewResponse(value, tx.Error)
}

var BasicInsertFunc = func(data *Model, repo *Repository) Response {
	value := data.Value
	//TODO Remove this after initial development
	repo.db.AutoMigrate(&value)
	tx := repo.db.Create(value)
	return NewResponse(value, tx.Error)
}

var BasicUpdateFunc = func(data *Model, repo *Repository) Response {
	value := data.Value
	tx := repo.db.Updates(value)
	return NewResponse(value, tx.Error)
}

var BasicDeleteFunc = func(data *Model, repo *Repository) Response {
	value := data.Value
	tx := repo.db.Delete(value)
	return NewResponse(value, tx.Error)
}
