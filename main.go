package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Entry struct {
	ID          uint32
	FirstName   string
	LastName    string
	PhoneNumber string
}

func Remove() {
	main()
}

func Add() {
	main()
}

func List() {
	var entries []Entry
	f, err := os.Open("phonebook.json")
	log.Fatal(err)
	defer f.Close()
	if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
		log.Fatal(err)
	}
	fmt.Println(entries)
	Menu()
}

func Menu() {
	var mode uint
	fmt.Println(`Choose your action:
	1) List entries
	2) Add new
	3) Remove by ID
	4) Quit`)
	fmt.Scanln(&mode)
	for {
		if mode == 1 {
			List()
		} else if mode == 2 {
			Add()
		} else if mode == 3 {
			Remove()
		} else if mode == 4 {
			return
		} else {
			fmt.Println("Incorrect choice")
			Menu()
		}
	}
}

func main() {
	Menu()
}
