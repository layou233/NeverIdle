package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/layou233/neveridle/controller"
	"github.com/layou233/neveridle/waste"
)

const Version = "0.2.3"

var (
	FlagCPUPercent             = flag.Float64("cp", 0, "Percent of CPU waste")
	FlagCPU                    = flag.Duration("c", 0, "Interval for CPU waste")
	FlagMemory                 = flag.Int("m", 0, "GiB of memory waste")
	FlagMemoryPercent          = flag.Float64("mp", 0, "Percent of memory waste (e.g., 0.2 for 20%)")
	FlagNetwork                = flag.Duration("n", 0, "Interval for network speed test")
	FlagNetworkConnectionCount = flag.Int("t", 10, "Set concurrent connections for network speed test")
	FlagNightStart             = flag.Int("night-start", 0, "深夜开始小时 (0-23)")
	FlagNightEnd               = flag.Int("night-end", 6, "深夜结束小时 (0-23，支持跨夜如 22-6)")
	FlagIdleThreshold          = flag.Int("idle", 5, "网络空闲连接数阈值（小于此值才执行浪费）")
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
	} else if *FlagMemoryPercent != 0 {
		nothingEnabled = false
		fmt.Println("====================")
		fmt.Println("Starting memory wasting with percent", *FlagMemoryPercent)
		go waste.MemoryPercent(*FlagMemoryPercent)
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
		go waste.Network(*FlagNetwork, *FlagNetworkConnectionCount, *FlagNightStart, *FlagNightEnd, *FlagIdleThreshold)
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
