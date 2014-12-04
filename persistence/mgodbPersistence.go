package persistence

import (
	"encoding/json"
	"fmt"
	// "github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/mgo/bson"
	"github.com/shaalx/sstruct/mgodb"
)

type MgoPersistence struct {
	Persistence
	MgoDB *mgodb.DB
}

func init() {
	fmt.Println("init...")
}

func (m *MgoPersistence) init() {
	fmt.Println("mgoPersistence init...")
}

func (m MgoPersistence) Do(bs []byte) bool {
	var i interface{}
	json.Unmarshal(bs, &i)
	return m.MgoDB.Save(i)
}

func (m *MgoPersistence) QueryOne() *bson.M {
	return m.MgoDB.Select(nil)
}
