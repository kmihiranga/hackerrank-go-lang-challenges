package main


func veryBigSum(arr []int64) int64{

	var total int64 = 0
	for i, _ := range arr {
		total += arr[i]
	}

	return total
}