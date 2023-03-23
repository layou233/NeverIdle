package waste

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"go.einride.tech/pid"
	"log"
	"runtime"
	"time"
)

func CPUPercent(referencePercent float64) {
	size := 100000.0
	controller = RunPID(InitMachine(size), referencePercent, size/1000, false)
}

var controller *pid.Controller

type Machine struct {
	runtimePeriod   time.Duration // ms
	maxControlValue float64
	busyTime        int64
	idleTime        time.Duration

	revolution float64
}

func InitMachine(max float64) *Machine {
	e := &Machine{runtimePeriod: time.Second, maxControlValue: max}
	e.busyTime = 0
	e.idleTime = time.Duration(e.maxControlValue)
	for i := 0; i < runtime.NumCPU(); i++ {
		go e.Run()
	}
	return e
}

func (m *Machine) Run() {
	for {
		startTime := time.Now().UnixNano()
		for time.Now().UnixNano()-startTime < m.busyTime {
		}
		time.Sleep(m.idleTime)
	}
}

func (m *Machine) Measure() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatalln(err.Error())
		return -1
	}
	return percent[0]
}

func (m *Machine) Control(value float64) {
	// value range [0, maxControlValue] (unit: Nanosecond)
	m.revolution += value
	if m.revolution < 0 {
		m.revolution = 0
		controller.Reset()
	} else if m.revolution > m.maxControlValue {
		m.revolution = m.maxControlValue
	}
	value = m.revolution / m.maxControlValue
	totalTime := m.runtimePeriod.Nanoseconds()
	m.busyTime = int64(float64(totalTime) * value)
	m.idleTime = time.Duration(totalTime - m.busyTime)
}
