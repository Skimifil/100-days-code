package main

import "fmt"

func main() {
	var golang_course = "Assistir ao curso de Golang da Nana do canal TechWorld With Nana"
	var taskProjectGo = "Fazer um projeto em Go"
	var reaward_after_finish = "Comer uma bolachinha"

	var tasksItems = []string{golang_course, taskProjectGo}

	fmt.Println("### TAREFAS ###")
	fmt.Println("Lista de afazeres")
	for index, task := range tasksItems {
		fmt.Printf("	- %d. %s\n", index+1, task)
	}

	fmt.Println("\n### RECOMPENSAS ###")
	fmt.Println("Ao fim de finalizar todas as trefas, parabéns, segue seu presente por ter terminado o curso")
	fmt.Println(reaward_after_finish)
}
