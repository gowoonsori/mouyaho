package badge

import (
	"gorm.io/gorm/utils/tests"
	"testing"
)

func Test_Arial_Font_read(t *testing.T) {
	getArialDrawer()
}

func Test_Arial_Font_Measure(t *testing.T) {
	//given
	fd := getArialDrawer()
	expect := defaultArialFontWidth

	//when
	singleD := fd.measureString("1")
	doubleD1 := fd.measureString("12")
	doubleD2 := fd.measureString("68")
	tripleD1 := fd.measureString("123")
	tripleD2 := fd.measureString("485")

	//then
	tests.AssertEqual(t, singleD, expect)
	tests.AssertEqual(t, doubleD1, expect+arialWidthPerWord)
	tests.AssertEqual(t, doubleD2, expect+arialWidthPerWord)
	tests.AssertEqual(t, tripleD1, expect+arialWidthPerWord*2)
	tests.AssertEqual(t, tripleD2, expect+arialWidthPerWord*2)
}

func Test_Vera_Font_read(t *testing.T) {
	getVeraDrawer()
}

func Test_Vera_Font_Measure(t *testing.T) {
	//given
	fd := getVeraDrawer()
	expect := defaultVeraFontWidth

	//when
	singleD := fd.measureString("1")
	doubleD1 := fd.measureString("12")
	doubleD2 := fd.measureString("68")
	tripleD1 := fd.measureString("123")
	tripleD2 := fd.measureString("485")

	//then
	tests.AssertEqual(t, singleD, expect)
	tests.AssertEqual(t, doubleD1, expect+veraWidthPerWord)
	tests.AssertEqual(t, doubleD2, expect+veraWidthPerWord)
	tests.AssertEqual(t, tripleD1, expect+veraWidthPerWord*2)
	tests.AssertEqual(t, tripleD2, expect+veraWidthPerWord*2)
}
