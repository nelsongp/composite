package subtask

import (
	"fmt"
	"github.com/nelsongp/composite/comp"
	"strings"
)

type SubTask struct {
	Name   string
	Prices int
}

//Add..
func (s *SubTask) Add(i comp.Item) {
	fmt.Println("Yo no acepto más subtareas")
}

// String ...
func (s *SubTask) String() string {
	sb := strings.Builder{}
	sb.WriteString("\t\t|-- ")
	sb.WriteString(s.Name)
	sb.WriteString("\n")

	return sb.String()
}

// Price ...
func (s *SubTask) Price() int {
	return s.Prices
}
