package iot

import "strconv"

func (c *Client) InvokeDataAPIService(apiSrn string, params map[string]string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "InvokeDataAPIService")
	c.Params.Set("ApiSrn", apiSrn)
	for k, v := range params {
		i := 0
		c.Params.Set("Param."+strconv.Itoa(i+1)+".ParamName", k)
		c.Params.Set("Param."+strconv.Itoa(i+1)+".ParamValue", v)
		i++
	}
	res, err = c.GetResponse()
	return
}

func (c *Client) GetDataAPIServiceDetail(apiSrn string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "GetDataAPIServiceDetail")
	c.Params.Set("ApiSrn", apiSrn)
	res, err = c.GetResponse()
	return
}
