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
	ID          uint
	FirstName   string
	LastName    string
	PhoneNumber string
}

func mainMenu() (choice int) {
	fmt.Println(`Choose your action:
1) List all entries.
2) Add new entry.
3) Remove an entry by ID.
4) Quit`)
	fmt.Scanln(&choice)
	return choice
}

func openFile() []Entry {
	var entries []Entry
	file, err := os.Open("phonebook.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err := json.NewDecoder(bufio.NewReader(file)).Decode(&entries); err != nil {
		log.Fatal(err)
	}
	return entries
}

func saveFile(entries []Entry) {
	file, err := os.Create("phonebook.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err := json.NewEncoder(file).Encode(entries); err != nil {
		log.Fatal(err)
	}
}

func listEntries() {
	entries := openFile()
	sort.Slice(entries, func(i, j int) bool { return entries[i].LastName < entries[j].LastName })
	fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
	fmt.Println("-------------------------------------------------------------------------")
	for i, e := range entries {
		fmt.Printf("%10d %20s %20s %20s\n", e.ID, e.LastName, e.FirstName, e.PhoneNumber)
		if (i+1)%20 == 0 && i < len(entries)-1 {
			fmt.Print("Press <ENTER> to continue...")
			fmt.Scanln()
			fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
			fmt.Println("-------------------------------------------------------------------------")
		}
	}
}

func addNew() {
	var ID uint
	entries := openFile()
	for i, e := range entries {
		if i != 0 && entries[i].ID > entries[i-1].ID {
			ID = e.ID + 1
		}
	}
	entry := Entry{ID: ID}
	fmt.Print("Enter last name:")
	fmt.Scan(&entry.LastName)
	fmt.Print("Enter first name:")
	fmt.Scan(&entry.FirstName)
	fmt.Print("Enter phone number:")
	fmt.Scan(&entry.PhoneNumber)
	entries = append(entries, entry)
	saveFile(entries)
}

func removeByID() {
	entries := openFile()
	var ID uint
	fmt.Print("Enter ID:")
	fmt.Scan(&ID)
	for i, e := range entries {
		if ID == e.ID {
			entries = append(entries[:i], entries[i+1:]...)
		}
	}
	saveFile(entries)
}

func main() {
	for {
		choice := mainMenu()
		if choice == 1 {
			listEntries()
		} else if choice == 2 {
			addNew()
		} else if choice == 3 {
			removeByID()
		} else if choice == 4 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
			// After addNew() and removeByID() completion, an error "ERR: wrong choice 0" appears, the program continues to work.
		}
	}
}
