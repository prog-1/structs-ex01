package main

import "fmt"

func List(){
	file, err := os.Open("phone book.json")
	if err != nil {
  		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
  		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
  		log.Fatal(err)
	}
}

func Add(){
	file, err := os.Open("phone book.json")
	if err != nil {
  		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
  		s.Text()
	} 
	if err := s.Err(); err != nil {
  		log.Fatal(err)
	}

	type phone struct {
		id uint
		name string
		surname string
		number int
	}
	var scan int
	fmt.Println("Enter ID,Name,Surname and Phone number")
	fmt.Scan(&scan)
	scan := []phone{}

	if err := json.NewDecoder(f).Decode(&scan);err != nil{
		fmt.Println(nil,err)
	}
	fmt.Println(scan,err)
}

func Remove(){
	file, err := os.Open("phone book.json")
	if err != nil {
  		log.Fatal(err)
	}
	defer file.Close()

	type phone struct {
		name string
		surname string
		number int
	}

}

func main() {
	var action,answer int
	for {
		fmt.Println("Confirm using main menu,write yes")
		fmt.Scan(&answer)
		if answer == yes{
			for t := 0, t < 20, t++ {
				fmt.Println("Choose your action:
					1) List all entries.
					2) Add new entry.
					3) Remove an entry by ID.")
				fmt.Scan(&action)
				if action == 1{
					List()
				} if else action == 2{
					Add()
				} if else action == 3{
					Remove()
				} else {
					fmt.Println("Wrong number")
				}
			}
		}
	}
}