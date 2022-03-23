package task

import (
	"github.com/nelsongp/composite/comp"
	"strconv"
	"strings"
)

type Task struct {
	Name        string
	Responsable string
	Prices      int
	SubTasks    []comp.Item
}

//Add..
func (t *Task) Add(i comp.Item) {
	t.SubTasks = append(t.SubTasks, i)
}

// String ...
func (t *Task) String() string {
	sb := strings.Builder{}
	sb.WriteString("\t|--")
	sb.WriteString(t.Name)
	sb.WriteString(" - ")
	sb.WriteString(t.Responsable)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(t.Price()))
	sb.WriteString("\n")
	for _, v := range t.SubTasks {
		sb.WriteString(v.String())
	}

	return sb.String()
}

// Price ...
func (t *Task) Price() int {
	price := t.Prices
	for _, v := range t.SubTasks {
		price += v.Price()
	}

	return price
}
