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

func newEntry() {
	const phonebookFile = "phonebook.json"
	var entries []Entry
	var ID uint32
	e := Entry{ID: ID}
	f, err := os.Open(phonebookFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter last name:")
	scanner.Scan()
	e.LastName = scanner.Text()
	fmt.Println("Enter first name:")
	scanner.Scan()
	e.FirstName = scanner.Text()
	fmt.Println("Enter phone number:")
	scanner.Scan()
	e.PhoneNumber = scanner.Text()

	for _, i := range entries {
		if e.ID != i.ID {
			e.ID = i.ID
		}

	}
	e.ID = e.ID + 1
	entries = append(entries, e)
	f, err = os.Create("phonebook.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	if err := json.NewEncoder(w).Encode(entries); err != nil {
		log.Fatal(err)
	}

}

func printHeader() {
	fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
	fmt.Println("-------------------------------------------------------------------------")
}

func listEntries() {
	const phonebookFile = "phonebook.json"
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

func removeById() {
	const phonebookFile = "phonebook.json"
	var entries []Entry
	f, err := os.Open(phonebookFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter ID:")
	var id uint32
	fmt.Scanln(&id)
	for j, i := range entries {
		if i.ID == id {
			entries = append(entries[:j], entries[j+1:]...)
			f, _ = os.Create("phonebook.json")
			defer f.Close()
			en := bufio.NewWriter(f)
			defer en.Flush()
			if err := json.NewEncoder(en).Encode(entries); err != nil {
				log.Fatal(err)
			}
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

func main() {
	for {
		choice := mainMenu()
		if choice == 1 {
			listEntries()
		} else if choice == 2 {
			newEntry()
		} else if choice == 3 {
			removeById()
		} else if choice == 4 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}
