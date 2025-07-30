package main

import "fmt"

/*
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

func jumpingOnClouds(c []int32) int32 {
	if len(c) <= 1 {
		return 0
	}

	if len(c) == 2 {
		if c[1] == 0 {
			return 1
		}
		return 0
	}

	jumps := 0
	i := 0
	clouds := len(c)

	for i < clouds-1 {
		if i+2 < clouds && c[i+2] == 0 {
			i += 2
		} else if i+1 < clouds && c[i+1] == 0 {
			i++
		}
		jumps++
	}

	return int32(jumps)
}

func main() {
	c := []int32{0, 0, 1, 0, 0, 1, 0}
	fmt.Println(jumpingOnClouds(c))
}
