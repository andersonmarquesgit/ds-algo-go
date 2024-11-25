package main

func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v // XOR
	}
	return result
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	println(singleNumber(nums)) // 4
	nums2 := []int{2, 2, 1}
	println(singleNumber(nums2)) // 1
}

/**
Iteração 1:
result ^= 4
result: 0000
     4: 0100
-------------
Resultado: 0100 (em decimal: 4)

Iteração 2:
result ^= 1
result: 0100
     1: 0001
-------------
Resultado: 0101 (em decimal: 5)

Iteração 3:
result ^= 2
result: 0101
     2: 0010
-------------
Resultado: 0111 (em decimal: 7)

Iteração 4:
result ^= 1
result: 0111
     1: 0001
-------------
Resultado: 0110 (em decimal: 6)

Iteração 5:
result ^= 2
result: 0110
     2: 0010
-------------
Resultado: 0100 (em decimal: 4)
*/
