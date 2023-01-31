package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var (
	FlagCPU     = flag.Duration("c", 0, "Interval for CPU waste")
	FlagMemory  = flag.Int("m", 0, "GiB of memory waste")
	FlagNetwork = flag.Duration("n", 0, "Interval for network speed test")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("NeverIdle", Version, "- Getting worse from here.")
	fmt.Println("Platform:", runtime.GOOS, ",", runtime.GOARCH, ",", runtime.Version())
	fmt.Println("GitHub: https://github.com/layou233/NeverIdle")

	flag.Parse()
	nothingEnabled := true

	if *FlagMemory != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting memory wasting of", *FlagMemory, "GiB")
		go WasteMemory(*FlagMemory)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if *FlagCPU != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting CPU wasting with interval", *FlagCPU)
		go WasteCPU(*FlagCPU)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if *FlagNetwork != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting network speed testing with interval", *FlagNetwork)
		go WasteNetwork(*FlagNetwork)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if nothingEnabled {
		flag.PrintDefaults()
	} else {
		// fatal error: all goroutines are asleep - deadlock!
		// select {} // fall asleep

		for {
			time.Sleep(24 * time.Hour)
		}
	}
}
