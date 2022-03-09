package badge

import (
	"bytes"
	"fmt"
	"html/template"
)

var HeartBadgeWriter Writer

func init() {
	tb, err := template.New("like-badge").Parse(heartTemplate)
	if err != nil {
		panic(fmt.Errorf("[Error] HeartBadge TemplateFile open error: %w", err))
	}
	HeartBadgeWriter = &heartBadgeWriter{heartBadgeTemplate: tb}
}

type Writer interface {
	ParseFile(badge HeartBadge) ([]byte, error)
}

type heartBadgeWriter struct {
	heartBadgeTemplate *template.Template
}

type HeartBadge struct {
	BgColor     string
	BorderColor string
	IconColor   string
	ReactColor  string
	OnReact     bool
	TextColor   string
	Text        string
	ShareColor  string
	Edge        string
}

func NewHeartBadge(bgColor string, borderColor string, iconColor string, reactColor string, onReact bool, textColor string, text string, shareColor string, edge string) *HeartBadge {
	return &HeartBadge{BgColor: bgColor, BorderColor: borderColor, IconColor: iconColor, ReactColor: reactColor, OnReact: onReact, TextColor: textColor, Text: text, ShareColor: shareColor, Edge: edge}
}

func (hbw *heartBadgeWriter) ParseFile(badge HeartBadge) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := hbw.heartBadgeTemplate.Execute(buf, badge); err != nil {
		return nil, fmt.Errorf("[Error] HeartBadge File parsing error: %w", err)
	}
	return buf.Bytes(), nil
}
