package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type Record struct {
	ID        uint
	FirstName string
	LastName  string
	PhoneN    string
}

func removeEntry() (errr string) {
	var Records []Record
	f, err := os.Open("Data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&Records); err != nil {
		log.Fatal(err)
	}
	id := -1
	fmt.Println("ID:")
	fmt.Scanln(&id)
	for i, v := range Records {
		if v.ID == uint(id) {
			Records = append(Records[:i], Records[i+1:]...)
			f.Close()
			f, _ = os.Create("Data.json")
			defer f.Close()
			w := bufio.NewWriter(f)
			defer w.Flush()
			if err := json.NewEncoder(w).Encode(Records); err != nil {
				log.Fatal(err)
			}
			return ""
		}
	}
	return "ID not found"
}

func getNextID(rec []Record) uint {
	id := -1
	for _, v := range rec {
		if id < int(v.ID) {
			id = int(v.ID)
		}
	}
	return uint(id + 1)
}

func ListEntries() {
	var Records []Record
	f, err := os.Open("Data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&Records); err != nil {
		log.Fatal(err)
	}
	sort.Slice(Records, func(i, j int) bool { return Records[i].LastName < Records[j].LastName })
	fmt.Println("ID              Last Name                First Name                Phone#")
	fmt.Println("-------------------------------------------------------------------------")
	i := 0
	for _, rec := range Records {
		if i == 20 {
			fmt.Println("Please press <ENTER> to continue...")
			fmt.Scanln()
			fmt.Println("ID              Last Name                First Name                Phone#")
			fmt.Println("-------------------------------------------------------------------------")
		}
		fmt.Printf("%-16v%-25v%-26v%v\n", rec.ID, rec.LastName, rec.FirstName, rec.PhoneN)
		i++
	}
}

func addNewEntry(scanner bufio.Scanner) {
	var Records []Record
	f, err := os.Open("Data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&Records); err != nil {
		log.Fatal(err)
	}
	var Entry Record
	fmt.Print("Last Name: ")
	scanner.Scan()
	Entry.LastName = scanner.Text()
	fmt.Print("First Name: ")
	scanner.Scan()
	Entry.FirstName = scanner.Text()
	fmt.Print("Phone#: ")
	scanner.Scan()
	Entry.PhoneN = scanner.Text()
	Entry.ID = getNextID(Records)
	Records = append(Records, Entry)
	f.Close()
	f, _ = os.Create("Data.json")
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	if err := json.NewEncoder(w).Encode(Records); err != nil {
		log.Fatal(err)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Please choose you action:\n1) List all entries.\n2) Add new entry.\n3) Remove an entry by ID.(under construction)\n4) Quit")
		fmt.Print(">>>")
		scanner.Scan()
		if scanner.Text() == "1" {
			ListEntries()
		}
		if scanner.Text() == "2" {
			addNewEntry(*scanner)
		}
		if scanner.Text() == "3" {
			if err := removeEntry(); err != "" {
				fmt.Println(err)
			}

		}
		if scanner.Text() == "4" {
			break
		}
	}
}
