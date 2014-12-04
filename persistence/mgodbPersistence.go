package persistence

import (
	"encoding/json"
	// "github.com/shaalx/sstruct/fetch"
	"github.com/shaalx/sstruct/mgodb"
)

type MgoPersistence struct {
	Persistence
}

func (m MgoPersistence) Do(bs []byte) bool {
	server := []string{"", "newsmgo", "firstbanner"}
	dbserver := mgodb.SetLocalDB(server...)
	defer dbserver.Close()
	var i interface{}
	json.Unmarshal(bs, &i)
	return dbserver.Save(i)
}
