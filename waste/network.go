package waste

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

func Network(interval time.Duration) {
	for {
		user, err := speedtest.FetchUserInfo()
		if err != nil {
			fmt.Println("[NETWORK] Error when fetching user info:", err)
			time.Sleep(time.Minute)
			continue
		}
		serverList, err := speedtest.FetchServers(user)
		if err != nil {
			fmt.Println("[NETWORK] Error when fetching servers:", err)
			time.Sleep(time.Minute)
			continue
		}

		targets := make(speedtest.Servers, 0, len(serverList))
		for _, server := range serverList {
			if server.Latency > 0 {
				targets = append(targets, server)
			}
		}
		if len(targets) == 0 {
			fmt.Println("[NETWORK] No available server to test. Retry in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		// pick random
		s := targets[rand.Int31n(int32(len(targets)))]

		err = s.PingTest()
		if err != nil {
			s.Latency = -1
		}

		err = s.DownloadTest(false)
		if err != nil {
			s.DLSpeed = -1
		}

		err = s.UploadTest(false)
		if err != nil {
			s.ULSpeed = -1
		}

		fmt.Println("[NETWORK] SpeedTest Ping:", s.Latency, ", Download:", s.DLSpeed, ", Upload:", s.ULSpeed, "via", s.String())

		speedtest.GlobalDataManager.Reset()
		runtime.GC()
		time.Sleep(interval)
	}
}
