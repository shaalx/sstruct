package structs

import (
	"github.com/shaalx/sstruct/common"
)

type News struct {
	Content     interface{} `content`
	UnixDate    int64
	DisplayDate string
}

func (self *News) Init() {
	self.UnixDate = common.NowToUnix()
	self.DisplayDate = common.UnixFormatS(self.UnixDate)
}
