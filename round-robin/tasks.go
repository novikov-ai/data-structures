package main

type Task struct {
	Name       string
	Complexity int
}

type Performer struct {
	Name         string
	Productivity int
	Tasks        []*Task
}

// Method A
func assignTasksRoundRobin(performers []*Performer, tasks []*Task) {
	for _, p := range performers {
		p.Tasks = []*Task{}
	}
	n := len(performers)
	for i, task := range tasks {
		idx := i % n
		performers[idx].Tasks = append(performers[idx].Tasks, task)
	}
}

// Method B
func redistributeFirstTasks(performers []*Performer) {
	if len(performers) == 0 {
		return
	}
	n := len(performers)
	firstTasks := make([]*Task, n)
	for i, p := range performers {
		if len(p.Tasks) > 0 {
			firstTasks[i] = p.Tasks[0]
		}
	}
	for _, p := range performers {
		if len(p.Tasks) > 0 {
			p.Tasks = p.Tasks[1:]
		}
	}
	for i := 0; i < n; i++ {
		if firstTasks[i] != nil {
			next := (i + 1) % n
			performers[next].Tasks = append([]*Task{firstTasks[i]}, performers[next].Tasks...)
		}
	}
}