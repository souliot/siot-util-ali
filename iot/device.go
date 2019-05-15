package iot

type Device struct {
}
type DeviceProp struct {
}

// Product
func (c *Client) QueryDevice(ProductKey string) (res Response, err error) {
	c.Params.Set("Action", "QueryDevice")
	c.Params.Set("ProductKey	", ProductKey)
	c.Params.Set("CurrentPage", "1")
	c.Params.Set("PageSize", "50")
	res, err = c.GetResponse()
	return
}
