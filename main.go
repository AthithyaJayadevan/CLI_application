package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type tasks []task

var dbfile = "mydb.json"

var scanner *bufio.Scanner

type task struct {
	Name     string `json:"name"`
	Duration int    `json : "duration"`
	Status   string `json:"status"`
}

func addtask(name string, duration int) {
	newtask := &task{name, duration, "Not done"}
	db, err := os.Open(dbfile)

	bytevalue, _ := ioutil.ReadAll(db)
	if err != nil {
		fmt.Printf("There was an error in opening the JSON DB : %v", err)
	}

	var alltasks tasks
	json.Unmarshal(bytevalue, &alltasks)

	if len(alltasks) == 0 {
		fmt.Print("Initial stage entry. JSON file empty. Adding corresponding task to JSON file")
		alltasks = append(alltasks, *newtask)
		file, _ := json.MarshalIndent(alltasks, "", " ")
		_ = ioutil.WriteFile("mydb.json", file, 0644)
	} else {
		alltasks = append(alltasks, *newtask)
		file, _ := json.MarshalIndent(alltasks, "", " ")
		_ = ioutil.WriteFile("mydb.json", file, 0644)
	}

	fmt.Printf("Task successfully added to DB")

}

func completetask(name string) {
	var alltasks tasks
	db, err := os.Open(dbfile)
	if err != nil {
		fmt.Printf("There was an error in opening the JSON DB : %v", err)
		os.Exit(1)
	}
	bytevalue, _ := ioutil.ReadAll(db)
	json.Unmarshal(bytevalue, &alltasks)

	for _, item := range alltasks {
		if item.Name == name {
			item.Status = "Done"
		}
	}
	file, _ := json.MarshalIndent(alltasks, "", " ")
	_ = ioutil.WriteFile("mydb.json", file, 0644)

	fmt.Printf("Task %s successfullt marked as complete", name)
}

func deletetask(name string) {
	var alltasks tasks
	db, err := os.Open(dbfile)
	if err != nil {
		fmt.Printf("There was an error in opening the JSON DB : %v", err)
		os.Exit(1)
	}
	bytevalue, _ := ioutil.ReadAll(db)
	json.Unmarshal(bytevalue, &alltasks)
	for i, item := range alltasks {
		if item.Name == name {
			alltasks = append(alltasks[:i], alltasks[i+1:]...)
		}
	}
	file, _ := json.MarshalIndent(alltasks, "", " ")
	_ = ioutil.WriteFile("mydb.json", file, 0644)
	fmt.Printf("Task %s successfully removed from CLI list", name)
}

func listtasks() {
	var alltasks tasks
	db, err := os.Open(dbfile)
	if err != nil {
		fmt.Printf("There was an error in opening the JSON DB : %v", err)
		os.Exit(1)
	}
	bytevalue, _ := ioutil.ReadAll(db)
	json.Unmarshal(bytevalue, &alltasks)
	if len(alltasks) == 0 {
		fmt.Printf("Oops, Sorry. You have not added any tasks yet !!!\n")
		os.Exit(1)
	}
	for _, item := range alltasks {
		fmt.Printf("TASK:%s, DURATION: %d, STATUS: %s\n", item.Name, item.Duration, item.Status)
	}
	fmt.Printf("All tasks listed out...")
}

func networkinfo(cmd string) {
	if cmd == "networks_around" {
		command := exec.Command("netsh wlan show networks")
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			fmt.Printf("An unexpected error occured :%v", err)
		}
	}

}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("cli is a simple enough task manager that helps you to keep track of tasks")
		fmt.Printf("Available commands are : add, complete, delete, list")
		os.Exit(1)
	}
	cmd := flag.Arg(0)
	switch cmd {
	case "list":
		listtasks()
	case "add":
		ntask := strings.Join(flag.Args()[1:], " ")
		splittaks := strings.Split(ntask, " ")
		taskname := strings.Join(splittaks[:len(splittaks)-1], " ")
		nduration, err := strconv.Atoi(splittaks[len(splittaks)-1])
		if err != nil {
			fmt.Printf("there is something wrong with the input : %v", err)
		}
		addtask(taskname, nduration)
	case "delete":
		ntask := flag.Arg(1)
		deletetask(ntask)
	case "do":
		ntask := flag.Arg(1)
		completetask(ntask)

	case "network":
		networkcmd := flag.Arg(1)
		networkinfo(networkcmd)

	}

}
