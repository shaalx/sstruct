package bean

import (
	"github.com/shaalx/sstruct/utils"
)

type News struct {
	Content     interface{} `content`
	UnixDate    int64
	DisplayDate string
}

func (self *News) Init() {
	self.UnixDate = utils.NowToUnix()
	self.DisplayDate = utils.UnixFormatS(self.UnixDate)
}