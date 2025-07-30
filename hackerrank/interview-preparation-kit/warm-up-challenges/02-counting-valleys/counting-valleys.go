package main

import "fmt"

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 *
 * [DDUUUUDD]
 *       /\
 *      /  \
 *  \  /
 *   \/
 */

/*
Sample Input
8
UDDDUDUU

Sample Output
1

Explanation
If we represent _ as sea level, a step up as /, and a step down as \, the hike can be drawn as:

_/\      _
   \    /
    \/\/
The hiker enters and leaves one valley.
*/

func countingValleys(steps int32, path string) int32 {
	seaLevel := 0
	valleys := 0
	for _, step := range path {
		if step == 'U' {
			seaLevel++
		} else {
			seaLevel--
		}
		if seaLevel == 0 && step == 'U' {
			valleys++
		}
	}

	return int32(valleys)
}

func main() {
	path := "UDDDUDUU"
	fmt.Println(countingValleys(8, path))

	path2 := "DDUUDDUDUUUD"
	fmt.Println(countingValleys(12, path2))

	path3 := "DDDUUUUDDDUUUUDD"
	fmt.Println(countingValleys(16, path3))
}
