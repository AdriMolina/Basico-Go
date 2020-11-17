package main

import (
	"fmt"
)

type taskList struct {
	task []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.task = append(t.task, tl)
}

//Definicion de un struc con propiedades
type task struct {
	nombre      string
	descripcion string
	completado  bool
}

//Cambiar propiedades de compñetado
func (t task) marcarCompleta() {
	t.completado = true
}

func (t task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	// Inicialización de un struct tipo task
	t := task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go de platzi en esta semana",
	}

	t1 := task{
		nombre:      "Completar mi curso de Segurdad",
		descripcion: "Completar mi curso de Go de platzi en esta semana",
	}
	lista := &taskList{
		task: []*task{
			t, t1,
		},
	}
	fmt.Println(lista.task[0])
	lista.agregarALista(t3)
	fmt.Println(len(lista.task))

}
