// 和底层接口通讯

package das

import (
	"net/http"
	"net/url"
	"github.com/zloathleo/go-httpserver/common/logger"
)

//底层接口通讯软件ip和端口
var dasUrl = ""

func sendData2DDI(dataJson string) (*http.Response,error) {
	resp, err := http.PostForm(dasUrl+"/control",
		url.Values{"content": {dataJson}})
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		logger.Println(err)
		return resp,err
	} else {
		return resp,nil
	}

}
