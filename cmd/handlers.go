package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./html/home.page",
		"./html/header.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ts.Execute(w, app.cacheData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (app *application) doNextFit(w http.ResponseWriter, r *http.Request) {
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

	// random split of data
	if processNumber == 1 {
		processes = append(processes, 250, 80, 40, 50)
	}
	if processNumber == 2 {
		processes = append(processes, 170, 140, 35, 115)
	}
	if processNumber == 3 {
		processes = append(processes, 250, 130, 10, 80)
	}

	app.nextFitImpl(processes, processNumber)

	app.home(w, r)
}

func (app *application) nextFitImpl(processes []int, processNumber int) {
	memoryBlocks := app.cacheData.MemoryBlocks
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

func (app *application) LRUHome(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./html/LRU.page",
		"./html/header.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ts.Execute(w, app.lruCacheData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (app *application) doLRU(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	var currentMemoryData []int
	pageFaults := 0
	pageHits := 0

	capacity, err := strconv.Atoi(r.PostForm.Get("capacity"))
	if err != nil {
		panic(err)
	}

	strProcessList := r.PostForm.Get("processList")

	strProcessListSlice := strings.Split(strProcessList, " ")
	processList := toIntArray(strProcessListSlice)

	app.lruImpl(&processList, &currentMemoryData, capacity, &pageFaults, &pageHits)

	strLRUCacheDataSlice := toStringArray(currentMemoryData)
	strLRUCacheData := strings.Join(strLRUCacheDataSlice, ",")

	app.lruCacheData.CurrentMemoryData = strLRUCacheData
	app.lruCacheData.PageFaults = pageFaults
	app.lruCacheData.PageHits = pageHits

	app.LRUHome(w, r)
}

func (app *application) lruImpl(processList *[]int, currentMemoryData *[]int, capacity int, pageFaults *int, pageHits *int) {
	for _, val := range *processList {
		app.infoLog.Printf("Current memory data: %v \n page faults: %d \n page hits: %d \n next val: %d",
			currentMemoryData, pageFaults, pageHits, val)
		if !contains(*currentMemoryData, val) {
			if len(*currentMemoryData) == capacity {
				*currentMemoryData = (*currentMemoryData)[1:]
				*currentMemoryData = append(*currentMemoryData, val)
			} else {
				*currentMemoryData = append(*currentMemoryData, val)
			}
			*pageFaults++
		} else {
			*currentMemoryData = removeElementFromSlice(*currentMemoryData, val)
			*currentMemoryData = append(*currentMemoryData, val)
			*pageHits++
		}
	}
}

func removeElementFromSlice(s []int, value int) []int {
	for i, v := range s {
		if v == value {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	return s
}

func contains(s []int, val int) bool {
	for _, a := range s {
		if a == val {
			return true
		}
	}
	return false
}

func toIntArray(str []string) []int {
	var converted []int
	for _, i := range str {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		converted = append(converted, j)
	}
	return converted
}

func toStringArray(int []int) []string {
	var converted []string
	for _, i := range int {
		j := strconv.Itoa(i)
		converted = append(converted, j)
	}
	return converted
}
