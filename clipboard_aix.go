//go:build aix
// +build aix

package clipboard

import (
	"errors"
)


func readAll() (string, error) {
	return "", errors.New("Clipboard operations are not supported on AIX")
}

func writeAll(text string) error {
	return errors.New("Clipboard operations are not supported on AIX")
}
