//go:build unix && !linux

package controller

import "syscall"

func SetPriority(nice int) error {
	return syscall.Setpriority(syscall.PRIO_PROCESS, 0, nice)
}

func SetWorstPriority() {
	SetPriority(19)
}
