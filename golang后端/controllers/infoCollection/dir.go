package infoCollection

import "fmt"

func DirScanNew(target string) string{
	DirScanAll := [...]string{"http://www.baidu.com/root","http://asdfhasdf/a","http://sadfhasd/1111111"}
	return fmt.Sprintf("%v",DirScanAll)
}