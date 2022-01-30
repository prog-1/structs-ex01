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
		} else if choice == 3 {
			remove()
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

func add() {
	var entries []Entry
	var new Entry

	entries, err := loadPb()
	if err != nil {
		fmt.Println("Err")
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter last name:")
	scanner.Scan()
	new.LastName = scanner.Text()
	fmt.Println("Enter first name:")
	scanner.Scan()
	new.FirstName = scanner.Text()
	fmt.Println("Enter phone number:")
	scanner.Scan()
	new.PhoneNumber = scanner.Text()

	for _, i := range entries {
		if new.ID != i.ID {
			new.ID = i.ID
		}

	}

	new.ID = new.ID + 1
	entries = append(entries, new)

	if err := savePb(entries); err != nil {
		fmt.Println("Err")
	}
}

func remove() {
	var id uint32
	i := 0
	entries, err := loadPb()
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println("Enter ID:")
	fmt.Scan(&id)

	for _, e := range entries {
		if e.ID != id {
			entries[i] = e
			i++
		}
	}
	entries = entries[:i]
	if err := savePb(entries); err != nil {
		fmt.Println("Err")
	}
}
func loadPb() ([]Entry, error) {
	f, err := os.Open(phonebookFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var entries []Entry
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		return nil, err
	}
	return entries, err
}

func savePb(entries []Entry) error {
	f, err := os.Create("phonebook.json")
	if err != nil {
		return err
	}
	return json.NewEncoder(f).Encode(entries)
}
