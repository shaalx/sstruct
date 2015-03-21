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

func (m MgoPersistence) Do(bs []byte, notice string) bool {
	news := bean.News{}
	news.Init()
	news.Notice = notice
	err := json.Unmarshal(bs, &news.Content)
	if log.IsError("{mongodb data unmarshal news}", err) {
		return false
	}
	return m.MgoDB.Save(news)
}

// 在某一时间内，算作某时刻
func (m MgoPersistence) DoWithUnixDate(bs []byte, notice string, date int64) bool {
	news := bean.News{}
	news.InitWithUnixDate(date)
	news.Notice = notice
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

func (m *MgoPersistence) QueryNewses(selector bson.M) []bean.News {
	return m.MgoDB.SelectNewses(selector)
}

func (m *MgoPersistence) QuerySortedLimitNNewses(selector bson.M, n int, srotor ...string) []bean.News {
	return m.MgoDB.SelectSortLimitNNewses(selector, n, srotor...)
}
