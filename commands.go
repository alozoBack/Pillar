package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/gorcon/rcon"
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
					log.Info("EXECCOMMAND:" + command)
					output := execMinecraftRCON(command)
					m.response = fmt.Sprintf("Command: %s -> %s", command, output)
				}
			} else if selectedData[0] == "Give admin permission" {
				adminGive := os.Getenv("ADMINCOMMAND")
				if !m.showInput {
					m.showInput = true
					m.selectedRow = m.table.Cursor()
					m.textInput.SetValue("")
					m.textInput.Focus()
					return m, nil
				} else {
					m.showInput = false
					user := m.textInput.Value()
					log.Info("adminGive:" + adminGive)
					log.Info("PERMSCOMMAND:" + strings.Replace(adminGive, "%user%", user, -1))
					output := execMinecraftRCON(strings.Replace(adminGive, "%user%", user, -1))
					m.response = fmt.Sprintf("permission for %s given %s", user, output)
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

// detector list select
func executeAction(action string) string {

	switch action {
	/*TODO
	case "LuckPerms editor open":
		return execMinecraftRCON("lp editor")
	*/
	case "TPS":
		return execMinecraftRCON("tps")

	case "Players":
		return execMinecraftRCON("list")

	case "Stop server":
		return execMinecraftRCON("stop")
	}

	return "Unkown action"
}

func execMinecraftRCON(command string) string {
	hostRcon := os.Getenv("HOSTRCON")
	passwordRcon := os.Getenv("PASSWORDRCON")

	conn, err := rcon.Dial(hostRcon, passwordRcon)
	if err != nil {
		log.Error(err)
		return fmt.Sprintf("ERROR %s", err)
	}
	defer conn.Close()
	response, err := conn.Execute(command)
	if err != nil {
		log.Error(err)
		return fmt.Sprintf("ERROR %s", err)
	}
	return response
}
