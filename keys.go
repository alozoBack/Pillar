package main

import "github.com/charmbracelet/bubbles/key"

// Keybindings for navigation
var keys = struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
}{
	Up:    key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "move up")),
	Down:  key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "move down")),
	Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
}
