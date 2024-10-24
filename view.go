package main

import (
	"fmt"
)

func renderView(m model) string {
	if m.showInput {

		return fmt.Sprintf(
			"Enter the command: %s",
			m.textInput.View(),
		)
	}

	return fmt.Sprintf(
		"%s\n %s \n",
		m.table.View(),
		m.response,
	)
}
