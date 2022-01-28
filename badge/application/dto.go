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
	IsTransparency  bool
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
	if v, ok := queryStrings["transparency"]; ok {
		if b, err := strconv.ParseBool(v); err == nil {
			result.IsTransparency = b
		}
	}
	return result
}

func (u *UrlInfo) CreateBadgeUrl() string {
	var b bytes.Buffer
	b.WriteString(u.Url)
	b.WriteString("?")
	b.WriteString(u.CountTextColor)
	b.WriteString(u.BackgroundColor)
	b.WriteString(u.LikeIconColor)
	b.WriteString(u.ShareIconColor)
	b.WriteString(strconv.FormatBool(u.IsTransparency))
	return b.String()
}
