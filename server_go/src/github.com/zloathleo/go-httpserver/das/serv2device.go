// 和底层接口通讯

package das

import (
	"log"
	"net/http"
	"net/url"
)

//底层接口通讯软件ip和端口
const DDIIp = "localhost"
const DDIPort = "8080"

func sendData2DDI(dataJson string) error {
	resp, err := http.PostForm("http://"+DDIIp+":"+DDIPort+"/control",
		url.Values{"content": {dataJson}})
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		log.Print(err)
		return err
	} else {
		return nil
	}

}
