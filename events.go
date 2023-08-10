package sqlite

import (
	"fmt"
	bus "github.com/hexades/hexabus"
)

type model struct {
	value any
}

func (m *model) String() string {
	return fmt.Sprintf("%v", m.value)
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
