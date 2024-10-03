package main

import (
	"fmt"
	"regexp"
)

func main() {
	emailPattern := `^[\w.-]+@[a-zA-Z\d.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)

	emails := []string{"user@example.com", "invalid-email@", "test.user@domain.co"}

	for _, email := range emails {
		if re.MatchString(email) {
			fmt.Printf("%s 是有效的電子郵件地址。\n", email)
		} else {
			fmt.Printf("%s 是無效的電子郵件地址。\n", email)
		}
	}
}
