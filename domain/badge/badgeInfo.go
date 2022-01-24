package badge

type BadgeInfo struct {
	Vertical bool

	LeftIcon  LeftIcon
	RightIcon RightIcon

	Color  string
	Width  float64
	Height float64
	Radius float64
}
