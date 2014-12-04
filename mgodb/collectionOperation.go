package mgodb

import (
	"fmt"
	"github.com/shaalx/sstruct/mgo/bson"
	// . "github.com/shaalx/sstruct/structs"
	"log"
)

// log and check the error
func Log(err error) {
	if err != nil {
		// fmt.Println(err.Error())
	}
}

// check data exists in mongodb
func (d *DB) Exist(selector bson.M) bool {
	collection := d.collection
	n, err := collection.Find(selector).Count()
	Log(err)
	if n == 0 || err != nil {
		return false
	}
	return true
}

// count data in mongodb
func (d *DB) Count(selector bson.M) int {
	collection := d.collection
	n, err := collection.Find(selector).Count()
	Log(err)
	return n
}

// save data to mongodb
func (d *DB) Save(data interface{}) bool {
	if data == nil {
		return false
	}
	collection := d.collection
	err := collection.Insert(data)
	Log(err)
	if err != nil {
		return false
	}
	fmt.Println("data save >>", d.ToString())
	return true
}

// delete data from mongodb
func (d *DB) Delete(selector bson.M) bool {
	collection := d.collection
	err := collection.Remove(selector)
	Log(err)
	if err != nil {
		return false
	}
	return true
}

// update data(selector) with change
func (d *DB) Update(selector bson.M, change interface{}) bool {
	collection := d.collection
	err := collection.Update(selector, change)
	Log(err)
	if err != nil {
		return false
	}
	// fmt.Println("data update >>", d.ToString())
	return true
}

// select one data(into interface{})
func (d *DB) Select(selector bson.M) interface{} {
	collection := d.collection
	var result interface{}
	err := collection.Find(selector).One(&result)
	Log(err)
	if err != nil {
		return nil
		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
	}
	return result
}

// select data into bson.M{}
func (d *DB) SelectAny(selector bson.M) *bson.M {
	collection := d.collection
	var result bson.M
	err := collection.Find(selector).One(&result)
	Log(err)
	if err != nil {
		return nil
	}
	return &result
}

// // select a homegroup
// func (d *DB) SelectHomegroup(selector bson.M) *Homegroup {
// 	collection := d.collection
// 	var result Homegroup
// 	err := collection.Find(selector).One(&result)
// 	Log(err)
// 	if err != nil {
// 		return nil
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 	}
// 	return &result
// }

// // select an app
// func (d *DB) SelectApp(selector bson.M) *App {
// 	collection := d.collection
// 	var result App
// 	err := collection.Find(selector).One(&result)
// 	Log(err)
// 	if err != nil {
// 		return nil
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 	}
// 	return &result
// }

// // select a digapp
// func (d *DB) SelectDigapp(selector bson.M) *Digapp {
// 	collection := d.collection
// 	var result Digapp
// 	err := collection.Find(selector).One(&result)
// 	Log(err)
// 	if err != nil {
// 		return nil
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 	}
// 	return &result
// }

// // select an apprate
// func (d *DB) SelectApprate(selector bson.M) *Apprate {
// 	collection := d.collection
// 	var result Apprate
// 	err := collection.Find(selector).One(&result)
// 	Log(err)
// 	if err != nil {
// 		return nil
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 	}
// 	return &result
// }

// select sort data into bson.M{}
func (d *DB) SelectSort(selector bson.M, sortor ...string) *bson.M {
	collection := d.collection
	var result bson.M
	err := collection.Find(selector).Sort(sortor...).One(&result)
	Log(err)
	if err != nil {
		return nil
		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
	}
	return &result
}

// // select sorted app
// func (d *DB) SelectSortApp(selector bson.M, sortor ...string) *App {
// 	collection := d.collection
// 	var result App
// 	err := collection.Find(selector).Sort(sortor...).One(&result)
// 	Log(err)
// 	if err != nil {
// 		return nil
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 	}
// 	return &result
// }

// select sorted data into bson.M{}
func (d *DB) SelectSortAny(selector bson.M, sortor ...string) *bson.M {
	collection := d.collection
	var result bson.M
	err := collection.Find(selector).Sort(sortor...).One(&result)
	Log(err)
	if err != nil {
		return nil
	}
	return &result
}

// // select sorted homegroup
// func (d *DB) SelectSortHomegroup(selector bson.M, sortor ...string) *Homegroup {
// 	collection := d.collection
// 	var result Homegroup
// 	err := collection.Find(selector).Sort(sortor...).One(&result)
// 	Log(err)
// 	if err != nil {
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 		return nil
// 	}
// 	return &result
// }

// // select sorted topchart
// func (d *DB) SelectSortTopchart(selector bson.M, sortor ...string) *Topchart {
// 	collection := d.collection
// 	var result Topchart
// 	err := collection.Find(selector).Sort(sortor...).One(&result)
// 	if err != nil {
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 		return nil
// 	}
// 	return &result
// }

// // select sorted digapp
// func (d *DB) SelectSortDigapp(selector bson.M, sortor ...string) *Digapp {
// 	collection := d.collection
// 	var result Digapp
// 	err := collection.Find(selector).Sort(sortor...).One(&result)
// 	if err != nil {
// 		// log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 		fmt.Print(".")
// 		return nil
// 	}
// 	return &result
// }

// // select sorted digapps limit within 20
// func (d *DB) SelectSortLimitDigapps(selector bson.M, sortor ...string) []Digapp {
// 	collection := d.collection
// 	var result []Digapp
// 	err := collection.Find(selector).Sort(sortor...).Limit(20).All(&result)
// 	if err != nil {
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 		return nil
// 	}
// 	return result
// }

// // select sorted apprate
// func (d *DB) SelectSortApprate(selector bson.M, sortor ...string) *Apprate {
// 	collection := d.collection
// 	var result Apprate
// 	err := collection.Find(selector).Sort(sortor...).One(&result)
// 	if err != nil {
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 		return nil
// 	}
// 	return &result
// }

// // select sorted apprate
// func (d *DB) SelectTopchartHistory(selector bson.M) []Topcharthistory {
// 	collection := d.collection
// 	var result []Topcharthistory
// 	err := collection.Find(selector).All(&result)
// 	if err != nil {
// 		log.Printf("%s select error : %v\n", d.ToString(), err.Error())
// 		return nil
// 	}
// 	return result
// }
