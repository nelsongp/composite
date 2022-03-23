package project

import (
	"github.com/nelsongp/composite/comp"
	"strconv"
	"strings"
)

type Project struct {
	Name string
	Task []comp.Item
}

//Add..
func (p *Project) Add(i comp.Item) {
	p.Task = append(p.Task, i)
}

//String
func (p *Project) String() string {
	sb := strings.Builder{}
	sb.WriteString(p.Name)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(p.Price()))
	sb.WriteString("\n")
	for _, v := range p.Task {
		sb.WriteString(v.String())
	}
	return sb.String()
}

//Price
func (p *Project) Price() int {
	price := 0
	for _, v := range p.Task {
		price += v.Price()
	}
	return price
}
