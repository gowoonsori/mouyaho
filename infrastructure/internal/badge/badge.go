package badge

import (
	"bytes"
	"fmt"
	"html/template"
)

var (
	b2i              = map[bool]int{false: 0, true: 1}
	isReactClassName = map[bool]string{false: "unlike", true: "like"}
)

const (
	defaultBadgeHeight   = float64(30)
	defaultIconRectWidth = float64(35)
	defaultTxtRectWidth  = float64(10)

	defaultTextY     = 18
	defaultTextWidth = 19
)

type reactBadge struct {
	LeftIconColor string

	Text      string
	TextColor string

	RightIconColor string

	BackgroundColor string

	XRadius string
	YRadius string

	IsClear bool
	IsReact bool
}

func NewBadge(leftIconColor string, text string, textColor string, rightIconColor string, backgroundColor string, XRadius string, YRadius string, isClear bool, isReact bool) *reactBadge {
	return &reactBadge{LeftIconColor: leftIconColor, Text: text, TextColor: textColor, RightIconColor: rightIconColor, BackgroundColor: backgroundColor, XRadius: XRadius, YRadius: YRadius, IsClear: isClear, IsReact: isReact}
}

type Writer interface {
	RenderBadgeFile(b reactBadge) ([]byte, error)
}

type likeBadgeWriter struct {
	likeBadgeTemplate *template.Template
}

func NewLikeBadgeWriter() (Writer, error) {
	tb, err := template.New("like-badge").Parse(likeBadgeTemplate)
	if err != nil {
		return nil, fmt.Errorf("[err] LikeBadgeNewWriter %w", err)
	}

	return &likeBadgeWriter{likeBadgeTemplate: tb}, nil
}

func (lbw *likeBadgeWriter) RenderBadgeFile(b reactBadge) ([]byte, error) {
	drawer := getArialDrawer()
	height := defaultBadgeHeight
	textWidth := drawer.measureString(b.Text)

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
			Icon: icon{Color: color(b.LeftIconColor),
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
			Text: text{Msg: b.Text, Color: color(b.TextColor),
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
			Icon: icon{Color: color(b.RightIconColor),
				Bound: bound{
					Width:  0,
					Height: 0,
					X:      9 + textWidth - defaultTextWidth,
					Y:      -11,
				},
			},
		},
		Width:          defaultIconRectWidth + defaultTxtRectWidth + textWidth + defaultIconRectWidth,
		Height:         defaultBadgeHeight,
		XRadius:        b.XRadius,
		YRadius:        b.YRadius,
		ReactClassName: isReactClassName[b.IsReact],
		Opacity:        b2i[!b.IsClear],
	}

	buf := &bytes.Buffer{}
	if err := lbw.likeBadgeTemplate.Execute(buf, lb); err != nil {
		return nil, fmt.Errorf("[err] RenderLikeBadge %w", err)
	}
	return buf.Bytes(), nil
}
