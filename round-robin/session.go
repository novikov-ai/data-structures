package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Simulation struct {
	Performers  []*Performer
	Cycle       int
	Log         []string
	Timer       *time.Ticker
	Paused      bool
	Settings    Settings
	StopChannel chan bool
	Mutex       sync.Mutex
}

var sim Simulation

var selectedIndex int = 0

func startSimulationTimer(updateUI func()) {
	sim.Mutex.Lock()
	if sim.Timer != nil {
		sim.StopChannel <- true
	}
	sim.Paused = false
	sim.StopChannel = make(chan bool)
	sim.Timer = time.NewTicker(sim.Settings.TimerInterval)
	sim.Mutex.Unlock()

	go func() {
		for {
			select {
			case <-sim.Timer.C:
				sim.Mutex.Lock()
				sim.Cycle++
				// Process each performer's current task.
				for _, p := range sim.Performers {
					if len(p.Tasks) > 0 {
						p.Tasks[0].Complexity -= p.Productivity
						if p.Tasks[0].Complexity <= 0 {
							logEntry := fmt.Sprintf("Cycle %d: %s finished task %s", sim.Cycle, p.Name, p.Tasks[0].Name)
							sim.Log = append(sim.Log, logEntry)
							p.Tasks = p.Tasks[1:]
						}
					}
				}
				// 50% chance to invoke Method B.
				if rand.Float64() < 0.5 {
					redistributeFirstTasks(sim.Performers)
				}
				sim.Mutex.Unlock()
				updateUI()
			case <-sim.StopChannel:
				sim.Timer.Stop()
				return
			}
		}
	}()
}

func startNewSession(updateUI func()) {
	if sim.Timer != nil {
		sim.StopChannel <- true
	}
	sim.Mutex.Lock()
	sim.Cycle = 0
	sim.Log = []string{}
	selectedIndex = 0
	numPerformers := rand.Intn(sim.Settings.MaxPerformers-sim.Settings.MinPerformers+1) + sim.Settings.MinPerformers
	sim.Performers = make([]*Performer, numPerformers)
	for i := 0; i < numPerformers; i++ {
		sim.Performers[i] = &Performer{
			Name:         randomPerformerName(),
			Productivity: rand.Intn(sim.Settings.MaxProductivity-sim.Settings.MinProductivity+1) + sim.Settings.MinProductivity,
			Tasks:        []*Task{},
		}
	}
	numTasks := rand.Intn(sim.Settings.MaxTasks-sim.Settings.MinTasks+1) + sim.Settings.MinTasks
	tasks := make([]*Task, numTasks)
	for i := 0; i < numTasks; i++ {
		tasks[i] = &Task{
			Name:       randomTaskName(),
			Complexity: rand.Intn(sim.Settings.MaxComplexity-sim.Settings.MinComplexity+1) + sim.Settings.MinComplexity,
		}
	}
	assignTasksRoundRobin(sim.Performers, tasks)
	sim.Mutex.Unlock()
	updateUI()
	startSimulationTimer(updateUI)
}

func togglePause(updateUI func()) {
	sim.Mutex.Lock()
	defer sim.Mutex.Unlock()
	if sim.Paused {
		startSimulationTimer(updateUI)
	} else {
		if sim.Timer != nil {
			sim.StopChannel <- true
		}
		sim.Paused = true
	}
	updateUI()
}