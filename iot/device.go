package iot

import "strconv"

// Device
func (c *Client) QueryDevice(ProductKey string) (res Response, err error) {
	c.Params.Set("Action", "QueryDevice")
	c.Params.Set("ProductKey	", ProductKey)
	c.Params.Set("CurrentPage", "1")
	c.Params.Set("PageSize", "50")
	res, err = c.GetResponse()
	return
}

func (c *Client) QueryDeviceDetail(IotId string) (res Response, err error) {
	c.Params.Set("Action", "QueryDeviceDetail")
	c.Params.Set("IotId	", IotId)
	res, err = c.GetResponse()
	return
}

func (c *Client) QueryDeviceProp(IotId string) (res Response, err error) {
	c.Params.Set("Action", "QueryDeviceProp")
	c.Params.Set("IotId	", IotId)
	res, err = c.GetResponse()
	return
}

func (c *Client) QueryDevicePropertyStatus(IotId string) (res Response, err error) {
	c.Params.Set("Action", "QueryDevicePropertyStatus")
	c.Params.Set("IotId	", IotId)
	res, err = c.GetResponse()
	return
}

// Identifier 设备的属性Identifier
// StartTime、EndTime 要查询的属性记录的开始时间、结束时间。取值为毫秒值时间戳。
// PageSize 返回结果中每页显示的记录数。数量限制：每页最多可显示50条。
// Asc 返回结果中属性记录的排序方式，取值：0：倒序；1：正序。
func (c *Client) QueryDevicePropertyData(IotId string, Identifier string, StartTime int64, EndTime int64, PageSize int, Asc int) (res Response, err error) {
	c.Params.Set("Action", "QueryDevicePropertyData")
	c.Params.Set("IotId", IotId)
	c.Params.Set("Identifier", Identifier)
	c.Params.Set("StartTime", strconv.Itoa(int(StartTime)))
	c.Params.Set("EndTime", strconv.Itoa(int(EndTime)))
	c.Params.Set("PageSize", strconv.Itoa(PageSize))
	c.Params.Set("Asc", strconv.Itoa(Asc))
	res, err = c.GetResponse()
	return
}
