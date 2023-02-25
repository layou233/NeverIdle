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

		var targets = speedtest.Servers{}
		for _, server := range serverList {
			if server.Latency > 0 {
				targets = append(targets, server)
			}
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

		fmt.Println("[NETWORK] SpeedTest Ping:", s.Latency, ",", s.DLSpeed, ",", "Upload:", s.ULSpeed, "via", s.String())

		speedtest.GlobalDataManager.Reset()
		runtime.GC()
		time.Sleep(interval)
	}
}
