package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	for {
		choice := mainMenu()
		if choice == 1 {
			listEntries()
		} else if choice == 2 {
			add()
		} else if choice == 4 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}

func mainMenu() (choice int) {
	fmt.Println(`Choose your action:
1) List entries
2) Add new
3) Remove by ID
4) Quit`)
	fmt.Scanln(&choice)
	return choice
}

type Entry struct {
	ID          uint32
	FirstName   string
	LastName    string
	PhoneNumber string
}

func printHeader() {
	fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
	fmt.Println("-------------------------------------------------------------------------")
}

const (
	phonebookFile = "phonebook.json"
	pageLines     = 20
)

func listEntries() {
	var entries []Entry

	f, err := os.Open(phonebookFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}

	sort.Slice(entries, func(i, j int) bool { return entries[i].LastName < entries[j].LastName })

	printHeader()
	for i, e := range entries {
		fmt.Printf("%10d %20s %20s %20s\n", e.ID, e.LastName, e.FirstName, e.PhoneNumber)
		if (i+1)%20 == 0 && i < len(entries)-1 {
			fmt.Print("Press <ENTER> to continue...")
			fmt.Scanln()
			printHeader()
		}
	}
}

func newID(entries []Entry) (id uint) {
	for _, v := range entries {
		if id < uint(v.ID) {
			id = uint(v.ID)
		}
	}
	return id + 1
}

func add() {
	var Entries []Entry
	var New Entry

	f, err := os.Open(phonebookFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&Entries); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter last name:")
	fmt.Scan(&New.LastName)
	fmt.Println("Enter first name:")
	fmt.Scan(&New.FirstName)
	fmt.Println("Enter phone number of contact:")
	fmt.Scan(&New.PhoneNumber)
	newID(Entries)
	Entries = append(Entries, New)

	f, _ = os.Create("phonebook.json")
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	if err := json.NewEncoder(w).Encode(New); err != nil {
		log.Fatal(err)
	}
}
