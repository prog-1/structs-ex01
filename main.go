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
1) List entries
2) Add new
3) Remove by ID
4) Quit`)
	fmt.Scanln(&choice)
	return choice
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

func removeId() {
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

	fmt.Println("Enter ID of the contact that you want to be  removed.")
	fmt.Scan(&id)

	var id uint32
	for a, b := range entries {
		if b.ID == id {
			entries = append(entries[:a], entries[a+1:]...)
			f, _ = os.Create("phonebook.json")
			s := bufio.NewWriter(f)
			defer s.Flush()
			defer f.Close()
			if err := json.NewEncoder(s).Encode(entries); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	for {
		choice := mainMenu()
		if choice == 1 {
			listEntries()
		} else if choice == 2 {
			//not today
		} else if choice == 3 {
			//not today
		} else if choice == 4 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}
