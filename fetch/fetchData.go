package fetch

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/shaalx/sstruct/log"
	"net/http"
)

/*根据给定的URL,fetch the data*/
func Do(url, ipaddr string) []byte {
	request := httplib.Get(url)
	request.SetTransport(newTransport(ipaddr))
	request.Header("Host", "itunes.apple.com")
	request.Header("X-Apple-Store-Front", "143465-19,21 t:native")
	request.Header("Accept", "*/*")
	request.Header("Accept-Language", "zh-cn")
	request.Header("X-Dsid", "1458643138")
	request.Header("Connection", "keep-alive")
	request.Header("Proxy-Connection", "keep-alive")
	request.Header("Design-Agent", "AppStore/2.0 iOS/7.1.1 model/iPod5,1 build/11D201 (4; dt:81)")
	bs, err := request.Bytes()
}

func newTransport(ipaddr string) *http.Transport {
	transport :=
		&http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				//本地地址  ipaddr是本地外网IP
				lAddr, err := net.ResolveTCPAddr(netw, ipaddr+":0")
				if err != nil {
					return nil, err
				}
				//被请求的地址
				rAddr, err := net.ResolveTCPAddr(netw, addr)
				if err != nil {
					return nil, err
				}
				conn, err := net.DialTCP(netw, lAddr, rAddr)
				if err != nil {
					return nil, err
				}
				deadline := time.Now().Add(35 * time.Second)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		}
	return transport
}
