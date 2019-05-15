package main

import (
	"net/url"

	sutil "github.com/souliot/siot-util"
)

func main() {
	uri := "http://192.168.50.200:8082/v1/video"

	// GET
	var param = url.Values{}
	param.Add("page", "1")
	param.Add("pageSize", "2")
	data := param.Encode()

	// POST
	// m := map[string]interface{}{
	// 	"page":     1,
	// 	"pageSize": 2,
	// }

	// mjson, _ := json.Marshal(m)
	// data := string(mjson)

	mime := "application/json"

	header := map[string][]string{
		"X-Access-Token": {"UlRZ2SSToJ0fcaoeaJSA8g=="},
		"Content-Type":   {mime},
	}
	sutil.HttpDo("GET", uri, data, header)
	// beego.Info(str)
}
