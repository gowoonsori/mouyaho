package fonts

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func Test_Read_CURRENT_DIR(t *testing.T) {
	_, b, _, _ := runtime.Caller(0) //해당 코드의 위치
	fmt.Println(filepath.Dir(b))

	dir, _ := os.Getwd() //현재 실행한 위치
	fmt.Println(dir)

	dir, _ = filepath.Abs("./") //현재 실행한 위치
	fmt.Println(dir)

	dir, _ = filepath.Abs(filepath.Dir(os.Args[0])) //빌드된 파일의 위치(tmp/)
	fmt.Println(dir)

	dir, _ = os.Executable() //빌드된 파일의 위치(tmp/)
	fmt.Println(filepath.Dir(dir))
}
