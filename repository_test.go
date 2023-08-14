package sqlite

import (
	"fmt"
	"log"
	"os"

	//"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQLiteRepositorySuite(t *testing.T) {
	_ = os.Remove("test_sqlite.db")
	newRepository()
	openDB(t)
	insertData(t)
	readData(t)
	updateData(t)
}

var td = &TestData{Identifier: "foo", SomeValue: "bar"}

func openDB(t *testing.T) {
	evt := NewEvent("test_sqlite.db", BasicOpenFunc)
	sendEvent(evt)
	log.Println("Sent open event:", evt)
	response := evt.Receive()
	assert.Nil(t, response.Err)
}

func insertData(t *testing.T) {
	t.Log("Insert ")
	evt := NewEvent(td, BasicInsertFunc)
	sendEvent(evt)
	response := evt.Receive()
	assert.NotNil(t, response)
	assert.Nil(t, response.Err)
	t.Log("End Insert")
}

func readData(t *testing.T) {
	t.Log("Read ")
	query := &TestData{Identifier: "foo"}
	rd := NewEvent(query, ReadFirstFunc)
	sendEvent(rd)
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
	sendEvent(event)
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
