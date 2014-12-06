package mgodb

import (
	"github.com/shaalx/sstruct/log"
	"github.com/shaalx/sstruct/pkg3/mgo"
	"strings"
)

// 数据库结构体
type DB struct {
	db            string
	collectionStr string
	session       *mgo.Session
	collection    *mgo.Collection
}

/*设置数据库的：
服务器 str[0];数据库 str[1]; 集合 str[2]*/
func SetDB(str ...string) *DB {
	db := new(DB)
	dialUrl := "mongodb://database:password@" + str[0] + "/collections"
	session, err := mgo.Dial(dialUrl)
	if log.IsError("{connect mongodb}", err) {
		return nil
	}
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(str[1]).C(str[2])
	db.db = str[1]
	db.collectionStr = str[2]
	db.session = session
	db.collection = collection
	return db
}

func SetLocalDB(str ...string) *DB {
	db := new(DB)
	session, err := mgo.Dial("127.0.0.1:27017")
	if log.IsError("{connect mongodb}", err) {
		return nil
	}
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(str[1]).C(str[2])
	db.db = str[1]
	db.collectionStr = str[2]
	db.session = session
	db.collection = collection
	return db
}

func (d *DB) Session() *mgo.Session {
	return d.session
}

func (d *DB) Collection() *mgo.Collection {
	return d.collection
}

/*关闭数据库*/
func (d *DB) Close() {
	d.session.Close()
}

/*返回当前数据库集合（collection）的string*/
func (d *DB) ToString() string {
	return strings.Join([]string{d.db, d.collectionStr}, "/")
}
