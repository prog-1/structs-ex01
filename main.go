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

func Sort(rec []Record) {
	for j := range rec {
		for i := 0; i < len(rec)-1-j; i++ {
			tmp := []string{rec[i].LastName, rec[i+1].LastName}
			sort.Strings(tmp)
			if tmp[0] != rec[i].LastName {
				rec[i], rec[i+1] = rec[i+1], rec[i]
			}
		}
	}
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
	Sort(Records)
	fmt.Println("ID              Last Name                First Name                Phone#")
	fmt.Println("-------------------------------------------------------------------------")
	i := 0
	for _, rec := range Records {
		if i == 20 {
			fmt.Println("Please press <ENTER> to continue...")
			fmt.Scanln()
		}
		fmt.Printf("%-16v%-25v%-26v%v\n", rec.ID, rec.LastName, rec.FirstName, rec.PhoneN)
		i++
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please choose you action:\n1) List all entries.\n2) Add new entry.(under construction)\n3) Remove an entry by ID.(under construction)")
	fmt.Print(">>>")
	scanner.Scan()
	if scanner.Text() == "1" {
		ListEntries()
	}

}
