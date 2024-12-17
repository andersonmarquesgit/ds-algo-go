/*
Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Example 1:

Input: rowIndex = 3
Output: [1,3,3,1]
Example 2:

Input: rowIndex = 0
Output: [1]
Example 3:

Input: rowIndex = 1
Output: [1,1]

Constraints:

0 <= rowIndex <= 33

Follow up: Could you optimize your algorithm to use only O(rowIndex) extra space?
*/

package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 { // Base case
		return []int{1}
	}

	if rowIndex == 1 { // Base case
		return []int{1, 1}
	}

	previousRow := getRow(rowIndex - 1)
	row := make([]int, rowIndex+1)
	row[0], row[rowIndex] = 1, 1

	for i := 1; i < rowIndex; i++ {
		row[i] = previousRow[i-1] + previousRow[i]
	}

	return row
}
