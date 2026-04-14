package waste

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

var Buffers []*GiBObject
var MiBBuffers [][]byte

const (
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
)

type GiBObject struct {
	B [GiB]byte
}

func Memory(gib int) {
	Buffers = make([]*GiBObject, 0, gib)
	for gib > 0 {
		o := new(GiBObject)
		rand.Read(o.B[:])
		Buffers = append(Buffers, o)
		gib -= 1
	}
}

// MemoryPercent dynamically keeps system memory usage at the specified percentage (0.0 to 1.0)
func MemoryPercent(targetPercent float64) {
	MiBBuffers = make([][]byte, 0)
	
	// 每次分配/释放的步长为 50MiB
	stepSize := 50 * MiB

	for {
		v, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("[MEMORY] 获取内存信息失败:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		targetBytes := uint64(float64(v.Total) * targetPercent)
		
		if v.Used < targetBytes {
			// 需要增加内存
			diff := targetBytes - v.Used
			if diff > uint64(stepSize) {
				// 每次最多分配 stepSize 避免突然 OOM
				buf := make([]byte, stepSize)
				rand.Read(buf) // 填入随机数据防止被系统优化
				MiBBuffers = append(MiBBuffers, buf)
			}
		} else if v.Used > targetBytes {
			// 需要释放内存
			diff := v.Used - targetBytes
			if diff > uint64(stepSize) && len(MiBBuffers) > 0 {
				// 释放一个块
				MiBBuffers[len(MiBBuffers)-1] = nil
				MiBBuffers = MiBBuffers[:len(MiBBuffers)-1]
				runtime.GC() // 强制垃圾回收以释放内存
			}
		}
		
		time.Sleep(3 * time.Second) // 每3秒检查一次
	}
}
