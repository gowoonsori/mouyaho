package badge

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"likeIt/infrastructure/internal/fonts"
	"sync"
)

type FontType int

const (
	fontDPI               = 72
	fontSize              = 11
	extraVeraSansDx       = 13
	arialFontFamily       = "Arial,Sans,Verdana,Helvetica,sans-serif"
	veraFontFamily        = "Vera,Sans,Verdana,Helvetica,sans-serif"
	defaultArialFontWidth = 19
	arialWidthPerWord     = 6
	defaultVeraFontWidth  = 20
	veraWidthPerWord      = 7
)

var (
	arialDrawer = initArialFontDrawer()
	veraDrawer  = initVeraFontDrawer()
)

type fontDrawer interface {
	measureString(string) float64
	getDefaultWidth() float64
	getFontFamily() string
}

type fontInfo struct {
	sync.Mutex
	fontSize     int
	extraDx      int
	fontFamily   string
	defaultWidth float64
	drawer       *font.Drawer
}

func (fd *fontInfo) getFontFamily() string {
	return fd.fontFamily
}

func (fd *fontInfo) getDefaultWidth() float64 {
	return fd.defaultWidth
}

func (fd *fontInfo) measureString(s string) float64 {
	fd.Lock()
	p := fd.drawer.MeasureString(s)
	fd.Unlock()

	// must be more than 0.
	size := fd.fixedToPoint(p)
	if size <= 0 {
		return 0
	}

	// add extra margin.
	return size + float64(fd.extraDx)
}

func (fd *fontInfo) fixedToPoint(p fixed.Int26_6) float64 {
	var result float64

	if p < 0 {
		reverse := -p
		result += float64(reverse>>6) * -1
	} else {
		result += float64(p >> 6)
	}
	return result
}

func getArialDrawer() fontDrawer {
	return arialDrawer
}

func initArialFontDrawer() fontDrawer {
	b := fonts.GetArialFont()

	f, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}

	return &fontInfo{
		fontSize:   fontSize,
		extraDx:    extraVeraSansDx,
		fontFamily: arialFontFamily,
		drawer: &font.Drawer{
			Face: truetype.NewFace(f, &truetype.Options{
				Size:    fontDPI,
				DPI:     fontSize,
				Hinting: font.HintingFull,
			}),
		},
		defaultWidth: defaultArialFontWidth,
	}
}

func getVeraDrawer() fontDrawer {
	return veraDrawer
}

func initVeraFontDrawer() fontDrawer {
	b := fonts.GetVeraFont()

	f, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}

	return &fontInfo{
		fontSize:   fontSize,
		extraDx:    extraVeraSansDx,
		fontFamily: veraFontFamily,
		drawer: &font.Drawer{
			Face: truetype.NewFace(f, &truetype.Options{
				Size:    fontDPI,
				DPI:     fontSize,
				Hinting: font.HintingFull,
			}),
		},
		defaultWidth: defaultVeraFontWidth,
	}
}
