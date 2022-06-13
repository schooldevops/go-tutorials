package email

import "strings"

var freemails = []string{"gmail.com", "yahoo.com", "outlook.com"}

func IsFreemail(email string) bool {
	for _, provider := range freemails {
		if strings.Contains(email, provider) {
			return true
		}
	}
	return false
}
