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

	phone := struct {
		name string
		surname string
		number int
	}
	
	gg := []float64{s}
	fmt.Println("Enter name,surname and phone number")
	fmt.Scan(&phone)
	add := append(gg ,phone)
	fmt.Printl(add)

func Remove(){
	file, err := os.Open("phone book.json")
	if err != nil {
  		log.Fatal(err)
	}
	defer file.Close()

	phone := struct {
		name string
		surname string
		number int
	}

	scan := fmt.Fscan(file, &phone)
	scann := []float64{scan}
	var id int
	fmt.Println("Enter ID")
	fmt.Scan(&id)
	for indx := 1,indx < len(scann),indx++{
		if id == indx{
			a := [:indx]float64
			b := [indx:]float64
			fmt.Println(a,b)
		}
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