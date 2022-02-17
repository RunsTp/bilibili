package bilibili

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	cookies       []*http.Cookie
	cookiesString string
	timeout       time.Duration
	logger        resty.Logger
}

var std = &Client{}

// SetTimeout 设置http请求超时时间
func SetTimeout(timeout time.Duration) {
	std.SetTimeout(timeout)
}
func (c *Client) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// GetTimeout 获取http请求超时时间，默认20秒
func GetTimeout() time.Duration {
	return std.GetTimeout()
}
func (c *Client) GetTimeout() time.Duration {
	if c.timeout == 0 {
		return time.Second * 20
	}
	return c.timeout
}

// SetLogger 设置logger
func SetLogger(logger resty.Logger) {
	std.SetLogger(logger)
}
func (c *Client) SetLogger(logger resty.Logger) {
	c.logger = logger
}

// GetLogger 获取logger，默认使用resty默认的logger
func GetLogger() resty.Logger {
	return std.GetLogger()
}
func (c *Client) GetLogger() resty.Logger {
	return c.logger
}

// GetCookiesString 获取字符串格式的cookies，方便自行存储后下次使用。只有正常登录或者调用 SetCookiesString 之后，
// 这个函数才会有返回值。如果调用的是 SetCookies 或者没有进行过正常登录，则返回空字符串。
//
// 配合下面的 SetCookiesString 使用。
func GetCookiesString() string {
	return std.cookiesString
}
func (c *Client) GetCookiesString() string {
	return c.cookiesString
}

// SetCookiesString 设置Cookies，但是是字符串格式，配合 GetCookiesString 使用。有些功能必须登录或设置Cookies后才能使用。
func SetCookiesString(cookiesString string) {
	std.SetCookiesString(cookiesString)
}
func (c *Client) SetCookiesString(cookiesString string) {
	c.cookiesString = cookiesString
	c.cookies = (&resty.Response{RawResponse: &http.Response{Header: http.Header{
		"Set-Cookie": strings.Split(cookiesString, "\n"),
	}}}).Cookies()
}

// GetCookies 获取Cookies。配合下面的SetCookies使用。
func GetCookies() []*http.Cookie {
	return std.GetCookies()
}
func (c *Client) GetCookies() []*http.Cookie {
	return c.cookies
}

// SetCookies 设置Cookies。有些功能必须登录之后才能使用，设置Cookies可以代替登录。
func SetCookies(cookies []*http.Cookie) {
	std.SetCookies(cookies)
}
func (c *Client) SetCookies(cookies []*http.Cookie) {
	c.cookies = cookies
	c.cookiesString = ""
}

// 获取resty的一个request
func (c *Client) resty() *resty.Client {
	client := resty.New().SetTimeout(c.GetTimeout())
	if c.logger != nil {
		client.SetLogger(c.logger)
	}
	return client
}
