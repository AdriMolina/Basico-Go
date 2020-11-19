package main

import "fmt"

type task1List1 struct {
	tasks []*task1
}

func (t *task1List1) appendtask1(tl *task1) {

	t.tasks = append(t.tasks, tl)
}

func (t *task1List1) removetask1(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task11 struct {
	name        string
	description string
	completed   bool
}

func (t *task1) markAsCompleted() {
	t.completed = true
}

func (t *task1) updateName(name string) {
	t.name = name
}

func (t *task1) updateDescription(description string) {
	t.description = description
}

func main() {
	t1 := &task1{
		name:        "Terminar Curso de Go",
		description: "Terminar el Curso de Go en Platzi en las proximas dos semanas",
	}
	t2 := &task1{
		name:        "Terminar Curso de JavaScript",
		description: "Terminar mi curso de Async/Await en JavaScript",
	}
	list := &task1List{
		tasks: []*task1{
			t1, t2,
		},
	}
	fmt.Printf("%+v\n", *list.tasks[0])
	fmt.Printf("%+v\n", *list.tasks[1])
	list.tasks[1].markAsCompleted()
	fmt.Printf("%+v\n", *list.tasks[1])
	t3 := &task1{
		name:        "Construir mi propio servidor web",
		description: "Construir mi propio web server utilizando Go",
	}
	list.appendtask1(t3)
	fmt.Printf("%+v\n", *list.tasks[2])
	list.removetask1(1)
	fmt.Printf("%+v\n", *list.tasks[1])
}
