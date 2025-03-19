package main

import (
	"fmt"
	"regexp"
	"strings"
)

func numUniqueEmails(emails []string) int {
	uniqueEmailsMap := make(map[string]bool)

	for _, email := range emails {
		curEmail := strings.Split(email, "@")
		domain := curEmail[1]
		localName := strings.Split(curEmail[0], "+")
		regex := regexp.MustCompile(`[^a-z]`)
		firstPart := strings.ToLower(regex.ReplaceAllString(localName[0], ""))

		finalEmail := firstPart + "@" + domain

		if !uniqueEmailsMap[finalEmail] {
			uniqueEmailsMap[finalEmail] = true
		}
	}

	return len(uniqueEmailsMap)
}

// Ambos os loops são O(n), mas o segundo é mais eficiente por não ter que fazer a substituição de caracteres usando regex
func numUniqueEmailsOptimized(emails []string) int {
	uniqueEmailsMap := make(map[string]bool)

	for _, email := range emails {
		curEmail := strings.Split(email, "@")
		domain := curEmail[1]
		local := curEmail[0]

		// Remover tudo após '+'
		if idx := strings.Index(local, "+"); idx != -1 {
			local = local[:idx]
		}

		cleanLocal := make([]byte, 0, len(local))
		for i := 0; i < len(local); i++ {
			if local[i] != '.' {
				cleanLocal = append(cleanLocal, local[i])
			}
		}

		finalEmail := string(cleanLocal) + "@" + domain
		uniqueEmailsMap[finalEmail] = true
	}

	return len(uniqueEmailsMap)
}

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}

	fmt.Println(numUniqueEmails(emails))
	fmt.Println(numUniqueEmailsOptimized(emails))

}
