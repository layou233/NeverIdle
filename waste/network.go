package waste

import (
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

// isNightTime 判断当前是否处于深夜时段
// 支持跨夜（例如 nightStart=22, nightEnd=6 表示 22:00~次日 06:00）
func isNightTime(startHour, endHour int) bool {
	h := time.Now().Hour()
	if startHour <= endHour {
		return h >= startHour && h < endHour
	}
	// 跨夜情况
	return h >= startHour || h < endHour
}

// isNetworkIdle 判断网络是否空闲（通过 ss 统计 established 连接数）
// 返回 true = 空闲，可以执行浪费
func isNetworkIdle(threshold int) bool {
	// 使用 ss -Htn state established | wc -l  （-H 无表头，-t TCP，-n 数字格式）
	cmd := exec.Command("sh", "-c", `ss -Htn state established | wc -l`)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("[NETWORK] 空闲检测失败（ss 命令执行错误）:", err)
		return false // 安全起见，不执行
	}

	countStr := strings.TrimSpace(string(out))
	count, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Println("[NETWORK] 解析连接数失败:", err)
		return false
	}

	fmt.Printf("[NETWORK] 当前 established 连接数: %d (阈值 %d)\n", count, threshold)
	return count < threshold
}

func Network(interval time.Duration, connectionCount int, nightStart, nightEnd, idleThreshold int) {
	cache := false
	st := speedtest.New()
	st.SetNThread(connectionCount)
	var targets speedtest.Servers

	for {
		// ====================== 新增：深夜 + 空闲 双重判断 ======================
		if !isNightTime(nightStart, nightEnd) || !isNetworkIdle(idleThreshold) {
			fmt.Println("[NETWORK] 非深夜或网络不空闲，跳过本次浪费，5分钟后重检...")
			time.Sleep(5 * time.Minute) // 频繁检查，避免错过深夜窗口
			continue
		}
		// =====================================================================

		if !cache {
			_, err := st.FetchUserInfo()
			if err != nil {
				fmt.Println("[NETWORK] Error when fetching user info:", err)
				time.Sleep(time.Minute)
				continue
			}
			serverList, err := st.FetchServers()
			if err != nil {
				fmt.Println("[NETWORK] Error when fetching servers:", err)
				time.Sleep(time.Minute)
				continue
			}

			targets = *serverList.Available()
			if len(targets) == 0 {
				fmt.Println("[NETWORK] No available server to test. Retry in 5 seconds...")
				time.Sleep(5 * time.Second)
				continue
			}
			if float64(len(targets))/float64(len(serverList)) > 0.5 {
				cache = true
			}
		}

		// pick random as main server
		s := targets[rand.Int31n(int32(len(targets)))]

		err := s.PingTest(nil)
		if err != nil {
			s.Latency = -1
		}

		err = s.MultiDownloadTestContext(context.Background(), targets)
		if err != nil {
			s.DLSpeed = -1
		}

		err = s.MultiUploadTestContext(context.Background(), targets)
		if err != nil {
			s.ULSpeed = -1
		}

		fmt.Println("[NETWORK] SpeedTest Ping:", s.Latency, ", Download:", s.DLSpeed, ", Upload:", s.ULSpeed, "mainServer", s.String())

		st.Reset()
		runtime.GC()
		time.Sleep(interval) // 执行完一次浪费后，按你设置的间隔（-n）再睡
	}
}
