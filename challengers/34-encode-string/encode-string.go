package main

import (
	"fmt"
	"strconv"
)

/*
We have the input string: aaaabbccc
And the output is: 4a2b3c

Clarify:
- If we have a same character on the end array, for example, the character 'a' again: aaaabbccca
The output is 4a2b3c1a

- If we have input abcd
The output is 1a1b1c1d

Nesse caso optei por usar byte por considerar que a string é composta por caracteres ASCII.
Mas poderia usar rune para suportar caracteres Unicode.
Por exemplo, se a string fosse "çáááçááá" o byte não suportaria, mas o rune sim.

------
*Quando usar byte ou rune:*
Use byte se:
A string contém apenas caracteres ASCII.
Eficiência é uma prioridade e você sabe que não haverá caracteres multibyte.
Use rune se:
A string pode conter caracteres Unicode (acentos, emojis, símbolos).
Você quer garantir a manipulação correta de cada caractere, independentemente do número de bytes que ele ocupa.
------

Plan:
- If the string is empty, return an empty string
- Create a variable to store the encoded string
- Create a variable to store the count of the current character
- Create a variable to store the current character
- Iterate over the string
- If the current character is equal to the character in the position i, increment the count
- If the current character is different from the character in the position i, add the count and the current character to the encoded string
- Update the current character to the character in the position i
- Reset the count to 1
- Add the last character to the encoded string
- Return the encoded string
*/

func encode(str string) string { // O(n)
	if len(str) == 0 {
		return ""
	}

	strEncode := ""
	countChar := 0
	currentChar := str[0]

	for i := 0; i < len(str); i++ {
		if currentChar == str[i] {
			countChar++
		} else {
			strEncode += strconv.Itoa(countChar) + string(currentChar)
			currentChar = str[i]
			countChar = 1
		}
	}

	strEncode += strconv.Itoa(countChar) + string(currentChar) // Add the last character
	return strEncode
}

func main() {
	str := "aaaabbccc"
	fmt.Println(encode(str))

	str2 := "abcd"
	fmt.Println(encode(str2))

	str3 := "aaaabbccca"
	fmt.Println(encode(str3))

	str4 := "a"
	fmt.Println(encode(str4))
}
