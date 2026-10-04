package validator

import "strings"

func IsValidEmail(email string) bool {
	lastDot := strings.LastIndex(email, ".")
	if strings.ContainsRune(email, '@') && lastDot != -1 && lastDot < len(email)-1 {
		return true
	}

	return false
}

func IsAdult(age int) bool {
	if age < 18 {
		return false
	}

	return true
}
