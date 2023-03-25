//go:build !unix && !windows

package controller

import (
	"os"
)

func SetPriority(int) error {
	return os.ErrInvalid
}

func SetWorstPriority() {}
