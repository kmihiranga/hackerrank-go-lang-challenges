package main

func simpleArraySum(arrSum [6]int32) int32{
	
	var total int32 = 0

	for i := 0; i < len(arrSum); i++ {
		total = total + arrSum[i];
	}
	return total
}