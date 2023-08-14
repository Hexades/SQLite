package hsqlite

import (
	"fmt"
)

type Model struct {
	Value any
}

func (m *Model) String() string {
	return fmt.Sprintf("%v", m.Value)
}

func NewEvent(value any, executable Executable) Event {
	return &EventModel{data: &Model{Value: value}, executable: executable}
}

type Event interface {
	Execute(repo *Repository)
	Send(Response)
	Receive() Response
}

type EventModel struct {
	responseChannel chan Response
	data            *Model
	executable      Executable
}

func (e *EventModel) Execute(repo *Repository) {
	e.Send(e.executable(e.data, repo))
}

func (e *EventModel) getChannel() chan Response {
	if e.responseChannel == nil {
		e.responseChannel = make(chan Response, 1)
	}
	return e.responseChannel
}

func (e *EventModel) Send(val Response) {
	e.getChannel() <- val
}

func (e *EventModel) Receive() Response {
	return <-e.getChannel()
}
func NewResponse(value any, err error) Response {
	return Response{value, err}
}

type Response struct {
	Value any
	Err   error
}
