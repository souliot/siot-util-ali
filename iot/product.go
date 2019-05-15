package iot

type Product struct {
}
type ProductTag struct {
}

// Product
func (c *Client) CreateProduct() (res Response, err error) {
	return
}
func (c *Client) UpdateProduct() (res Response, err error) {
	return
}
func (c *Client) QueryProductList() (res Response, err error) {
	c.Params.Set("Action", "QueryProductList")
	c.Params.Set("CurrentPage", "1")
	c.Params.Set("PageSize", "10")
	res, err = c.GetResponse()

	return
}
func (c *Client) QueryProduct() (res Response, err error) {
	return
}
func (c *Client) DeleteProduct() (res Response, err error) {
	return
}
func (c *Client) ListProductByTags() (res Response, err error) {
	return
}

// Product Tags
func (c *Client) CreateProductTags() (res Response, err error) {
	return
}
func (c *Client) UpdateProductTags() (res Response, err error) {
	return
}
func (c *Client) DeleteProductTags() (res Response, err error) {
	return
}
func (c *Client) ListProductTags() (res Response, err error) {
	return
}
