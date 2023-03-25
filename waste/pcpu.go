package waste

import (
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/layou233/neveridle/controller"

	"github.com/shirou/gopsutil/v3/cpu"
	"go.einride.tech/pid"
	"golang.org/x/crypto/chacha20"
)

var c *pid.Controller

func CPUPercent(referencePercent float64) {
	maxStep := 100000.0
	rateImpact := maxStep / 1000
	c = controller.RunPID(newMachine(maxStep), referencePercent, rateImpact, false)
}

type machine struct {
	runtimePeriod   time.Duration // ms
	maxControlValue float64
	busyTime        int64
	idleTime        time.Duration

	revolution float64
}

func newMachine(maxStep float64) *machine {
	e := &machine{runtimePeriod: time.Second, maxControlValue: maxStep}
	e.busyTime = 0
	e.idleTime = time.Duration(e.maxControlValue)
	for i := 0; i < runtime.NumCPU(); i++ {
		go e.Run()
	}
	return e
}

func (m *machine) Run() {
	var buffer []byte
	if len(Buffers) > 0 {
		buffer = Buffers[0].B[:4*MiB]
	} else {
		buffer = make([]byte, 4*MiB)
	}
	_, _ = rand.Read(buffer)
	cipher, _ := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
	for {
		startTime := time.Now().UnixNano()
		for time.Now().UnixNano()-startTime < m.busyTime {
			cipher.XORKeyStream(buffer, buffer)
			newCipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
			if err == nil {
				cipher = newCipher
			}
		}
		time.Sleep(m.idleTime)
	}
}

func (m *machine) Measure() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatalln(err)
		return -1
	}
	return percent[0]
}

func (m *machine) Control(value float64) {
	// value range [0, maxControlValue] (unit: Nanosecond)
	m.revolution += value
	if m.revolution < 0 {
		m.revolution = 0
		c.Reset()
	} else if m.revolution > m.maxControlValue {
		m.revolution = m.maxControlValue
	}
	value = m.revolution / m.maxControlValue
	totalTime := m.runtimePeriod.Nanoseconds()
	m.busyTime = int64(float64(totalTime) * value)
	m.idleTime = time.Duration(totalTime - m.busyTime)
}
