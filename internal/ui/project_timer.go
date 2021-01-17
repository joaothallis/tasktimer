package ui

import (
	"time"

	"github.com/caarlos0/tasktimer/internal/model"
	tea "github.com/charmbracelet/bubbletea"
)

type projectTimerModel struct {
	tasks []model.Task
	d     time.Duration
}

func (m projectTimerModel) Init() tea.Cmd {
	return nil
}

func (m projectTimerModel) Update(msg tea.Msg) (projectTimerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case projectTimerUpdateMsg:
		m.tasks = msg.tasks
	}

	m.d = time.Duration(0)
	for _, t := range m.tasks {
		var z = t.EndAt
		if z.IsZero() {
			z = time.Now()
		}
		m.d += z.Sub(t.StartAt)
	}
	return m, nil
}

func (m projectTimerModel) View() string {
	return midGrayForeground("total: ") + boldSecondaryForeground(m.d.Round(time.Second).String())
}

// msgs and cmds

type projectTimerUpdateMsg struct {
	tasks []model.Task
}

func updateProjectTimerCmd(tasks []model.Task) tea.Cmd {
	return func() tea.Msg {
		return projectTimerUpdateMsg{tasks}
	}
}