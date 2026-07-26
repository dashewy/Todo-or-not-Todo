package main 

import (
	"fmt"
	"time"
	"errors"
	"os"
	"strconv"
	"github.com/aquasecurity/table"
)

type Todo struct {
	Task string
	InproTask string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time 
}

type Todos struct {
	Items []Todo
}

func (t *Todos) checkIdx(idx int) error {
	
	if idx < 0 || idx >= len(t.Items) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err 
	}

	return nil 
}

func (t *Todos) add(task string) {
	
	todo := Todo{ 
		Task: task,
		InproTask: "",
		Completed: false,
		CreatedAt: time.Now() ,
		CompletedAt: nil,
	}

	t.Items = append(t.Items, todo)

}

func (t *Todos) del(idx int) error {

	if err := t.checkIdx(idx); err != nil {
		return err
	}
	
	t.Items = append(t.Items[:idx], t.Items[idx+1:]...)
	return nil 
}

func (t *Todos) toggle(idx int) error {

	if err := t.checkIdx(idx); err != nil {
		return err 
	}
	if t.Items[idx].Task != "" {
		err := errors.New("Task not in progress")
		fmt.Println(err)
		return err
	}
  // want to be able to taggle on and off 
	
	status := t.Items[idx].Completed

	if !status {
		t.Items[idx].Completed = true
		completedTime := time.Now()
		t.Items[idx].CompletedAt = &completedTime
		t.Items[idx].Task = t.Items[idx].InproTask
		t.Items[idx].InproTask = ""
	}	else {
		t.Items[idx].Completed = false
		t.Items[idx].CompletedAt = nil 
	}
		
	return nil
}

func (t *Todos) edit(idx int, newTitle string) error {

	if err := t.checkIdx(idx); err != nil {
		return err 
	}
	
	t.Items[idx].Task = newTitle 

	return nil 
	
}

func (t *Todos) inpro(idx int) error {

	if err := t.checkIdx(idx); err != nil {
		return err 
	}
	
  if t.Items[idx].InproTask == "" {	
		t.Items[idx].InproTask = t.Items[idx].Task
		t.Items[idx].Task = ""
	} else {
		t.Items[idx].Task = t.Items[idx].InproTask
		t.Items[idx].InproTask = ""
	}

	return nil
}

func (t *Todos) printTab() {
	
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Completed", "Task", "In progress", "Created At", "Completed At")

	for idx, todo := range t.Items {
		compVal := "❌"
		completedAt := ""
	

		if todo.Completed {
			compVal = "✅"
		
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(idx), compVal, todo.Task, todo.InproTask, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
