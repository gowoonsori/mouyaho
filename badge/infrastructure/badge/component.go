package badge

type color string

type bound struct {
	Width  float64
	Height float64
	X      float64
	Y      float64
}

type rect struct {
	Color color
	Bound bound
}

type text struct {
	Msg   string
	Color color
	Bound bound
}

type icon struct {
	Color color
	Bound bound
}

type textBadge struct {
	Rect rect
	Text text
}

type iconBadge struct {
	Rect rect
	Icon icon
}

type likeBadge struct {
	FontFamily     string
	FontSize       int
	React          iconBadge // react icon
	Count          textBadge // count text
	Share          iconBadge // share icon
	Width          float64
	Height         float64
	XRadius        int
	YRadius        int
	ReactClassName string
	Opacity        int
}
