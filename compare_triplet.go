package main

func compareTriplet(a []int32, b []int32) []int32 {
	
	var alice int32 = 0
	var bob int32 = 0

	result := []int32{0, 0}

	for i := 0; i <= len(result); i++ {
		if a[i] > b[i] {
			alice += 1
		}else if a[i] < b[i] {
			bob += 1
		}else {
			alice += 0
			bob += 0
		}
	}

	result = []int32{alice, bob}
	return result	
}	