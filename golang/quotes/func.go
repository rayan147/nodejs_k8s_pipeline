package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func ops(a int, b int) (int, int) {
	return a + b, a - b
}

func sumTotal(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
	
// Function as argument
func doit(oparator func(int, int) int, a int, b int) int {
	return oparator(a, b)
}

// Mutiple return values
func multiple(a int, b int)  int {
	return a * b
}

func main() {
	res := sum(1, 2)
	fmt.Println(res)

	res1, res2 := ops(2, 2)
	fmt.Println("2+2=", res1,"2-2=", res2)
	b,_ := ops(1, 2)
	fmt.Println("10+2=", b)
	total := sumTotal(1, 2, 3, 4, 5)
	fmt.Println("total=", total)

	c:= doit(sum, 1, 2)
	fmt.Println("1+2=", c)
	d := doit(multiple, 2, 2)
	fmt.Println("2*2=", d)
}