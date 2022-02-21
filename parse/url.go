/*
 * @Author: reber
 * @Mail: reber0ask@qq.com
 * @Date: 2022-02-21 16:51:19
 * @LastEditTime: 2022-02-21 16:51:34
 */
package parse

import (
	"fmt"
	"net"
	"net/url"
	"path"
	"strings"
)

type ParseUrl struct {
	u *url.URL
}

// 解析 URL
func NewParseURL(targetURL string) *ParseUrl {
	urlObj, _ := url.Parse(targetURL)

	return &ParseUrl{
		u: urlObj,
	}
}

// 获取 BaseURL
func (p *ParseUrl) BaseURL() string {
	return fmt.Sprintf("%s://%s/", p.u.Scheme, p.u.Host)
}

// 获取 Scheme
func (p *ParseUrl) Scheme() string {
	return p.u.Scheme
}

// 获取 Username
func (p *ParseUrl) Username() string {
	return p.u.User.Username()
}

// 获取 Password
func (p *ParseUrl) Password() string {
	Pwd, _ := p.u.User.Password()
	return Pwd
}

// 获取 Host
func (p *ParseUrl) Host() string {
	Host, _, _ := net.SplitHostPort(p.u.Host)
	return Host
}

// 获取 Port
func (p *ParseUrl) Port() string {
	_, Port, _ := net.SplitHostPort(p.u.Host)
	return Port
}

// 获取 Path
func (p *ParseUrl) Path() string {
	return p.u.Path
}

// 获取 SuffixName
func (p *ParseUrl) SuffixName() string {
	fileType := path.Ext(p.u.Path)
	ext := strings.TrimLeft(fileType, ".")

	return ext
}

// 获取 RawQuery
func (p *ParseUrl) RawQuery() string {
	return p.u.RawQuery
}

// 获取 MapQuery
func (p *ParseUrl) MapQuery() url.Values {
	MapQuery, _ := url.ParseQuery(p.u.RawQuery)
	return MapQuery
}

// 获取 Fragment
func (p *ParseUrl) Fragment() string {
	return p.u.Fragment
}
