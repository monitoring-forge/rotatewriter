//go:build windows
// +build windows

package rotatewriter

import (
	"os"
)

func openFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}
