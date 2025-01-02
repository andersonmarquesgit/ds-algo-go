/*
Given two arrays of strings list1 and list2, find the common strings with the least index sum.

A common string is a string that appeared in both list1 and list2.

A common string with the least index sum is a common string such that if it appeared at list1[i] and list2[j] then i + j
should be the minimum value among all the other common strings.

Return all the common strings with the least index sum. Return the answer in any order.

Example 1:

Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["Piatti","The Grill at Torrey Pines",
"Hungry Hunter Steakhouse","Shogun"]
Output: ["Shogun"]
Explanation: The only common string is "Shogun".
Example 2:

Input: list1 = ["Shogun","Tapioca Express","Burger King","KFC"], list2 = ["KFC","Shogun","Burger King"]
Output: ["Shogun"]
Explanation: The common string with the least index sum is "Shogun" with index sum = (0 + 1) = 1.
Example 3:

Input: list1 = ["happy","sad","good"], list2 = ["sad","happy","good"]
Output: ["sad","happy"]
Explanation: There are three common strings:
"happy" with index sum = (0 + 1) = 1.
"sad" with index sum = (1 + 0) = 1.
"good" with index sum = (2 + 2) = 4.
The strings with the least index sum are "sad" and "happy".
*/
package main

func findRestaurant(list1 []string, list2 []string) []string {
	// create a map to store the index of each string in list1
	indexMap := make(map[string]int)
	for i, v := range list1 {
		indexMap[v] = i
	}

	// create a map to store the index sum of each common string
	indexSumMap := make(map[string]int)
	minIndexSum := len(list1) + len(list2)
	for i, v := range list2 {
		if index, ok := indexMap[v]; ok {
			indexSum := i + index
			indexSumMap[v] = indexSum
			if indexSum < minIndexSum {
				minIndexSum = indexSum
			}
		}
	}

	// create a slice to store the common strings with the least index sum
	result := []string{}
	for k, v := range indexSumMap {
		if v == minIndexSum {
			result = append(result, k)
		}
	}

	return result
}
