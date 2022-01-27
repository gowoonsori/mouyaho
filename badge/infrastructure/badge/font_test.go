package badge

import (
	"gorm.io/gorm/utils/tests"
	"io/ioutil"
	"testing"
)

func Test_Arial_Font_read(t *testing.T) {
	_, err := ioutil.ReadFile(*arialFontFile)
	if err != nil {
		t.Error(err)
	}
}

func Test_Arial_Font_Measure(t *testing.T) {
	//given
	fd := getArialDrawer()
	expect := defaultTextWidth

	//when
	singleD := fd.measureString("1")
	doubleD1 := fd.measureString("12")
	doubleD2 := fd.measureString("68")
	tripleD1 := fd.measureString("123")
	tripleD2 := fd.measureString("485")

	//then
	tests.AssertEqual(t, singleD, expect)
	tests.AssertEqual(t, doubleD1, expect+6)
	tests.AssertEqual(t, doubleD2, expect+6)
	tests.AssertEqual(t, tripleD1, expect+12)
	tests.AssertEqual(t, tripleD2, expect+12)
}
