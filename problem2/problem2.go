package problem2

func BlowOutCandles(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var max, sum int
	max = arr[0]
	for index := 0; index < len(arr); index++ {
		if arr[index] > max {
			max = arr[index]
			sum = 1
		} else if arr[index] == max {
			sum++
		}
	}

	return sum
}

func CountCandles(arr []int) map[int]int {
	countMap := make(map[int]int, 0)

	for index := 0;	index < len(arr) ; index++ {
		_, ok := countMap[arr[index]]
		if ok {
			countMap[arr[index]]++
		}else {
			countMap[arr[index]] = 1
		}


	}


	return countMap
}