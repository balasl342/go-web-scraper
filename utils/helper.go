// utils/helper.go
package utils

import (
	"log"
)

// CheckError is a utility function to handle errors.
func CheckError(prefix string, err error) {
	if err != nil {
		log.Fatalf(prefix, "%v", err)
	}
}
