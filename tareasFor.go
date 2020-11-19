package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) appendTask(tl *task) {

	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printList() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre", tarea.name)
		fmt.Println("Descripcion", tarea.description)
	}
}
func (t *taskList) printListComplete() {
	for _, tarea := range t.tasks {
		if tarea.completed {

			fmt.Println("Nombre", tarea.name)
			fmt.Println("Descripcion", tarea.description)
		}

	}
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markAsCompleted() {
	t.completed = true
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func main() {
	t1 := &task{
		name:        "Terminar Curso de Go",
		description: "Terminar el Curso de Go en Platzi en las proximas dos semanas",
	}
	t2 := &task{
		name:        "Terminar Curso de JavaScript",
		description: "Terminar mi curso de Async/Await en JavaScript",
	}
	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	fmt.Printf("%+v\n", *list.tasks[0])
	fmt.Printf("%+v\n", *list.tasks[1])
	list.tasks[1].markAsCompleted()
	fmt.Printf("%+v\n", *list.tasks[1])
	t3 := &task{
		name:        "Construir mi propio servidor web",
		description: "Construir mi propio web server utilizando Go",
	}
	list.appendTask(t3)

	//Esto es igual a ..........
	for i := 0; i < len(list.tasks); i++ {
		fmt.Println("Index", i, "nombre", list.tasks[i].name)
	}
	//Esto
	for index, tarea := range list.tasks {
		fmt.Println("Index", index, "nombre", tarea.name)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	list.printList()
	list.appendTask(t3)
	list.tasks[0].markAsCompleted()
	fmt.Println("Estas son las tareas completadas")
	list.printListComplete()

	mapaTareas := make(map[string]*taskList)
	mapaTareas["Nestor"] = list

	t4 := &task{
		name:        "Terminar Curso de Go",
		description: "Terminar el Curso de Go en Platzi en las proximas dos semanas",
	}
	t5 := &task{
		name:        "Terminar Curso de JavaScript",
		description: "Terminar mi curso de Async/Await en JavaScript",
	}

	list2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["RICARDO"] = list2
	fmt.Println("Tareas Nestor")
	mapaTareas["Nestor"].printList()
	mapaTareas["RICARDO"].printList()

}
