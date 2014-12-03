package fetch

import (
	// "./gojson"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/go-xorm/xorm"
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
	// "os"
	"strconv"
	"strings"
	"time"
)

func CrawlStoreApp(appurl AppUrl) {
	crawlStore(appurl.Url)
}

func CrawlStoreApps(appurls ...AppUrl) {
	for _, it := range appurls {
		CrawlStoreApp(it)
	}
}

/*根据给定的URL，将源数据加盖时间戳，然后存储到mongoDB中*/
func crawlStore(url string) {
	request := httplib.Get(url)
	request.Header("Host", "itunes.apple.com")
	request.Header("X-Apple-Store-Front", "143465-19,21 t:native")
	request.Header("Accept", "*/*")
	request.Header("Accept-Language", "zh-cn")
	request.Header("X-Dsid", "1458643138")
	request.Header("Connection", "keep-alive")
	request.Header("Proxy-Connection", "keep-alive")
	request.Header("Design-Agent", "AppStore/2.0 iOS/7.1.1 model/iPod5,1 build/11D201 (4; dt:81)")

	str, _ := request.String()
	v := setStringByStruct(str)
	onlyStore(v)
}
