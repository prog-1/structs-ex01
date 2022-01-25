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

const phonebookFile = "phonebook.json"

func Sort(entries []Entry) {
	sort.Slice(entries, func(i, j int) bool { return entries[i].LastName < entries[j].LastName })
}

func Header() {
	fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
	fmt.Println("-------------------------------------------------------------------------")
}
func ReadPhonebook() (entries []Entry, err error) {
	f, err := os.Open(phonebookFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	if err = json.NewDecoder(r).Decode(&entries); err != nil {
		return nil, err
	}
	return entries, nil
}
func ReadFeedback() (feedbacks []Feedbackstruct, err error) {
	f, err := os.Open("feedback.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	if err = json.NewDecoder(r).Decode(&feedbacks); err != nil {
		return nil, err
	}
	return feedbacks, nil
}
func List() {
	var entries []Entry
	entries, _ = ReadPhonebook()
	Sort(entries)
	Header()
	for i, j := range entries {
		fmt.Printf("%10d %20s %20s %20s\n", j.ID, j.LastName, j.FirstName, j.PhoneNumber)
		if (i+1)%20 == 0 && i < len(entries) {
			fmt.Println("Press <ENTER> to continue...")
			fmt.Scanln()
			Header()
		}
	}
}

func GiveID(entries []Entry) (i uint) {
	i = uint(len(entries)) + 1
	return i
}

func Add(scanner *bufio.Scanner) {
	// how to encode with user input??? Tutorial by me (version 1)
	var entries []Entry          // make variable where will be all data
	entries, _ = ReadPhonebook() // open and encode file with already been information and put it in variable you made in previous step

	var newentries Entry // make new variable with type of your struct, but be careful - it need to be NOT SLICE OF STRUCT, JUST STRUCT, CAUSE WE MAKING ONLY ONE NEW CONTACT, NOT 5 IN ONE TIME

	fmt.Println("Enter last name:") // Scan information you need
	scanner.Scan()                  // better not use bufio reader, it adds \r\n to text and you ll need 2 variables for string and error
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

	f, err := os.Create("phonebook.json") // and encode it back
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
	var entries []Entry          //var
	entries, _ = ReadPhonebook() // opening and decoding

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
			}
			break
		}
	}
	f, err := os.Create("phonebook.json") // and encode it back
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
	Text  string `json:"Text,omitempty"`
}

func Feedback(scanner *bufio.Scanner) { // i wanned to make stars struct and convert it to json file to keep.
	var feedbacks []Feedbackstruct //variable
	feedbacks, _ = ReadFeedback()  // open and encode REMEMBER RENAME FILE THAT WILL BE OPENED

	var newfeedback Feedbackstruct //just var

	fmt.Println("You want to leave a feedback? You are so great user! ^0^") // scan and action
	fmt.Print("Stars: ")
	fmt.Scan(&newfeedback.Stars)

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
	f, err := os.Create("feedback.json")       // encode FILE RENAIMING
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		action := UI()
		if action == 1 {
			List()
		} else if action == 2 {
			Add(scanner)
		} else if action == 3 {
			Remove()
		} else if action == 4 {
			Exit()
		} else if action == 5 {
			Feedback(scanner)
		}
	}
}
