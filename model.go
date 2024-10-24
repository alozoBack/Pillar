package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	table       table.Model
	textInput   textinput.Model
	showInput   bool
	selectedRow int
	response    string
}

func NewModel() model {

	columns := []table.Column{
		{Title: "", Width: 30},
	}

	rows := []table.Row{
		{"Exec command"},
		{"LuckPerms editor open"},
		{"Start server"},
		{"Stop server"},
		{"TPS"},
		{"Players"},
		{"Give admin permission"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	ti := textinput.New()
	ti.Placeholder = "Enter something"
	ti.Focus()

	return model{
		table:     t,
		textInput: ti,
		response:  "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return updateModel(m, msg)
}

func (m model) View() string {
	return renderView(m)
}
