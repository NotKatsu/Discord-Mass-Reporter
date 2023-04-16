package helpers

import (
	"regexp"
)

func IntCheck(number string) bool {
	var digitCheck = regexp.MustCompile(`^[0-9]+$`)

	if digitCheck.MatchString(number) == true {
		return true
	} else {
		return false
	}

}
