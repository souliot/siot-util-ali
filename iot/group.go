package iot

import "strconv"

func (c *Client) CreateDeviceGroup(SuperGroupId string, GroupName string, GroupDesc string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "CreateDeviceGroup")
	if SuperGroupId != "" {
		c.Params.Set("SuperGroupId", SuperGroupId)
	}
	c.Params.Set("GroupName", GroupName)
	c.Params.Set("GroupDesc", GroupDesc)

	res, err = c.GetResponse()
	return
}

func (c *Client) QueryDeviceGroupList(SuperGroupId string, PageSize int, CurrentPage int) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "QueryDeviceGroupList")
	if SuperGroupId != "" {
		c.Params.Set("SuperGroupId", SuperGroupId)
	}
	c.Params.Set("PageSize", strconv.Itoa(PageSize))
	c.Params.Set("CurrentPage", strconv.Itoa(CurrentPage))
	res, err = c.GetResponse()
	return
}

func (c *Client) BatchAddDeviceGroupRelations(GroupId string, Devices []map[string]string) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "BatchAddDeviceGroupRelations")
	c.Params.Set("GroupId", GroupId)

	for i, Device := range Devices {
		c.Params.Set("Device."+strconv.Itoa(i+1)+".ProductKey", Device["ProductKey"])
		c.Params.Set("Device."+strconv.Itoa(i+1)+".DeviceName", Device["DeviceName"])
	}
	res, err = c.GetResponse()
	return
}

func (c *Client) QueryDeviceListByDeviceGroup(GroupId string, PageSize int, CurrentPage int) (res Response, err error) {
	c.InitBaseParams()
	c.Params.Set("Action", "QueryDeviceListByDeviceGroup")
	c.Params.Set("GroupId", GroupId)

	c.Params.Set("PageSize", strconv.Itoa(PageSize))
	c.Params.Set("CurrentPage", strconv.Itoa(CurrentPage))
	res, err = c.GetResponse()
	return
}
