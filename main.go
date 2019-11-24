package main

import (
	"flag"
	"fmt"
	"os"
)

var tlist []task

type task struct {
	desc     string
	duration int
}

func calculator(a *int, b *int, operand *string) {
	switch *operand {
	case "+":
		result := *a + *b
		fmt.Printf("The calculated result is : %d\n", result)
	case "-":
		result := *a - *b
		fmt.Printf("The calaulted result is : %d\n", result)
	case "*":
		result := (*a) * (*b)
		fmt.Printf("The calculated result is : %d\n", result)
	case "/":
		result := (*a) / (*b)
		fmt.Printf("The calculated result is :%d\n", result)
	case "%":
		result := (*a) % (*b)
		fmt.Printf("The calculated result is :%d\n", result)
	default:
		fmt.Print("Invalid operand... Exiting CLI now...\n")
		os.Exit(1)
	}

}

func addtask(d *string, dur *int) {
	tsk := task{desc: *d, duration: *dur}
	tlist = append(tlist, tsk)
}

func main() {

	a := flag.Int("n1", 0, "First number")
	b := flag.Int("n2", 0, "Second number")
	d := flag.String("task", "", "Task to be performed")
	duration := flag.Int("duration", 0, "Duration of the task")
	operand := flag.String("operand", "", "Operand to be performed")
	flag.Parse()

	if *d == "" && *duration == 0 {
		fmt.Println("You have chosen calculator. Routing to calculator function...")
		calculator(a, b, operand)
	} else {
		fmt.Println("You have chosen to append the tasklist")
		fmt.Printf("Tasks before addition: %d\n", len(tlist))
		addtask(d, duration)
		fmt.Printf("Tasks after addition : %d\n", len(tlist))
	}
	fmt.Print("Hope you had fun\n")

}
