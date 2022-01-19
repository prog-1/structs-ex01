package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func UI() (action uint) {
	fmt.Println(`To choose your action enter number:
	1) List all entries.
	2) Add new entry.
	3) Remove an entry by ID.
	4) Exit
	5) Leave feedback`)
	fmt.Scan(&action)
	if action > 5 {
		fmt.Println("No such action, please restart program.")
	}
	return action
}

type Entry struct {
	ID          uint
	FirstName   string
	LastName    string
	PhoneNumber string
}

func Sort(entries []Entry) {
	sort.Slice(entries, func(i, j int) bool { return entries[i].LastName < entries[j].LastName })
}

func Header() {
	fmt.Printf("%4s %14s %14s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
	fmt.Println("-------------------------------------------------------------------------")
}

/*func GiveID(entries []Entry) uint {
	i := 1
	for i, j := range entries {
		j.ID = uint(i) + 1 //this prays that i will change it
	}
	return uint(i) + 1
}*/

func List() {
	var entries []Entry
	const phonebookFile = "phonebook.json"
	f, err := os.Open(phonebookFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}
	Sort(entries)
	Header()
	for i, j := range entries {
		fmt.Printf("%4d %14s %14s %20s\n", j.ID, j.LastName, j.FirstName, j.PhoneNumber)
		if i == 19 && i < len(entries) {
			fmt.Println("Press <ENTER> to continue...")
			fmt.Scanln()
			Header()
		}
	}

}

func Add(entries []Entry) {
	// how to encode with user input???
	f, err := os.Open("phonebook.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	//reader := bufio.NewReader(os.Stdin)
	if err := json.NewEncoder(w).Encode(entries); err != nil {
		log.Fatal(err)
	}

}

///func Remove(entries []Entry, s int) []Entry {
// return append(s[:s], slice[s+1:]...)
//}

func Exit() {
	os.Exit(0)
}

func Feedback() {
	var stars int
	fmt.Println("How many stars would you give us on Play Market?")
	fmt.Scan(&stars)
	if stars < 0 {
		fmt.Println("Isnt this too cruel... ＞﹏＜")
	} else {
		fmt.Println("Thanks for sharing your opinion. Its really important ^_^")
	} // i would wanna to make stars struct and convert it to json file to keep all inputed stars but i didnt understand how to encode with user input
}

func main() {
	action := UI()
	if action == 1 {
		List()
	} else if action == 2 {
		fmt.Println("This function isnt done yet")
	} else if action == 3 {
		fmt.Println("This function isnt done yet")
	} else if action == 4 {
		Exit()
	} else if action == 5 {
		Feedback()
	}
}
