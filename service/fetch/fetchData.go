package fetch

import (
	"crypto/tls"
	"github.com/shaalx/sstruct/pkg3/httplib"
	"github.com/shaalx/sstruct/service/log"
	"net"
	"net/http"
	"time"
)

/*根据给定的URL,fetch the data*/
func Do1(url, ipaddr string) []byte {
	request := httplib.Get(url)
	// request.SetTransport(newTransport(ipaddr))
	request.Header("Host", "itunes.apple.com")
	request.Header("X-Apple-Store-Front", "143465-19,21 t:native")
	request.Header("Accept", "*/*")
	request.Header("Accept-Language", "zh-cn")
	request.Header("X-Dsid", "1458643138")
	request.Header("Connection", "keep-alive")
	request.Header("Proxy-Connection", "keep-alive")
	request.Header("Design-Agent", "AppStore/2.0 iOS/7.1.1 model/iPod5,1 build/11D201 (4; dt:81)")
	// cookiestr := `JSESSIONID=4C992C1B138715E5EB262A0FE61CACC2;BIGipServerotn=1357906442.24610.0000;_jc_save_showZtkyts=true;current_captcha_type=C;_jc_save_detail=true;_jc_save_fromStation=%u77F3%u5BB6%u5E84%2CSJP;_jc_save_toStation=%u4E0A%u6D77%2CSHH;_jc_save_fromDate=2015-02-27;_jc_save_toDate=2014-12-22;_jc_save_wfdc_flag=dc`
	// cookie := http.Cookie{Name: "Cookie", Value: cookiestr}
	// request.SetCookie(&cookie)
	request.Param("api_key", "d2V8R7f5mVjS2jNfnv1p4Fvd2lYZwfFl7wtCLWCJ")
	request.Param("pattern", "all")
	bs, err := request.Bytes()
	if log.IsError("{fetch json data error}", err) {
		return nil
	}
	return bs
}

/*根据给定的URL,fetch the data*/
func Do(url, ipaddr string) []byte {
	// httplib.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	setting := httplib.BeegoHttpSettings{}
	setting.TlsClientConfig = &tls.Config{InsecureSkipVerify: true}
	// setting.EnableCookie = true
	httplib.SetDefaultSetting(setting)
	request := httplib.Get(url)
	// request.SetTransport(newTransport(ipaddr))
	request.Header("Host", "kyfw.12306.cn")
	request.Header("Accept", "*/*")
	request.Header("Accept-Language", "zh-CN,zh;q=0.8")
	request.Header("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 SE 2.X MetaSr 1.0")
	request.Header("X-Requested-With", "XMLHttpRequest")
	request.Header("Referer", "https://kyfw.12306.cn/otn/lcxxcx/init")
	request.Header("Connection", "keep-alive")
	// request.Header("Design-Agent", "AppStore/2.0 iOS/7.1.1 model/iPod5,1 build/11D201 (4; dt:81)")
	// request.Header("Cookies/Login", `JSESSIONID=4C992C1B138715E5EB262A0FE61CACC2; BIGipServerotn=1357906442.24610.0000; _jc_save_showZtkyts=true; current_captcha_type=C; _jc_save_detail=true; _jc_save_fromStation=%u77F3%u5BB6%u5E84%2CSJP; _jc_save_toStation=%u4E0A%u6D77%2CSHH; _jc_save_fromDate=2015-02-27; _jc_save_toDate=2014-12-22; _jc_save_wfdc_flag=dc`)
	// cookiestr := `JSESSIONID=4C992C1B138715E5EB262A0FE61CACC2;BIGipServerotn=1357906442.24610.0000;_jc_save_showZtkyts=true;current_captcha_type=C;_jc_save_detail=true;_jc_save_fromStation=%u77F3%u5BB6%u5E84%2CSJP;_jc_save_toStation=%u4E0A%u6D77%2CSHH;_jc_save_fromDate=2015-02-27;_jc_save_toDate=2014-12-22;_jc_save_wfdc_flag=dc`
	// cookie := http.Cookie{Name: "Cookie", Value: cookiestr}
	// request.SetCookie(&cookie)
	bs, err := request.Bytes()
	if log.IsError("{fetch json data error}", err) {
		return nil
	}
	return bs
}

/*
* 固定IP
 */
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
