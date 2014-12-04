package persistence

import (
	"encoding/json"
	// "github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/mgodb"
)

type MgoPersistence struct {
	Persistence
	Server []string
}

func (m MgoPersistence) Do(bs []byte) bool {
	// server := []string{"", "newsmgo", "firstbanner"}
	dbserver := mgodb.SetLocalDB(m.Server...)
	defer dbserver.Close()
	var i interface{}
	json.Unmarshal(bs, &i)
	return dbserver.Save(i)
}
