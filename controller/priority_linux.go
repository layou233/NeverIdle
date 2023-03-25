//go:build linux

package controller

import (
	"os"
	"syscall"
)

const ioprioClassShift = 13

type ioprioClass = int

const (
	ioprioClassRT ioprioClass = (iota + 1) << ioprioClassShift
	ioprioClassBE
	ioprioClassIdle
)

const (
	ioprioWhoProcess = iota + 1
	ioprioWhoPGRP
	ioprioWhoUser
)

func ioprioSet(class ioprioClass, value int) {
	// error return ignored
	syscall.Syscall(syscall.SYS_IOPRIO_SET,
		uintptr(ioprioWhoPGRP), uintptr(os.Getpid()),
		uintptr(class)|uintptr(value))
}

func SetPriority(nice int) error {
	// https://github.com/syncthing/syncthing/issues/4628
	// Move ourselves to a new process group so that we can use the process
	// group variants of Setpriority etc. to affect all of our threads in one
	// go. If this fails, bail, so that we don't affect things we shouldn't.
	if err := syscall.Setpgid(0, 0); err != nil {
		return err
	}

	return syscall.Setpriority(syscall.PRIO_PGRP, 0, nice)
}

func SetWorstPriority() {
	SetPriority(19)

	// 0 through 7 being the range. Error return ignored.
	ioprioSet(ioprioClassIdle, 7)
}
