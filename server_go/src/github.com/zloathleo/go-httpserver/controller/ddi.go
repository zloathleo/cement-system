/**
与接口程序交互
*/
package controller

import (
	"github.com/zloathleo/go-httpserver/common/logger"
	"net/http"
	"net/url"
	"time"
)

//命令下发到接口程序
func cmdPostDDI(dataJson string) (*http.Response, error) {
	//http request to ddi

	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := httpClient.PostForm(cmdDasTargetUrl+"/control", url.Values{"content": {dataJson}})
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		logger.Infof("命令下发异常,[%v].", err)
		return resp, err
	} else {
		return resp, nil
	}
}
