package application

import (
	"bytes"
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

func CreateUrlInfoFromMap(queryStrings map[string]string) *UrlInfo {
	result := &UrlInfo{}

	if v, ok := queryStrings["url"]; ok {
		result.Url = v
	}
	if v, ok := queryStrings["like_color"]; ok {
		result.LikeIconColor = v
	}
	if v, ok := queryStrings["text_color"]; ok {
		result.CountTextColor = v
	}
	if v, ok := queryStrings["share_color"]; ok {
		result.ShareIconColor = v
	}
	if v, ok := queryStrings["bg"]; ok {
		result.BackgroundColor = v
	}
	if v, ok := queryStrings["clear"]; ok {
		if b, err := strconv.ParseBool(v); err == nil {
			result.IsClear = b
		}
	}
	return result
}

func (u *UrlInfo) CreateBadgeUrl() string {
	var b bytes.Buffer
	b.WriteString(u.Url)

	return b.String()
}
