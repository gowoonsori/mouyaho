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
	defaultLeftColor     = "red"
	defaultTextColor     = "black"
	defaultRightColor    = "black"
	defaultBg            = "#eee"
	defaultText          = "0"
	xRadius              = "15"
	yRadius              = "15"

	defaultTextY     = 18
	defaultTextWidth = 19
)

//  badge domain
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

// Opt라는 dto를 통해 badge 생성
func CreateBadgeFromOpts(bo *_BadgeOpts) *reactBadge {
	if bo == nil {
		bo = NewBadgeOpts()
	}
	return &reactBadge{LeftIconColor: bo.leftIconColor, Text: bo.text, TextColor: bo.textColor, RightIconColor: bo.rightIconColor, BackgroundColor: bo.backgroundColor, XRadius: bo.xRadius, YRadius: bo.yRadius, IsClear: bo.isClear, IsReact: bo.isReact}
}

//-----------------------------------------------------
//Badge를 만들기 위한 dto
type _BadgeOpts struct {
	leftIconColor   string
	text            string
	textColor       string
	rightIconColor  string
	backgroundColor string
	xRadius         string
	yRadius         string
	isClear         bool
	isReact         bool
}

func NewBadgeOpts() *_BadgeOpts {
	return &_BadgeOpts{
		leftIconColor:   defaultLeftColor,
		text:            defaultText,
		textColor:       defaultTextColor,
		rightIconColor:  defaultRightColor,
		backgroundColor: defaultBg,
		xRadius:         xRadius,
		yRadius:         yRadius,
		isClear:         false,
		isReact:         false,
	}
}

func (bo *_BadgeOpts) LeftIconColor(leftIconColor string) *_BadgeOpts {
	bo.leftIconColor = leftIconColor
	return bo
}

func (bo *_BadgeOpts) Text(text string) *_BadgeOpts {
	bo.text = text
	return bo
}

func (bo *_BadgeOpts) TextColor(textColor string) *_BadgeOpts {
	bo.textColor = textColor
	return bo
}

func (bo *_BadgeOpts) RightIconColor(rightIconColor string) *_BadgeOpts {
	bo.rightIconColor = rightIconColor
	return bo
}

func (bo *_BadgeOpts) BackgroundColor(backgroundColor string) *_BadgeOpts {
	bo.backgroundColor = backgroundColor
	return bo
}

func (bo *_BadgeOpts) XRadius(xRadius string) *_BadgeOpts {
	bo.xRadius = xRadius
	return bo
}

func (bo *_BadgeOpts) YRadius(yRadius string) *_BadgeOpts {
	bo.yRadius = yRadius
	return bo
}

func (bo *_BadgeOpts) IsClear(isClear bool) *_BadgeOpts {
	bo.isClear = isClear
	return bo
}

func (bo *_BadgeOpts) IsReact(isReact bool) *_BadgeOpts {
	bo.isReact = isReact
	return bo
}
