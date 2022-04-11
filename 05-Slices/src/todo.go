package src

import (
	s "github.com/inancgumus/prettyslice"
)

func Todo() {
	var todo []string
	todo = append(todo, "code")
	tommorow := []string{"TDD", "Interfaces", "methods"}
	todo = append(todo, tommorow...)

	s.Show("todo", todo)
}
