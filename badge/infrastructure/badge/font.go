package badge

import (
	"flag"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"io/ioutil"
	"sync"
)

type FontType int

const (
	fontDPI         = 72
	fontSize        = 11
	extraVeraSansDx = 13
	fontFamily      = "Arial,Sans,Verdana,Helvetica,sans-serif"
	fontsPath       = "../fonts/"
)

var (
	arialFontFile = flag.String("fontfile", fontsPath+"ARIAL.TTF", "filename of the ttf font")
	arialDrawer   = initArialFontDrawer()
)

type fontDrawer interface {
	measureString(string) float64
}

type fontInfo struct {
	sync.Mutex
	fontSize   int
	extraDx    int
	fontFamily string
	drawer     *font.Drawer
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

	// 26 bit integer(with 1 sign)
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
	b, err := ioutil.ReadFile(*arialFontFile)
	if err != nil {
		panic(err)
	}

	f, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}

	return &fontInfo{
		fontSize:   fontSize,
		extraDx:    extraVeraSansDx,
		fontFamily: fontFamily,
		drawer: &font.Drawer{
			Face: truetype.NewFace(f, &truetype.Options{
				Size:    fontDPI,
				DPI:     fontSize,
				Hinting: font.HintingFull,
			}),
		},
	}
}
