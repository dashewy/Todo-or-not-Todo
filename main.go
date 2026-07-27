package main 

import (
	"fmt"
	"flag"
	"strings"
	"os"
	"strconv"
)

type Args struct {
	Add string
	Del int 
	Edit string 
	Toggle int
	Inpro int
	List bool
	Close bool
}

func NewArg() *Args {

	arg := Args{}

	flag.StringVar(&arg.Add, "add", "", "Add a new task")
	flag.StringVar(&arg.Edit, "edit", "", "Edit task title use id:newTitle")
	flag.IntVar(&arg.Del, "delete", -1, "Delete a task by index")
	flag.IntVar(&arg.Toggle, "toggle", -1, "Update task to completed")
	flag.IntVar(&arg.Inpro, "inpro", -1, "Update task to be inprogress")
	flag.BoolVar(&arg.List, "list", false, "show the current todos")
	flag.BoolVar(&arg.Close, "close", false, "save all uncompleted task")

	flag.Parse()

	return &arg
}

func (arg *Args) Execute(t *Todos) {

	switch {
	case arg.List:
		t.printTab()
	case arg.Add != "":
		t.add(arg.Add)
	case arg.Edit != "":
		split := strings.SplitN(arg.Edit, ":", 2)
		if len(split) != 2 {
			fmt.Println("Please use the format id:newTitle")
			os.Exit(1)
		}

		idx, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Invalid index")
			os.Exit(1)
		}

		t.edit(idx, split[1])

	case arg.Toggle != -1:
		t.toggle(arg.Toggle)
	case arg.Del != -1:
		t.del(arg.Del)
	case arg.Inpro != -1:
		t.inpro(arg.Inpro)
	case arg.Close:
		fmt.Println("closing now")
	default:
		fmt.Println("Ivalid Cmd")
	}
}

func main() {
	storage := NewStorage[Todos]("todos.json")
	archived := NewStorage[[]string]("history.json")
	args := NewArg()

	todos, err := storage.Load()
	if err != nil {
		todos = Todos{}
	}

	args.Execute(&todos) 

	if args.Close {
		closed := todos.getClosed()
		
		if len(closed) > 0 {
			
			history, err := archived.Load()
			if err != nil {
				history = []string{}
			}
			// reverse args to prepend 
			history = append(closed, history...)
			// may make config so want to have this adjustable
			maxItems := 25
			if len(history) > maxItems {
				history = history[:maxItems]
			}
			
			if err := archived.Save(history); err != nil {
				fmt.Println("Error saving history:", err)
			} 
			fmt.Println("Items saved in history")
		} else {
			fmt.Println("No completed items to be saved")
		}

	}
	
	if err:= storage.Save(todos); err != nil {
		fmt.Println("Error saving:", err)
	}

}
