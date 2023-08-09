package sqlite

import (
	"github.com/hexades/hexabus"
)

const dbtype = "sqlite"


func NewOpen(connection string) Open {
	return &open{connection: connection}
}

type open struct {
	bus.RepositoryID
	bus.Event
	connection string
}

type Open interface {
	bus.RepositoryRequestEvent
	OpenConnection() string
}

func (o *open) OpenConnection() string {
	return o.connection
}

func NewRead(queryValue any) Read {
	return &read{queryValue: queryValue}
}

type Read interface {
	bus.RepositoryRequestEvent
	ReadQuery() any
}
type read struct {
	bus.Event
	queryValue any
}

func (r *read) ReadQuery() any {
	return r.queryValue
}

func NewInsert(value any) Insert {
	return &insert{value: value}
}

type insert struct {
	bus.Event
	value any
}
type Insert interface {
	bus.RepositoryRequestEvent
	InsertQuery() any
}

func (evt *insert) InsertQuery() any {
	return evt.value
}
func NewUpdate(value any) Update {
	return &update{value: value}
}

type update struct {
	bus.Event
	value any
}

type Update interface {
	bus.RepositoryRequestEvent
	UpdateQuery() any
}

func (evt *update) UpdateQuery() any {
	return evt.value
}
func NewDelete(value any) Delete {
	return &delete{value: value}
}

type Delete interface {
	bus.RepositoryRequestEvent
	DeleteQuery() any
}

type delete struct {
	bus.Event
	value any
}

func (evt *delete) DeleteQuery() any {
	return evt.value
}
