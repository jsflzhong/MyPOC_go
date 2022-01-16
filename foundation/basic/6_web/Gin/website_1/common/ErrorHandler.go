package common

import "fmt"

func CheckError(err error) bool{
	if err != nil {
		fmt.Println("@@@[CheckError]error:",err)
		return true
	}
	return false
}
