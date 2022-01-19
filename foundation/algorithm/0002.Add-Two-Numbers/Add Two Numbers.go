package foundation

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func addTwoNumbers(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	//carry:进位. 例如9+3后, carry就是1
	value1, value2, carry, resultList := 0, 0, 0, &LinkedList{Value: 0}
	//保留住head头结点最后返回.
	head := resultList
	//注意, 要加carry != 0的条件, 否则最后一位进位就被忽略不统计了. 例如结果应该是3,5,2,6,1. 如果忽略了这个条件, 那么结果就会是3,5,2,6
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			value1 = 0
		} else {
			value1 = l1.Value
			l1 = l1.Next
		}
		if l2 == nil {
			value2 = 0
		} else {
			value2 = l2.Value
			l2 = l2.Next
		}
		resultList.Next = &LinkedList{Value: (value1 + value2 + carry) % 10}
		resultList = resultList.Next
		carry = (value1 + value2 + carry) / 10
	}
	return head.Next
}
