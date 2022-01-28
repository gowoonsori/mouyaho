package badge

import (
	"bytes"
	"fmt"
	"html/template"
)

var b2i = map[bool]int{false: 0, true: 1}
var isReactClassName = map[bool]string{false: "react_off", true: "react_on"}

const (
	defaultBadgeHeight   = float64(30)
	defaultIconRectWidth = float64(35)
	defaultTxtRectWidth  = float64(10)
	XRadius              = 15
	YRadius              = 15
	defaultTextY         = 18
	defaultTextWidth     = 19
)

type Writer interface {
	RenderBadge(b LikeBadge) ([]byte, error)
}

type LikeBadge struct {
	IsReact       bool
	LikeIconColor string

	CountText      string
	CountTextColor string

	ShareIconColor string

	BackgroundColor string
	IsTransparency  bool
}

type likeBadgeWriter struct {
	likeBadgeTemplate *template.Template
}

func NewLikeBadgeWriter() (Writer, error) {
	tb, err := template.New("like-badge").Parse(likeBadgeTemplate)
	if err != nil {
		return nil, fmt.Errorf("[err] LikeBadgeNewWriter %w", err)
	}

	writer := &likeBadgeWriter{
		likeBadgeTemplate: tb,
	}
	return writer, nil
}

func (bw *likeBadgeWriter) RenderBadge(b LikeBadge) ([]byte, error) {
	drawer := getArialDrawer()
	height := defaultBadgeHeight
	textWidth := drawer.measureString(b.CountText)

	lb := &likeBadge{
		FontFamily: fontFamily,
		FontSize:   fontSize,
		React: iconBadge{
			Rect: rect{Color: color(b.BackgroundColor),
				Bound: bound{
					Width:  defaultIconRectWidth,
					Height: height,
					X:      0,
					Y:      0,
				}},
			Icon: icon{Color: color(b.LikeIconColor),
				Bound: bound{
					Width:  0,
					Height: 0,
					X:      15,
					Y:      7,
				},
			},
		},
		Count: textBadge{
			Rect: rect{Color: color(b.BackgroundColor),
				Bound: bound{
					Width:  defaultTxtRectWidth + textWidth,
					Height: height,
					X:      defaultIconRectWidth,
					Y:      0,
				}},
			Text: text{Msg: b.CountText, Color: color(b.CountTextColor),
				Bound: bound{
					Width:  0,
					Height: 0,
					X:      textWidth/2.0 + defaultIconRectWidth,
					Y:      defaultTextY,
				}},
		},
		Share: iconBadge{
			Rect: rect{Color: color(b.BackgroundColor),
				Bound: bound{
					Width:  defaultIconRectWidth,
					Height: height,
					X:      defaultIconRectWidth + textWidth + defaultTxtRectWidth,
					Y:      0,
				}},
			Icon: icon{Color: color(b.ShareIconColor),
				Bound: bound{
					Width:  0,
					Height: 0,
					X:      9 + textWidth - defaultTextWidth,
					Y:      -11,
				},
			},
		},
		Width:   defaultIconRectWidth + defaultTxtRectWidth + textWidth + defaultIconRectWidth,
		Height:  defaultBadgeHeight,
		Rx:      XRadius,
		Ry:      YRadius,
		IsReact: isReactClassName[b.IsReact],
		Opacity: b2i[!b.IsTransparency],
	}

	buf := &bytes.Buffer{}
	if err := bw.likeBadgeTemplate.Execute(buf, lb); err != nil {
		return nil, fmt.Errorf("[err] RenderLikeBadge %w", err)
	}
	return buf.Bytes(), nil
}
