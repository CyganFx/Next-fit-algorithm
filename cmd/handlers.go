package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./html/home.page",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ts.Execute(w, app.CacheTable)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (app *application) doTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	var processes []int

	processNumber, err := strconv.Atoi(r.PostForm.Get("task"))
	if err != nil {
		fmt.Println("string to int convertion problem")
		return
	}

	if processNumber == 1 {
		processes = append(processes, 250, 80, 40, 50)
	}
	if processNumber == 2 {
		processes = append(processes, 170, 140, 35, 115)
	}
	if processNumber == 3 {
		processes = append(processes, 250, 130, 10, 80)
	}

	app.nextFit(processes, processNumber)

	files := []string{
		"./html/home.page",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ts.Execute(w, app.CacheTable)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (app *application) nextFit(processes []int, processNumber int) {
	memoryBlocks := app.CacheTable.MemoryBlocks
	memoryBlockLength := len(memoryBlocks)
	processesLength := len(processes)
	pointer := 0

	for i := 0; i < processesLength; i++ {
		for pointer < memoryBlockLength {
			if memoryBlocks[pointer].FreeMemoryLeft >= processes[i] {
				tempForLog := memoryBlocks[pointer].FreeMemoryLeft

				memoryBlocks[pointer].FreeMemoryLeft -= processes[i]

				app.infoLog.Printf("Memory block #%d: Was %d KB, Taken %d KB by process number #%d, memory left: %d",
					memoryBlocks[pointer].Id, tempForLog, processes[i], processNumber, memoryBlocks[pointer].FreeMemoryLeft)
				break
			}
			pointer = (pointer + 1) % memoryBlockLength
		}
	}
}
