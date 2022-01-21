package badge

const (
	defaultBadgeHeight = float64(20)
	defaultIconWidth   = float64(15)
	defaultIconHeight  = float64(15)
	defaultIconX       = float64(3)
	defaultIconY       = float64(2.5)
)

type BadgeInfo struct {
	FontType FontType

	LeftText            string
	LeftTextColor       string
	LeftBackgroundColor string

	RightText            string
	RightTextColor       string
	RightBackgroundColor string

	XRadius string
	YRadius string
}
