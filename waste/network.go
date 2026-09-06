package waste

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

func Network(interval time.Duration, connectionCount int) {
	cache := false
	client := speedtest.New()
	client.SetNThread(connectionCount)
	var targets speedtest.Servers
	for {
		if !cache {
			_, err := client.FetchUserInfo()
			if err != nil {
				fmt.Println("[NETWORK] Error when fetching user info:", err)
				time.Sleep(time.Minute)
				continue
			}
			serverList, err := client.FetchServers()
			if err != nil {
				fmt.Println("[NETWORK] Error when fetching servers:", err)
				time.Sleep(time.Minute)
				continue
			}

			avail := serverList.Available()
			if avail == nil || len(*avail) == 0 {
				fmt.Println("[NETWORK] No available server to test. Retry in 5 seconds...")
				time.Sleep(5 * time.Second)
				continue
			}
			targets = *avail
			if float64(len(targets))/float64(len(serverList)) > 0.5 {
				cache = true
			}
		}

		// pick random as main server
		s := targets[rand.Int31n(int32(len(targets)))]

		err := s.PingTestContext(context.Background(), nil)
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

		client.Reset()
		runtime.GC()
		time.Sleep(interval)
	}
}
