//////////////////////////////////////////////////////
//	Project:	Godo
//	Author:		Erick Muup
//	File:		Todo.go
//	Descr:		Contains struct definition for usage
//				in the application
//////////////////////////////////////////////////////

package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

// A single todo item.
type Todo struct {
	Task      string    `"task":`
	Date      time.Time `"Date":`
	Done      bool      `"Done":`
	Completed time.Time `"Completed":`
}

type Todos []Todo // A typedef for an array of todos

func (t *Todos) Add(task string) {
	// Append a task to the list of items
	todo := Todo{
		Task:      task,
		Date:      time.Now(),
		Done:      false,
		Completed: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	// complete a task
	ls := *t

	if index <= 0 || index >= len(ls) {
		return errors.New("invalid index")
	}

	(ls)[index-1].Completed = time.Now()
	(ls)[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {

	// Delete an item
	ls := *t

	if index <= 0 || index >= len(*t) {
		return errors.New("invalid index")
	}

	*t = append((ls)[:index-1], (ls)[index:]...)

	return nil
}

func (t *Todos) CountPending() int {
	total := 0

	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}

// file io

func (t *Todos) Load(filename string) error {
	// load items from disk
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)

	if err != nil {
		return err
	}

	return nil

}

func (t *Todos) Store(filename string) error {
	// store the todos into disk

	content, err := json.MarshalIndent(t, "", " ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, content, 0644)

	if err != nil {
		return nil
	}

	return nil

}
