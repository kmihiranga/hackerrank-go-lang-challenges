package main

import "fmt"

func main() {
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// m := map[string]int{ 
	// 	"James": 32,
	// 	"Miss MoneyPenny": 27,
	// }
	// fmt.Println(m)

	// v, ok := m["Kalana"]

	// fmt.Println(v)
	// fmt.Println(ok)

	// if v, ok := m["James"]; ok {
	// 	fmt.Println("This is the if ", v)
	// }

	/* Simple Array Sum */
	// arr := [6]int32{1, 2, 3, 4, 10, 11}
	// sum := simpleArraySum(arr)
	// fmt.Println(sum)

	// array triplet
	// a := []int32{17, 28, 30}
	// b := []int32{99, 16, 8}

	// very big sum

	// arr := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	// sum := veryBigSum(arr);


	// resultArr := compareTriplet(a, b)
	sqaure := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	multiArr := diagonalDifferences(sqaure)

	fmt.Println(multiArr)
}


