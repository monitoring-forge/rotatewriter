//go:build !windows
// +build !windows

package rotatewriter

import (
	"os"
	"syscall"
)

func openFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND|syscall.O_NOFOLLOW, 0644)
}
