package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type Entry struct {
	ID          uint32
	FirstName   string
	LastName    string
	PhoneNumber string
}

func Remove() {
	var entries []Entry
	const phonebookFile = "phonebook.json"
	var id uint32
	var choise int
	f, err := os.Open(phonebookFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter ID: ")
	fmt.Scan(&id)
	for i, v := range entries {
		if v.ID == id {
			fmt.Printf("Do you really want to deleat %q %q. Enter <1> to confirm deletion: ", v.FirstName, v.LastName)
			fmt.Scan(&choise)
			if choise == 1 {
				entries = append(entries[:i], entries[i+1:]...)
				file, err := os.Create(phonebookFile)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				w := bufio.NewWriter(file)
				defer w.Flush()
				if err := json.NewEncoder(w).Encode(entries); err != nil {
					log.Fatal(err)
				}

			}
			return
		}
	}
}

func Add() {
	var entries []Entry
	var add Entry
	var id uint32

	const phonebookFile = "phonebook.json"
	scanner := bufio.NewScanner(os.Stdin)
	f, err := os.Open(phonebookFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}

	for _, v := range entries {
		if id < v.ID {
			id = v.ID
		}

	}
	add.ID = id + 1
	fmt.Print("Enter last name: ")
	scanner.Scan()
	add.LastName = scanner.Text()
	fmt.Print("Enter first name: ")
	scanner.Scan()
	add.FirstName = scanner.Text()
	fmt.Print("Enter phone number: ")
	scanner.Scan()
	add.PhoneNumber = scanner.Text()
	entries = append(entries, add)
	file, err := os.Create(phonebookFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	defer w.Flush()
	if err := json.NewEncoder(w).Encode(entries); err != nil {
		log.Fatal(err)
	}
}

func PrintHeader() {
	fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
	fmt.Println("-------------------------------------------------------------------------")
}

func List() {
	const linesPerPage = 20
	var entries []Entry
	const phonebookFile = "phonebook.json"
	f, err := os.Open(phonebookFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].LastName < entries[j].LastName })
	PrintHeader()
	for i, e := range entries {
		fmt.Printf("%10d %20s %20s %20s\n", e.ID, e.LastName, e.FirstName, e.PhoneNumber)
		if (i+1)%linesPerPage == 0 && i < len(entries)-1 {
			fmt.Print("Press <ENTER> to continue...")
			fmt.Scanln()
			PrintHeader()
		}
	}
}

func Menu() (mode int) {
	fmt.Println(`Choose your action:
	1) List entries
	2) Add new
	3) Remove by ID
	4) Quit`)
	fmt.Scanln(&mode)
	return mode
}

func main() {
	for {
		choice := Menu()
		if choice == 1 {
			List()
		} else if choice == 2 {
			Add()
		} else if choice == 3 {
			Remove()
		} else if choice == 4 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}
