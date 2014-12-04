package persistence

import (
	"encoding/json"
	"github.com/shaalx/sstruct/log"
	"github.com/shaalx/sstruct/mgodb"
	"github.com/shaalx/sstruct/pkg3/mgo/bson"
)

type MgoPersistence struct {
	Persistence
	MgoDB *mgodb.DB
}

func (m *MgoPersistence) init() {
	log.LOGS.Notice("mgoPersistence init...")
}

func (m MgoPersistence) Do(bs []byte) bool {
	var i interface{}
	err := json.Unmarshal(bs, &i)
	if log.IsError("{mongodb data unmarshal}", err) {
		return false
	}
	return m.MgoDB.Save(i)
}

func (m *MgoPersistence) QueryOne() *bson.M {
	return m.MgoDB.Select(nil)
}
