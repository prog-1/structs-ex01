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
	if action > 5 || action == 0 {
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
	fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
	fmt.Println("-------------------------------------------------------------------------")
}

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
		fmt.Printf("%10d %20s %20s %20s\n", j.ID, j.LastName, j.FirstName, j.PhoneNumber)
		if (i+1)%20 == 0 && i < len(entries) {
			fmt.Println(" ")
			fmt.Println("Press <ENTER> to continue...")
			fmt.Println(" ")
			fmt.Scanln()
			Header()
		}
	}
}

func GiveID(entries []Entry) (i uint) {
	i = uint(len(entries)) + 1 //Thats not ideal and will work wrong if the last contact will be deleted
	return i
}

func Add() {
	// how to encode with user input??? Tutorial by me (version 1)
	var entries []Entry // make variable where will be all data

	f, err := os.Open("phonebook.json") // open and encode file with already been information and put it in variable you made in previous step
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}

	var newentries Entry // make new variable with type of your struct, but be careful - it need to be NOT SLICE OF STRUCT, JUST STRUCT, CAUSE WE MAKING ONLY ONE NEW CONTACT, NOT 5 IN ONE TIME

	scanner := bufio.NewScanner(os.Stdin) // Scan information you need
	fmt.Println("Enter last name:")       // better not use bufio reader, it add \n\r to text and you ll need 2 variables for string and error
	scanner.Scan()
	scanner.Scan()
	newentries.LastName = scanner.Text()
	fmt.Println("Enter first name:")
	scanner.Scan()
	newentries.FirstName = scanner.Text()
	fmt.Println("Enter phone number:")
	scanner.Scan()
	newentries.PhoneNumber = scanner.Text()
	newentries.ID = GiveID(entries)

	entries = append(entries, newentries) // add scanned data to decoded earlier data

	f, err = os.Create("phonebook.json") // and encode it back
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

func Remove() {
	var entries []Entry                 //var
	f, err := os.Open("phonebook.json") // opening and decoding
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}
	var noid, nosure uint // actions, work with data
	fmt.Println("Please, enter ID of contact you want to delete:")
	fmt.Scan(&noid)
	for i, j := range entries {
		if j.ID == noid {
			fmt.Println(`Are you sure? 
*the company "Ltd Littleunidragon" is not responsible for "accidentally" permanently deleted contacts
1) Yes
2) No `)
			fmt.Scan(&nosure)
			if nosure == 1 {
				entries = append(entries[:i], entries[i+1:]...)
				fmt.Println("Contact with ID ", noid, " was removed")
				break
			} else {
				break
			}
		}
	}
	f, err = os.Create("phonebook.json") // and encode it back
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

func Exit() {
	os.Exit(0)
}

type Feedbackstruct struct {
	Stars int
	Text  string
}

func Feedback() { // i wanned to make stars struct and convert it to json file to keep.
	var feedbacks []Feedbackstruct //variable

	f, err := os.Open("feedback.json") // open and encode REMEMBER RENAME FILE THAT WILL BE OPENED
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&feedbacks); err != nil {
		log.Fatal(err)
	}
	var newfeedback Feedbackstruct //just var

	fmt.Println("You want to leave a feedback? You are so great user! ^0^") // scan and action
	fmt.Print("Stars: ")
	fmt.Scan(&newfeedback.Stars)

	scanner := bufio.NewScanner(os.Stdin) // Scan information you need
	fmt.Print("Commentary: ")
	scanner.Scan()
	scanner.Scan()
	newfeedback.Text = scanner.Text()

	if newfeedback.Stars < 0 {
		fmt.Println("Isnt this too cruel... ＞﹏＜")
	} else {
		fmt.Println("Thanks for sharing your opinion. It is really important for us. ^_^")
	}
	feedbacks = append(feedbacks, newfeedback) // add new
	f, err = os.Create("feedback.json")        // encode FILE RENAIMING
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	if err := json.NewEncoder(w).Encode(feedbacks); err != nil {
		log.Fatal(err)
	}
}

func main() {
	for {
		action := UI()
		if action == 1 {
			List()
		} else if action == 2 {
			Add()
		} else if action == 3 {
			Remove()
		} else if action == 4 {
			Exit()
		} else if action == 5 {
			Feedback()
		}
	}
}
