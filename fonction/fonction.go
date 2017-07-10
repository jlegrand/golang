package fonction

func max(num1, num2 int) (result int) {

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return
}

func Max(num1, num2 int) int {
	return max(num1, num2)
}
