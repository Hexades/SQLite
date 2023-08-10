package sqlite

import (
	bus "github.com/hexades/hexabus"
)

type SQLiteEvent interface {
	Execute(repo *repository)
}

func NewOpen(connection string, openFunc OpenFunction) *Open {
	return &Open{connection: connection, openFunc: &openFunc}
}

type Open struct {
	bus.RequestResponseEvent
	connection string
	openFunc   *OpenFunction
}

func (e *Open) Execute(repo *repository) {
	(*e.openFunc)(e, repo)
}

func NewRead(queryValue any, readFunc ReadFunction) *Read {
	return &Read{queryValue: queryValue, readFunc: &readFunc}
}

type Read struct {
	bus.RequestResponseEvent
	readFunc   *ReadFunction
	queryValue any
}

func (e *Read) Execute(repo *repository) {
	(*e.readFunc)(e, repo)
}

func NewInsert(value any, insertFunc InsertFunction) *Insert {
	return &Insert{value: value, insertFunc: &insertFunc}
}

type Insert struct {
	bus.RequestResponseEvent
	value      any
	insertFunc *InsertFunction
}

func (e *Insert) Execute(repo *repository) {
	(*e.insertFunc)(e, repo)
}

func NewUpdate(value any, updateFunc UpdateFunction) *Update {
	return &Update{value: value, updateFunc: &updateFunc}
}

type Update struct {
	bus.RequestResponseEvent
	value      any
	updateFunc *UpdateFunction
}

func (e *Update) Execute(repo *repository) {
	(*e.updateFunc)(e, repo)
}

func NewDelete(value any, deleteFunc DeleteFunction) *Delete {
	return &Delete{value: value, deleteFunc: &deleteFunc}
}

type Delete struct {
	bus.RequestResponseEvent
	value      any
	deleteFunc *DeleteFunction
}

func (e *Delete) Execute(repo *repository) {
	(*e.deleteFunc)(e, repo)
}
