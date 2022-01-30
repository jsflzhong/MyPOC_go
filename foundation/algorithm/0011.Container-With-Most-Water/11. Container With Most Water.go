package leetcode

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		// 例: [1,8,6,2,5,4,8,3,7]
		if height[start] < height[end] {
			//取容器能盛水的高度, 应该是最低的那个值
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}

// 用对撞指针
func my_maxArea(inputArray []int) int {
	length := len(inputArray)
	if length < 2 {
		return -1
	}
	leftP, rightP, containerHeight, volume := 0, length-1, 0, 0
	//1.对撞指针,找出最大的长宽乘积, 即最大的容器的容积.
	for leftP < rightP {
		tempVolume, width := 0, rightP-leftP
		//找出最大的短板
		if inputArray[leftP] < inputArray[rightP] {
			//容器能盛水的多少,取决于短板
			containerHeight = inputArray[leftP]
			//往中间移动指针, 看有没有比短板高一些的板
			leftP++
		} else {
			containerHeight = inputArray[rightP]
			rightP--
		}

		//反例: 不要像这样在这里计算宽度, 因为上面已经把指针++或--了,如果在这里减出来的宽度是不准的! 要在上面移动指针之前减出来!
		//tempVolume = (rightP - leftP) * containerHeight

		tempVolume = width * containerHeight

		if tempVolume > volume {
			volume = tempVolume
		}
	}
	return volume
}
