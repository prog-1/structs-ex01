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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func Remove() string {
	const phonebookFile = "phonebook.json"
	var entries []Entry
	f, err := os.Open(phonebookFile)
	check(err)
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}
	id := -1
	fmt.Println("Enter ID:")
	fmt.Scanln(&id)
	for i, v := range entries {
		if v.ID == uint32(id) {
			entries = append(entries[:i], entries[i+1:]...)
			f, _ = os.Create("phonebook.json")
			defer f.Close()
			w := bufio.NewWriter(f)
			defer w.Flush()
			if err := json.NewEncoder(w).Encode(entries); err != nil {
				log.Fatal(err)
			}
			return ""
		}
	}
	return "ID does not exist"
}
func Add() {
	const phonebookFile = "phonebook.json"
	var entries []Entry
	var new Entry
	f, err := os.Open(phonebookFile)
	check(err)
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
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
	var id uint32
	for _, v := range entries {
		if id < v.ID {
			id = v.ID
		}

	}
	new.ID = id + 1
	entries = append(entries, new)
	f, err = os.Create("phonebook.json")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	if err := json.NewEncoder(w).Encode(entries); err != nil {
		log.Fatal(err)
	}

}

func List() {
	const phonebookFile = "phonebook.json"
	var entries []Entry

	var i int
	f, err := os.Open(phonebookFile)
	check(err)
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].LastName < entries[j].LastName })
	fmt.Printf("%5s%24s%24s%26s\n", "ID", "Last Name", "First name", "Phone number")
	fmt.Println("----------------------------------------------------------------------------")
	for _, v := range entries {
		if i == 20 {
			fmt.Println("Please press <ENTER> to continue...")
			fmt.Scan()
			fmt.Printf("%5s%24s%24s%26s\n", "ID", "Last Name", "First name", "Phone number")
			fmt.Println("----------------------------------------------------------------------------")
			i = 0
		}
		fmt.Printf("%5d%25s%24s%26s\n", v.ID, v.LastName, v.FirstName, v.PhoneNumber)
		i++
	}
	fmt.Println("Please press <ENTER> to return to menu...")
	fmt.Scanln(&i)
}
func Menu() int {
	fmt.Println(`Please choose your action:
1) List all entries.
2) Add new entry.
3) Remove an entry by ID.
4) Stop program
Enter number(1-4):`)
	var a uint
	fmt.Scan(&a)
	switch a {
	case 1:
		List()
	case 2:
		Add()
	case 3:
		Remove()
	case 4:
		return 4
	default:
		fmt.Println("Incorect value")

	}
	return 0
}
func main() {
	for {
		a := Menu()
		if a == 4 {
			return
		}
	}
}
