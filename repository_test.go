package sqlite

import (
	"fmt"
	"log"
	"os"
	//"reflect"
	"testing"

	//"time"

	"github.com/hexades/hexabus"
	"github.com/stretchr/testify/assert"
)


func TestSQLiteRepositorySuite(t *testing.T) {
	_ = os.Remove("test_sqlite.db")
	b := bus.Get()
	NewRepository()
	b.SendRepositoryEvent(NewOpen("test_sqlite.db"))
	insertData(t)
	readData(t)
}

var td = &TestData{Identifier: "foo", SomeValue: "bar"}

func insertData(t *testing.T) {
	t.Log("Insert ")
	evt:=NewInsert(td)
	bus.Get().SendRepositoryEvent(evt)
	response:=evt.Receive()
	assert.NotNil(t,response)
	assert.Nil(t,response.Err)
	t.Log("End Insert")
}

func readData(t *testing.T) {
	t.Log("Read ")
	query:=&TestData{Identifier: "foo"}
	rd := NewRead(query)
	bus.Get().SendRepositoryEvent(rd)
	response:=rd.Receive()
	assert.NotNil(t,response)
	fmt.Println("Response Value: ",response)

	switch val:=response.Value.(type) {
		case TestData:
			log.Println(val)
	}
	assert.Equal(t,td,response.Value)
	t.Log("End Read")
}

type TestData struct {
	//gorm.Model
	Identifier string `gorm:"primaryKey"`
	SomeValue  string `gorm:"<-"`
}
