package main

import (
	"fmt"
	"testing"
)

/*
关于GO的test语法, 去看这个文件:Function1_test.go
*/
func TestJusttest(t *testing.T) {
	fmt.Println("@@@ 11111111")
	t.Log("@@@ 22222")
}
