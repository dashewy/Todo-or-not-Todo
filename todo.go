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

  // want to be able to taggle on and off 
	
	item := &t.Items[idx]

	if !item.Completed {
	
		if t.Items[idx].Task != "" {
			err := errors.New("Task not in progress")
			fmt.Println(err)
			return err
		}

		item.Completed = true
		completedTime := time.Now()
		item.CompletedAt = &completedTime
		
		item.Task = t.Items[idx].InproTask
		item.InproTask = ""
	}	else {
		item.Completed = false
		item.CompletedAt = nil 
	}
		
	return nil
}

func (t *Todos) edit(idx int, newTitle string) error {

	if err := t.checkIdx(idx); err != nil {
		return err 
	}
	item := &t.Items[idx]

	if item.InproTask != "" {
		item.InproTask = newTitle
	} else {
		item.Task = newTitle 
	}
	return nil 
	
}

func (t *Todos) inpro(idx int) error {

	if err := t.checkIdx(idx); err != nil {
		return err 
	}
	
	item := &t.Items[idx]
	
	if item.Completed {
		err := errors.New("item already complete, reset to change")
		fmt.Println(err)
		return err 
	}

  if item.InproTask == "" {	
		item.InproTask = item.Task
		item.Task = ""
	} else {
		item.Task = item.InproTask
		item.InproTask = ""
	}

	return nil
}

func (t *Todos) getClosed() []string {
	var openItems []Todo
	var closedItems []string
	
	for _, item := range t.Items {

		if item.Completed {
			closedItems = append(closedItems, item.Task)
		} else {
			openItems = append(openItems, item)
		}
	}
	// seperates completed and not completed instead of deleting index 
	t.Items = openItems 
	return closedItems
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
