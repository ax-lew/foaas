package errors

import (
	"fmt"
)

func NewInternalError(message string) error {
	return fmt.Errorf("internal error: %s", message)
}
