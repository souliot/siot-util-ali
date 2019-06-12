package iot

import "strconv"

// Device
func (c *Client) QueryDevice(ProductKey string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "QueryDevice")
	c.Params.Set("ProductKey	", ProductKey)
	c.Params.Set("CurrentPage", "1")
	c.Params.Set("PageSize", "50")
	res, err = c.GetResponse()
	return
}

func (c *Client) QueryDeviceDetail(IotId string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "QueryDeviceDetail")
	c.Params.Set("IotId	", IotId)
	res, err = c.GetResponse()
	return
}

func (c *Client) QueryDeviceProp(IotId string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "QueryDeviceProp")
	c.Params.Set("IotId	", IotId)
	res, err = c.GetResponse()
	return
}

func (c *Client) QueryDevicePropertyStatus(IotId string) (res Response, err error) {
	c.InitBaseParams()
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
	c.InitBaseParams()
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

func (c *Client) QueryDeviceByTags(PageSize int, CurrentPage int, Tag map[string]string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "QueryDeviceByTags")
	c.Params.Set("PageSize", strconv.Itoa(PageSize))
	c.Params.Set("CurrentPage", strconv.Itoa(CurrentPage))

	for k, v := range Tag {
		i := 0
		c.Params.Set("Tag."+strconv.Itoa(i+1)+".TagKey", k)
		c.Params.Set("Tag."+strconv.Itoa(i+1)+".TagValue", v)
		i++
	}
	res, err = c.GetResponse()
	return
}

func (c *Client) SaveDeviceProp(ProductKey string, DeviceName string, Props string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "SaveDeviceProp")
	c.Params.Set("ProductKey", ProductKey)
	c.Params.Set("DeviceName", DeviceName)
	c.Params.Set("Props", Props)

	res, err = c.GetResponse()
	return
}

func (c *Client) BatchUpdateDeviceNickname(DeviceNicknameInfos []map[string]string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "BatchUpdateDeviceNickname")
	for i, DeviceNicknameInfo := range DeviceNicknameInfos {
		c.Params.Set("DeviceNicknameInfo."+strconv.Itoa(i+1)+".ProductKey", DeviceNicknameInfo["ProductKey"])
		c.Params.Set("DeviceNicknameInfo."+strconv.Itoa(i+1)+".DeviceName", DeviceNicknameInfo["DeviceName"])
		c.Params.Set("DeviceNicknameInfo."+strconv.Itoa(i+1)+".Nickname", DeviceNicknameInfo["Nickname"])
	}
	res, err = c.GetResponse()
	return
}
