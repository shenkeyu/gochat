package kehuduan

import "fmt"

////////////////////////////////////////////////////////
//
//错误检查
//
////////////////////////////////////////////////////////
func checkError(err error, info string) (res bool) {

	if err != nil {
		fmt.Println(info + err.Error())
		return false
	}
	return true
}
