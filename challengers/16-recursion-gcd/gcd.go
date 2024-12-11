/*
Input Format
The first argument is an integer A.
The second argument is an integer B.


Output Format
Return the Greatest Common Divisor of A and B


Example Input
Input 1:
A = 6
B = 9


Example Output
Output 1:
3


Example Explanation
Explanation 1:
3 is the GCD of 6 and 9
*/

package main

func gcd(A int, B int) int {
	if B == 0 { // If B is 0, the GCD is A
		return A
	}
	return gcd(B, A%B) // Otherwise, call the function recursively with B and A%B, because the GCD of A and B is the same as the GCD of B and A%B
}

func main() {
	A := 6
	B := 9
	println(gcd(A, B)) // 3
}
