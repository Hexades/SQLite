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
	NewRepository()
	evt := NewEvent("test_sqlite.db", BasicOpenFunc)
	bus.SendRepositoryEvent(evt)
	response := evt.Receive()
	assert.Nil(t, response.Err)
}

func insertData(t *testing.T) {
	t.Log("Insert ")
	evt := NewEvent(td, BasicInsertFunc)
	bus.SendRepositoryEvent(evt)
	response := evt.Receive()
	assert.NotNil(t, response)
	assert.Nil(t, response.Err)
	t.Log("End Insert")
}

func readData(t *testing.T) {
	t.Log("Read ")
	query := &TestData{Identifier: "foo"}
	rd := NewEvent(query, ReadFirstFunc)
	bus.SendRepositoryEvent(rd)
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
	updateData := &TestData{Identifier: "foo", SomeValue: "Doodah"}
	event := NewEvent(updateData, BasicUpdateFunc)
	bus.SendRepositoryEvent(event)
	response := event.Receive()
	assert.NotNil(t, response)

	assert.Equal(t, updateData, response.Value)
	t.Log("End Read")
}

type TestData struct {
	//gorm.Model
	Identifier string `gorm:"primaryKey"`
	SomeValue  string `gorm:"<-"`
}
