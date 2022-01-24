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

func mainMenu() (choice int) {
	fmt.Println(`Choose your action:
1) List all entries.
2) Add new entry.
3) Remove an entry by ID.
4) Quit`)
	fmt.Scanln(&choice)
	return choice
}

func printHeader() {
	fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
	fmt.Println("-------------------------------------------------------------------------")
}

func listEntries() {
	var entries []Entry
	f, err := os.Open("phonebook.json")
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
		if (i+1)%20 == 0 && i < len(entries)-1 { // Avoid the prompt for the last row.
			fmt.Print("Press <ENTER> to continue...")
			fmt.Scanln()
			printHeader()
		}
	}
}

func newEntry() {
	var entries []Entry
	f, err := os.Open("phonebook.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}

	var ID uint32

	var id int // This piece of code finds new ID that will be greater than any existing ID.
	for _, e := range entries {
		if id < int(e.ID) {
			id = int(e.ID)
		}
	}
	id = id + 1

	e := Entry{ID: ID}

	fmt.Print("Enter last name: ")
	fmt.Scan(&e.LastName)

	fmt.Print("Enter first name: ")
	fmt.Scan(&e.FirstName)

	fmt.Print("Enter phone number: ")
	fmt.Scan(&e.PhoneNumber)

	e.ID = uint32(id)

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

func removeByID() {
	var entries []Entry
	f, err := os.Open("phonebook.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}

	var ID uint32
	fmt.Println("Please enter the entry ID of the contact you want to remove.")
	fmt.Scan(&ID)
	fmt.Printf("Are you sure that you want to remove the contact with the ID %v?\n", ID)
	var c string
	fmt.Println("Please enter 'Yes' if you are or 'No' if you are not.")
	for {
		fmt.Scan(&c)
		if c == "Yes" {
			for i, e := range entries {
				if ID == e.ID {
					entries = append(entries[:i], entries[i+1:]...)

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

					fmt.Printf("The contact with the ID %v has been removed.\n", ID)
					return
				}
			}
			fmt.Printf("The contact with the ID %v does not exist. Please try again.\n", ID)
			return
		} else if c == "No" {
			break
		} else {
			fmt.Println("ERROR: Please enter 'Yes' or 'No'.")
			continue
		}
	}
}

func main() {
	for {
		choice := mainMenu()
		if choice == 1 {
			listEntries()
		} else if choice == 2 {
			newEntry()
		} else if choice == 3 {
			removeByID()
		} else if choice == 4 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}
