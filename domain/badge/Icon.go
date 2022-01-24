package badge

type LeftIcon struct {
	Color      string
	Width      float64
	Height     float64
	ReactColor string
	Text       Text
}

type RightIcon struct {
	Color      string
	Width      float64
	Height     float64
	ReactColor string
}

type Text struct {
	TextColor string
	TextValue int
	FontType  string
	FontSize  float64
}
