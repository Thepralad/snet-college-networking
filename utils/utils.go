package utils

import (
	"strings"
)

func IsValidEmail(email string) bool {
	return strings.HasSuffix(email, "@salesiancollege.net")
}
