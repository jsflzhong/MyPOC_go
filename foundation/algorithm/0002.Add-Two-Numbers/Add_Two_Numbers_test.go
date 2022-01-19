package foundation

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &LinkedList{1, &LinkedList{2, &LinkedList{8, &LinkedList{9, nil}}}}
	l2 := &LinkedList{2, &LinkedList{3, &LinkedList{4, &LinkedList{6, nil}}}}
	fmt.Println("@@@The result should be: 3, 5, 2, 6, 1")
	resultList := addTwoNumbers(l1, l2)
	for resultList != nil {
		fmt.Printf("@@@The result number is :[%d] \n", resultList.Value)
		resultList = resultList.Next
	}

	l3 := &LinkedList{9, &LinkedList{9, &LinkedList{9, &LinkedList{9, nil}}}}
	l4 := &LinkedList{Value: 1}
	fmt.Println("@@@The result should be: 0,0,0,0,1")
	resultList2 := addTwoNumbers(l3, l4)
	for resultList2 != nil {
		fmt.Printf("@@@The result number is :[%d] \n", resultList2.Value)
		resultList2 = resultList2.Next
	}
}
