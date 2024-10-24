package main

import (
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// idk what is this :D
func updateModel(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			selectedData := m.table.SelectedRow()

			if selectedData[0] == "Exec command" {
				if !m.showInput {

					m.showInput = true
					m.selectedRow = m.table.Cursor()
					m.textInput.SetValue("")
					m.textInput.Focus()
					return m, nil
				} else {

					m.showInput = false
					command := m.textInput.Value()
					fmt.Printf("You entered: %s\n", command)

					output := executeMinecraftCommand("cmd " + command)
					m.response = fmt.Sprintf("Command: %s -> %s", command, output)
				}
			} else {

				response := executeAction(selectedData[0])
				m.response = response
			}

		case "esc":

			m.showInput = false
		}

	case tea.WindowSizeMsg:
		m.table.SetWidth(msg.Width)
		m.table.SetHeight(msg.Height - 4)

	}

	var cmd tea.Cmd
	if m.showInput {

		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// exec default command
func execCommand(command string) string {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Err in execute command: %s", err)
	}
	return string(output)
}

// detector list select
func executeAction(action string) string {
	switch action {
	case "Start server":
		return executeMinecraftCommand("srt")
		// doesn't implement 
	case "LuckPerms editor open":
		return executeMinecraftCommand("cmd lp editor")

	case "Stop server":
		return executeMinecraftCommand("stp")

	case "TPS":
		return executeMinecraftCommand("cmd tps")

	case "Players":
		return executeMinecraftCommand("cmd players")

	case "Give admin permission":
		return executeMinecraftCommand("cmd lp")
	}
	return "Unkown action"
}

// Exec minecraftd commands
func executeMinecraftCommand(command string) string {
	cmd := exec.Command("minecraftd", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Err in execute command: %s", err)
	}
	return string(output)
}
