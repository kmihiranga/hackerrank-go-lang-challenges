package main

func diagonalDifferences(arr [][]int32) int32{

	var diag1 int32 = 0
	var diag2 int32 = 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			// check first diagonal
			if(i == j) {
				diag1 += arr[i][j]
			}
			// check second diagonal
			if(i + j == len(arr) -1) {
				diag2 += arr[i][j]
			}
		}
	}
	return Abs(diag1 - diag2)

}

func Abs(x int32) int32{
	if(x < 0) {
		return -x
	}
	return x
}