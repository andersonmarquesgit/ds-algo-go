package main

import (
	"fmt"
	"math"
)

func findTwoSum(nums []int, target int) []int {
	// O(n^2)
	if len(nums) <= 1 {
		return nil
	}

	for p1 := 0; p1 < len(nums); p1++ { // O(n)
		numberToFind := target - nums[p1]
		for p2 := p1 + 1; p2 < len(nums); p2++ { // O(n)
			if numberToFind == nums[p2] {
				return []int{p1, p2}
			}
		}
	}

	// O(n * n) = O(n^2)

	return nil
}

func findTwoSumOptimized(nums []int, target int) []int {
	// O(n^2)
	if len(nums) <= 1 {
		return nil
	}

	numberToFindMap := make(map[int]int)

	for p1 := 0; p1 < len(nums); p1++ {
		if _, ok := numberToFindMap[nums[p1]]; ok {
			return []int{numberToFindMap[nums[p1]], p1}
		}

		numberToFind := target - nums[p1]
		numberToFindMap[numberToFind] = p1
	}

	return nil
}

func containerMostWater(heights []int) int {
	//O(n^2)
	maxArea := 0

	for a := 0; a < len(heights); a++ {
		for b := a + 1; b < len(heights); b++ {
			height := min(heights[a], heights[b])
			width := b - a
			area := height * width
			//currentArea := int(math.Min(float64(heights[a]), float64(heights[b])) * float64(b-a))

			//if area > maxArea {
			//	maxArea = area
			//}
			maxArea = int(math.Max(float64(area), float64(maxArea)))
		}
	}

	return maxArea
}

// 7, 1, 2, 3, 9
func containerMostWaterOptimized(heights []int) int {
	// O(n)
	maxArea := 0
	a := 0
	b := len(heights) - 1

	for a < b {
		height := min(heights[a], heights[b])
		width := b - a
		area := height * width
		maxArea = int(math.Max(float64(area), float64(maxArea)))

		if heights[a] <= heights[b] {
			a++
		} else {
			b--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func howMuchRainWater(elevations []int) int {
	// O(n^2)
	totalRainwater := 0
	for p := 0; p < len(elevations); p++ {
		maxL := 0
		maxR := 0
		currentHeight := elevations[p]

		//MaxL
		for leftP := p - 1; leftP >= 0; leftP-- {
			if elevations[leftP] > maxL {
				maxL = elevations[leftP]
			}
		}

		//MaxR
		for rightP := p + 1; rightP < len(elevations); rightP++ {
			if elevations[rightP] > maxR {
				maxR = elevations[rightP]
			}
		}

		currentWater := min(maxL, maxR) - currentHeight

		if currentWater >= 0 {
			totalRainwater += currentWater
		}

	}
	return totalRainwater
}

func howMuchRainWaterOptimized(elevations []int) int {
	// O(n)
	totalRainwater := 0
	leftP := 0
	rightP := len(elevations) - 1
	maxLeft := 0
	maxRight := 0

	for leftP < rightP {
		currentWater := 0
		currentElevation := 0

		if elevations[leftP] <= elevations[rightP] { // Esse if basicamente substitui a operação para descobrir a menor elevação
			currentElevation = elevations[leftP]
			if maxLeft > currentElevation {
				currentWater = maxLeft - currentElevation
				totalRainwater += currentWater
			} else {
				maxLeft = currentElevation
			}
			leftP++
		} else {
			currentElevation = elevations[rightP]
			if maxRight > currentElevation {
				currentWater = maxRight - currentElevation
				totalRainwater += currentWater
			} else {
				maxRight = currentElevation
			}
			rightP--
		}
	}
	return totalRainwater
}

func buildString(str string) string {
	newString := []byte{}

	for i := 0; i < len(str); i++ {
		if str[i] != '#' {
			newString = append(newString, str[i])
		} else {
			length := len(newString)
			if length > 0 {
				newString = newString[:length-1] // Remove o ultimo character adicionado
			}
		}
	}

	return string(newString)
}

func isSameString(s, t string) bool {
	// O(a + b) -- Time Complexity
	// O(a + b) -- Space Complexity, e aqui temos uma oportunidade de melhoria e fazer a alteração na própria string
	newS := buildString(s)
	newT := buildString(t)

	if len(newS) != len(newT) {
		return false
	}

	if newS == newT {
		return true
	}

	return false
}

func isSameStringOptimized(s, t string) bool {
	p1, p2 := len(s)-1, len(t)-1
	backCount := 0

	for p1 >= 0 || p2 >= 0 {
		if (p1 >= 0 && s[p1] == '#') || (p2 >= 0 && t[p2] == '#') {

			if p1 >= 0 && s[p1] == '#' {
				backCount = 2
				for backCount > 0 {
					p1--
					backCount--
					if p1 >= 0 && s[p1] == '#' {
						backCount += 2
					}
				}
			}

			if p2 >= 0 && s[p2] == '#' {
				backCount = 2
				for backCount > 0 {
					p2--
					backCount--
					if p2 >= 0 && t[p2] == '#' {
						backCount += 2
					}
				}
			}
		} else {
			if s[p1] != t[p2] {
				return false
			} else {
				p1--
				p2--
			}
		}
	}
	return true
}

func longestSubstringWithoutRepeating(str string) int {
	if len(str) <= 1 {
		return len(str)
	}

	longest := 0
	strMap := make(map[byte]bool)

	for i := 0; i < len(str); i++ {
		if _, ok := strMap[str[i]]; !ok {
			strMap[str[i]] = true
		} else {
			strMap = make(map[byte]bool)
			strMap[str[i]] = true
		}

		if len(strMap) > longest {
			longest = len(strMap)
		}
	}

	return longest
}

func main() {
	nums := []int{1, 3, 7, 9, 2}
	fmt.Println(findTwoSum(nums, 11))
	fmt.Println(findTwoSumOptimized(nums, 11))

	heights := []int{7, 1, 2, 3, 9}
	fmt.Println(containerMostWater(heights))
	fmt.Println(containerMostWaterOptimized(heights))

	elevations := []int{3, 1, 0, 1, 2}
	fmt.Println(howMuchRainWater(elevations))
	fmt.Println(howMuchRainWaterOptimized(elevations))

	s := "ab#z"
	t := "az"
	fmt.Println(isSameString(s, t))
	fmt.Println(isSameStringOptimized(s, t))

	s2 := "abc#d"
	t2 := "acc#d"
	fmt.Println(isSameString(s2, t2))
	fmt.Println(isSameStringOptimized(s2, t2))

	s3 := "a#####a"
	t3 := "a"
	fmt.Println(isSameString(s3, t3))
	fmt.Println(isSameStringOptimized(s3, t3))

	str1 := "abcbdca"
	fmt.Println(longestSubstringWithoutRepeating(str1))

	str2 := "ccccc"
	fmt.Println(longestSubstringWithoutRepeating(str2))

	str3 := ""
	fmt.Println(longestSubstringWithoutRepeating(str3))

	str4 := "abcbda"
	fmt.Println(longestSubstringWithoutRepeating(str4)) // Função com erro para esse caso de teste, necessário janela deslizante

	str5 := "a"
	fmt.Println(longestSubstringWithoutRepeating(str5))
}
