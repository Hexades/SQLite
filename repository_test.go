package sqlite

import (
	"fmt"
	"log"
	"os"

	//"reflect"
	"testing"

	//"time"

	bus "github.com/hexades/hexabus"
	"github.com/stretchr/testify/assert"
)

func TestSQLiteRepositorySuite(t *testing.T) {
	_ = os.Remove("test_sqlite.db")
	openDB(t)
	insertData(t)
	readData(t)
	updateData(t)
}

var td = &TestData{Identifier: "foo", SomeValue: "bar"}

func openDB(t *testing.T) {
	b := bus.Get()
	NewRepository()
	evt := NewOpen("test_sqlite.db", BasicOpenFunc)
	b.SendRepositoryEvent(evt)
	response := evt.Receive()
	assert.Nil(t, response.Err)
}

func insertData(t *testing.T) {
	t.Log("Insert ")
	evt := NewInsert(td, BasicInsertFunc)
	bus.Get().SendRepositoryEvent(evt)
	response := evt.Receive()
	assert.NotNil(t, response)
	assert.Nil(t, response.Err)
	t.Log("End Insert")
}

func readData(t *testing.T) {
	t.Log("Read ")
	query := &TestData{Identifier: "foo"}
	rd := NewRead(query, ReadFirstFunc)
	bus.Get().SendRepositoryEvent(rd)
	response := rd.Receive()
	assert.NotNil(t, response)
	fmt.Println("Response Value: ", response)

	switch val := response.Value.(type) {
	case TestData:
		log.Println(val)
	}
	assert.Equal(t, td, response.Value)
	t.Log("End Read")
}
func updateData(t *testing.T) {
	t.Log("Update ")
	event := NewUpdate(&TestData{Identifier: "foo", SomeValue: "Doodah"}, BasicUpdateFunc)
	bus.Get().SendRepositoryEvent(event)
	response := event.Receive()
	assert.NotNil(t, response)
	fmt.Println("Response Value: ", response)

	switch val := response.Value.(type) {
	case TestData:
		log.Println(val)
	}
	assert.Equal(t, event.value, response.Value)
	t.Log("End Read")
}

type TestData struct {
	//gorm.Model
	Identifier string `gorm:"primaryKey"`
	SomeValue  string `gorm:"<-"`
}
