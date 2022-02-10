package application

import (
	"net/url"
	"strconv"
)

type UrlInfo struct {
	Url             string
	LikeIconColor   string
	CountTextColor  string
	ShareIconColor  string
	BackgroundColor string
	IsClear         bool
}

func CreateUrlInfoFromUrl(reqUrl string) *UrlInfo {
	u := &UrlInfo{}
	if err := u.parsingQueryString(reqUrl); err != nil {
		return nil
	}

	return u
}

func (u *UrlInfo) parsingQueryString(reqUrl string) error {
	//url 생성
	p, _ := url.Parse(reqUrl)
	//query string decode
	rq, _ := url.QueryUnescape(p.RawQuery)
	//query string to map
	m, err := url.ParseQuery(rq)
	if err != nil {
		return err
	}

	if v, ok := m["url"]; ok {
		u.Url = v[0]
	}
	if v, ok := m["like_color"]; ok {
		u.LikeIconColor = v[0]
	}
	if v, ok := m["text_color"]; ok {
		u.CountTextColor = v[0]
	}
	if v, ok := m["share_color"]; ok {
		u.ShareIconColor = v[0]
	}
	if v, ok := m["bg"]; ok {
		u.BackgroundColor = v[0]
	}
	if v, ok := m["clear"]; ok {
		if b, err2 := strconv.ParseBool(v[0]); err2 == nil {
			u.IsClear = b
		}
	}

	return nil
}
