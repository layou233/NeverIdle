//go:build windows

package controller

import "golang.org/x/sys/windows"

func SetPriority(nice int) error {
	return windows.SetPriorityClass(windows.CurrentProcess(), uint32(nice))
}

func SetWorstPriority() {
	SetPriority(windows.IDLE_PRIORITY_CLASS)
}
