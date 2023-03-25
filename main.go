package main

import (
	"flag"
	"fmt"
	"github.com/layou233/neveridle/controller"
	"math/rand"
	"runtime"
	"time"

	"github.com/layou233/neveridle/waste"
)

const Version = "0.2.3"

var (
	FlagCPUPercent             = flag.Float64("cp", 0, "Percent of CPU waste")
	FlagCPU                    = flag.Duration("c", 0, "Interval for CPU waste")
	FlagMemory                 = flag.Int("m", 0, "GiB of memory waste")
	FlagNetwork                = flag.Duration("n", 0, "Interval for network speed test")
	FlagNetworkConnectionCount = flag.Int("t", 10, "Set concurrent connections for network speed test")
	FlagPriority               = flag.Int("p", 666, "Set process priority value")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("NeverIdle", Version, "- Getting worse from here.")
	fmt.Println("Platform:", runtime.GOOS, ",", runtime.GOARCH, ",", runtime.Version())
	fmt.Println("GitHub: https://github.com/layou233/NeverIdle")

	flag.Parse()
	nothingEnabled := true

	if *FlagPriority == 666 {
		fmt.Println("[PRIORITY] Use the worst priority by default.")
		controller.SetWorstPriority()
	} else {
		err := controller.SetPriority(*FlagPriority)
		if err != nil {
			fmt.Println("[PRIORITY] Error when set priority:", err)
		}
	}

	if *FlagMemory != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting memory wasting of", *FlagMemory, "GiB")
		go waste.Memory(*FlagMemory)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if *FlagCPU != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting CPU wasting with interval", *FlagCPU)
		go waste.CPU(*FlagCPU)
		runtime.Gosched()
		fmt.Println("====================")
	} else if *FlagCPUPercent != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting CPU wasting with percent", *FlagCPUPercent)
		waste.CPUPercent(*FlagCPUPercent)
		runtime.Gosched()
		fmt.Println("====================")
	}

	if *FlagNetwork != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting network speed testing with interval", *FlagNetwork)
		go waste.Network(*FlagNetwork, *FlagNetworkConnectionCount)
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
