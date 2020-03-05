package pallet

import "fmt"

// ErrorInvalidConfiguration is thrown when a missing configuration option is required.
type ErrorInvalidConfiguration struct {
	missingConfig string
}

// Error returns a string, implementing the error interface.
func (e ErrorInvalidConfiguration) Error() string {
	return fmt.Sprintf("%s is missing", e.missingConfig)
}
