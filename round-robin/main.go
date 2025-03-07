package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sim.Settings = defaultSettings()

	a := app.New()
	w := a.NewWindow("Round Robin Manager Bot")
	w.Resize(fyne.NewSize(1000, 700))

	performerList := widget.NewList(
		func() int {
			return len(sim.Performers)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			p := sim.Performers[i]
			taskInfo := "No current task"
			if len(p.Tasks) > 0 {
				taskInfo = fmt.Sprintf("%s (C:%d)", p.Tasks[0].Name, p.Tasks[0].Complexity)
			}
			o.(*widget.Label).SetText(fmt.Sprintf("%s [Prod:%d] | %s", p.Name, p.Productivity, taskInfo))
		},
	)

	taskList := widget.NewList(
		func() int {
			if len(sim.Performers) == 0 || selectedIndex < 0 || selectedIndex >= len(sim.Performers) {
				return 0
			}
			return len(sim.Performers[selectedIndex].Tasks)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("task")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			p := sim.Performers[selectedIndex]
			t := p.Tasks[i]
			o.(*widget.Label).SetText(fmt.Sprintf("%s (C:%d)", t.Name, t.Complexity))
		},
	)

	logEntry := widget.NewMultiLineEntry()
	logEntry.SetPlaceHolder("Simulation log...")

	taskScroll := container.NewVScroll(taskList)
	taskScroll.SetMinSize(fyne.NewSize(400, 250))
	logScroll := container.NewVScroll(logEntry)
	logScroll.SetMinSize(fyne.NewSize(400, 150))

	updateUI := func() {
		performerList.Refresh()
		taskList.Refresh()
		sim.Mutex.Lock()
		logText := ""
		for _, entry := range sim.Log {
			logText += entry + "\n"
		}
		sim.Mutex.Unlock()
		logEntry.SetText(logText)
	}

	performerList.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(sim.Performers) {
			selectedIndex = id
			taskList.Refresh()
		}
	}

	newButton := widget.NewButton("New", func() {
		startNewSession(updateUI)
	})
	pauseButton := widget.NewButton("Pause/Resume", func() {
		togglePause(updateUI)
	})
	settingsButton := widget.NewButton("Settings", func() {
		setWin := a.NewWindow("Settings")
		timerEntry := widget.NewEntry()
		timerEntry.SetText(fmt.Sprintf("%.0f", sim.Settings.TimerInterval.Seconds()))
		minPerfEntry := widget.NewEntry()
		minPerfEntry.SetText(strconv.Itoa(sim.Settings.MinPerformers))
		maxPerfEntry := widget.NewEntry()
		maxPerfEntry.SetText(strconv.Itoa(sim.Settings.MaxPerformers))
		minProdEntry := widget.NewEntry()
		minProdEntry.SetText(strconv.Itoa(sim.Settings.MinProductivity))
		maxProdEntry := widget.NewEntry()
		maxProdEntry.SetText(strconv.Itoa(sim.Settings.MaxProductivity))
		minTasksEntry := widget.NewEntry()
		minTasksEntry.SetText(strconv.Itoa(sim.Settings.MinTasks))
		maxTasksEntry := widget.NewEntry()
		maxTasksEntry.SetText(strconv.Itoa(sim.Settings.MaxTasks))
		minCompEntry := widget.NewEntry()
		minCompEntry.SetText(strconv.Itoa(sim.Settings.MinComplexity))
		maxCompEntry := widget.NewEntry()
		maxCompEntry.SetText(strconv.Itoa(sim.Settings.MaxComplexity))

		form := widget.NewForm(
			widget.NewFormItem("Timer (sec)", timerEntry),
			widget.NewFormItem("Min Performers", minPerfEntry),
			widget.NewFormItem("Max Performers", maxPerfEntry),
			widget.NewFormItem("Min Productivity", minProdEntry),
			widget.NewFormItem("Max Productivity", maxProdEntry),
			widget.NewFormItem("Min Tasks", minTasksEntry),
			widget.NewFormItem("Max Tasks", maxTasksEntry),
			widget.NewFormItem("Min Complexity", minCompEntry),
			widget.NewFormItem("Max Complexity", maxCompEntry),
		)
		form.OnSubmit = func() {
			if sec, err := strconv.Atoi(timerEntry.Text); err == nil {
				sim.Settings.TimerInterval = time.Duration(sec) * time.Second
			}
			if v, err := strconv.Atoi(minPerfEntry.Text); err == nil {
				sim.Settings.MinPerformers = v
			}
			if v, err := strconv.Atoi(maxPerfEntry.Text); err == nil {
				sim.Settings.MaxPerformers = v
			}
			if v, err := strconv.Atoi(minProdEntry.Text); err == nil {
				sim.Settings.MinProductivity = v
			}
			if v, err := strconv.Atoi(maxProdEntry.Text); err == nil {
				sim.Settings.MaxProductivity = v
			}
			if v, err := strconv.Atoi(minTasksEntry.Text); err == nil {
				sim.Settings.MinTasks = v
			}
			if v, err := strconv.Atoi(maxTasksEntry.Text); err == nil {
				sim.Settings.MaxTasks = v
			}
			if v, err := strconv.Atoi(minCompEntry.Text); err == nil {
				sim.Settings.MinComplexity = v
			}
			if v, err := strconv.Atoi(maxCompEntry.Text); err == nil {
				sim.Settings.MaxComplexity = v
			}
			setWin.Close()
			dialog.ShowInformation("Settings", "Settings updated", w)
		}
		setWin.SetContent(container.NewVBox(form))
		setWin.Resize(fyne.NewSize(300, 400))
		setWin.Show()
	})

	leftContainer := container.NewBorder(
		container.NewHBox(newButton, pauseButton),
		nil, nil, nil,
		performerList,
	)

	rightContainer := container.NewBorder(
		container.NewHBox(settingsButton),
		nil, nil, nil,
		container.NewVBox(
			widget.NewLabelWithStyle("Tasks for selected performer:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			taskScroll,
			widget.NewSeparator(),
			widget.NewLabelWithStyle("Log:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			logScroll,
		),
	)

	content := container.NewHSplit(leftContainer, rightContainer)
	content.Offset = 0.3

	w.SetContent(content)
	w.ShowAndRun()
}
