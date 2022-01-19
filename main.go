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
func Remove() {
	fmt.Println("Coming soon...")
	Menu()
}
func Add() {
	fmt.Println("Coming soon...")
	Menu()
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
	fmt.Printf("%s%24s%24s%26s\n", "ID", "Last Name", "First name", "Phone number")
	fmt.Println("----------------------------------------------------------------------------")
	for _, v := range entries {
		if i == 20 {
			fmt.Println("Please press <ENTER> to continue...")
			fmt.Scan()
			fmt.Printf("%s%24s%24s%26s\n", "ID", "Last Name", "First name", "Phone number")
			fmt.Println("----------------------------------------------------------------------------")
			i = 0
		}
		fmt.Printf("%d%25s%24s%26s\n", v.ID, v.LastName, v.FirstName, v.PhoneNumber)
		i++
	}
	fmt.Println("Please press <ENTER> to return to menu...")
	fmt.Scanln(&i)
	Menu()
}
func Menu() {
	fmt.Println(`Please choose you action:
1) List all entries.
2) Add new entry.
3) Remove an entry by ID.
4) Stop programm 
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
		return
	default:
		fmt.Println("Incorect value")
		Menu()
	}

}
func main() {
	Menu()
}
