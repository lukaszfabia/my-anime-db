package utils

import (
	"fmt"

	"golang.org/x/exp/rand"
)

/*
GenerateCode generates a random 6 digit code.

  - returns:

    string: the generated code
*/
func GenerateCode() string {
	min := int64(100000)
	max := int64(999999)
	return fmt.Sprintf("%06d", min+rand.Int63n(max-min))
}
