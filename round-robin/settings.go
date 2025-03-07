package main

import (
	"math/rand"
	"strconv"
	"time"
)

var performerNames = []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Heidi"}
var taskNames = []string{"Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel"}

func randomPerformerName() string {
	return performerNames[rand.Intn(len(performerNames))] + strconv.Itoa(rand.Intn(100))
}

func randomTaskName() string {
	return taskNames[rand.Intn(len(taskNames))] + "-" + strconv.Itoa(rand.Intn(1000))
}

type Settings struct {
	TimerInterval                    time.Duration // in seconds
	MinPerformers, MaxPerformers     int
	MinProductivity, MaxProductivity int
	MinTasks, MaxTasks               int
	MinComplexity, MaxComplexity     int
}

func defaultSettings() Settings {
	return Settings{
		TimerInterval:   1 * time.Second,
		MinPerformers:   3,
		MaxPerformers:   6,
		MinProductivity: 2,
		MaxProductivity: 10,
		MinTasks:        5,
		MaxTasks:        15,
		MinComplexity:   20,
		MaxComplexity:   100,
	}
}