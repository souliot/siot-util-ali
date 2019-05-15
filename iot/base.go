package iot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

var (
	HTTPMethod string = "GET"
	SEPARATOR  string = "&"
	Endpoint   string = "iot.cn-shanghai.aliyuncs.com"
)

type Response struct {
	RequestId string      `json:"RequestId"`
	Success   bool        `json:"Success"`
	Message   string      `json:"Message"`
	Code      string      `json:"Code"`
	PageCount int         `json:"PageCount"`
	PageSize  int         `json:"PageSize"`
	Page      int         `json:"Page"`
	Total     int         `json:"Total"`
	Data      interface{} `json:"Data"`
}

type Client struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Params          url.Values
	Request         *http.Request
}

func NewClient(AccessKeyId string, AccessKeySecret string) (c *Client) {
	c = &Client{}
	c.Endpoint = Endpoint
	c.AccessKeyId = AccessKeyId
	c.AccessKeySecret = AccessKeySecret
	c.Params = url.Values{}
	c.Params.Set("Format", "JSON")
	c.Params.Set("Version", "2018-01-20")
	c.Params.Set("SignatureMethod", "HMAC-SHA1")
	c.Params.Set("SignatureVersion", "1.0")
	c.Params.Set("RegionId", "cn-shanghai")
	c.Params.Set("AccessKeyId", AccessKeyId)

	c.GetRequest()
	return c
}

func (c *Client) SetFormat(format string) {
	c.Params.Set("Format", format)
}

func (c *Client) SetVersion(format string) {
	c.Params.Set("Format", format)
}

func (c *Client) GetRequest() {
	uri := "https://" + c.Endpoint
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		return
	}

	mime := "application/json"
	header := map[string][]string{
		"Content-Type": {mime},
	}
	req.Header = header
	c.Request = req
}

func (c *Client) GetResponse() (res Response, err error) {
	res = Response{}
	c.Signature()
	c.Request.URL.RawQuery = c.Params.Encode()
	client := &http.Client{}
	resp, err := client.Do(c.Request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	beego.Info(string(body))

	err = json.Unmarshal(body, &res)
	return
}

func (c *Client) GenerateSignString() string {
	var buffer bytes.Buffer
	buffer.WriteString(HTTPMethod)
	buffer.WriteString(SEPARATOR)
	t := percentEncode("/")
	buffer.WriteString(t)
	buffer.WriteString(SEPARATOR)
	t = percentEncode(c.Params.Encode())
	buffer.WriteString(t)

	return buffer.String()
}

func (c *Client) Signature() {
	c.Params.Del("Signature")
	c.Params.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	c.Params.Set("SignatureNonce", strconv.Itoa(int(time.Now().Unix())))
	key := c.AccessKeySecret + "&"
	h := hmac.New(sha1.New, []byte(key))
	io.WriteString(h, c.GenerateSignString())

	signature := string(base64Encode(h.Sum(nil)))
	c.Params.Set("Signature", signature)
}

func percentEncode(d string) (str string) {
	str = url.QueryEscape(d)
	rep := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	str = rep.Replace(str)
	return
}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}
