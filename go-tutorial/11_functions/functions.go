package main

import "fmt"

func sum(a int,  b int) int {
	return a + b
}

func multipleReturn(a int,  b int) (int, int) {
	return a + b, a - b
}

func sumWithVariadic(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main(){
	result := sum(1, 2)
	fmt.Println("result:", result)

	result1, result2 := multipleReturn(1, 2)
	fmt.Println("result1:", result1)
	fmt.Println("result2:", result2)

	result3 := sumWithVariadic(1, 2, 3, 4, 5)
	fmt.Println("result3:", result3)

	nums := []int{1, 2, 3, 4, 5}
	result4 := sumWithVariadic(nums...)
	fmt.Println("result4:", result4)

	nextInt := intSeq()
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())

	nextInt2 := intSeq()
	fmt.Println("nextInt2:", nextInt2())
}