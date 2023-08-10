package sqlite

import (
	bus "github.com/hexades/hexabus"
)

type model struct {
	value any
}

type Event interface {
	Execute(repo *repository)
}

type SQLiteEvent struct {
	bus.RequestResponseEvent
	data       *model
	executable SQLiteFunction
}

func (e *SQLiteEvent) Execute(repo *repository) {
	e.Send(e.executable(e.data, repo))
}

func NewEvent(value any, executable SQLiteFunction) *SQLiteEvent {
	return &SQLiteEvent{data: &model{value: value}, executable: executable}
}

type OpenEvent struct {
	bus.RequestResponseEvent
	data       *Open
	executable OpenFunction
}

func (e *OpenEvent) Execute(repo *repository) {
	e.Send(e.executable(e.data, repo))
}

func NewOpen(connection string, executable OpenFunction) *OpenEvent {
	return &OpenEvent{data: &Open{connection: connection}, executable: executable}
}

type Open struct {
	bus.RequestResponseEvent
	connection string
}
