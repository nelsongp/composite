package main

import (
	"fmt"
	"github.com/nelsongp/composite/project"
	"github.com/nelsongp/composite/subtask"
	"github.com/nelsongp/composite/task"
)

func main() {
	p := project.Project{Name: "Upload Images"}
	t1 := task.Task{Name: "Mockup", Responsable: "UI Desigener", Prices: 1000}

	t2 := task.Task{Name: "Markup", Responsable: "Web Designer"}
	st21 := subtask.SubTask{Name: "HTML", Prices: 300}
	st22 := subtask.SubTask{Name: "CSS", Prices: 700}
	t2.Add(&st21)
	t2.Add(&st22)

	t3 := task.Task{Name: "JS", Responsable: "Frontend Developer", Prices: 1000}

	t4 := task.Task{Name: "API Backend", Responsable: "Backend Developer"}
	st41 := subtask.SubTask{Name: "Authentication", Prices: 100}
	st42 := subtask.SubTask{Name: "DB connection", Prices: 900}
	t4.Add(&st41)
	t4.Add(&st42)

	t5 := task.Task{Name: "Database", Responsable: "DBA", Prices: 1000}

	p.Add(&t1)
	p.Add(&t2)
	p.Add(&t3)
	p.Add(&t4)
	p.Add(&t5)

	fmt.Println(&p)
}
