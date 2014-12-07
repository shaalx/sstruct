package persistence

import (
	"encoding/json"
	"github.com/shaalx/sstruct/bean"
	"github.com/shaalx/sstruct/persistence/mgodb"
	"github.com/shaalx/sstruct/pkg3/mgo/bson"
	"github.com/shaalx/sstruct/service/log"
)

type MgoPersistence struct {
	Persistence
	MgoDB *mgodb.DB
}

func (m *MgoPersistence) init() {
	log.LOGS.Notice("mgoPersistence init...")
}

func (m MgoPersistence) Do(bs []byte) bool {
	news := bean.News{}
	news.Init()
	err := json.Unmarshal(bs, &news.Content)
	if log.IsError("{mongodb data unmarshal news}", err) {
		return false
	}
	return m.MgoDB.Save(news)
}

func (m *MgoPersistence) QueryOne() *bson.M {
	return m.MgoDB.Select(nil)
}

func (m *MgoPersistence) QueryNewsOne(selector bson.M) *bean.News {
	return m.MgoDB.SelectNews(selector)
}

func (m *MgoPersistence) QuerySortedNewsOne(selector bson.M, srotor ...string) *bean.News {
	return m.MgoDB.SelectSortNews(selector, srotor...)
}
