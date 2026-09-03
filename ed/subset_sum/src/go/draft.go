package main

import "fmt"

func subsetSum(arr []int, n, k, index, sum int) bool {
	if sum == k {
		return true
	}

	if sum > k || index == n {
		return false
	}

	incluir := subsetSum(arr, n, k, index+1, sum+arr[index])
	naoIncluir := subsetSum(arr, n, k, index+1, sum)

	return incluir || naoIncluir
}
func main() {
	var n, k int
	if _, err := fmt.Scan(&n, &k); err != nil {
		return
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	resultado := subsetSum(arr, n, k, 0, 0)

	if resultado == true {
		fmt.Println(resultado)
	} else {
		fmt.Println(resultado)
	}
}
