package bean

import (
	"encoding/json"
	"github.com/shaalx/sstruct/service/log"
	"github.com/shaalx/sstruct/utils"
)

type News struct {
	Content     interface{} `content`
	Notice      string
	UnixDate    int64
	DisplayDate string
}

func (self *News) Init() {
	self.UnixDate = utils.NowToUnix()
	self.DisplayDate = utils.UnixFormatS(self.UnixDate)
}

func (self *News) InitWithUnixDate(date int64) string {
	self.UnixDate = date
	self.DisplayDate = utils.UnixFormatS(self.UnixDate)
	return self.DisplayDate
}

func Bytes2News(data []byte) *News {
	news := News{}
	err := json.Unmarshal(data, &news)
	if log.IsError("bytes --> News error.", err) {
		return nil
	}
	return &news
}
