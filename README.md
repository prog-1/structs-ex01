# Phone Book

## Description

Create a program which implements a phone book, allowing to list, add and delete phone book entries.
We recommend to store the phone book in files using [JSON](https://en.wikipedia.org/wiki/JSON) format.

### UI / Main menu

User interface (UI) should contain operation choices e.g.
```txt
Choose your action:
1) List all entries.
2) Add new entry.
3) Remove an entry by ID.
```

The database should be read once on startup, and saved every time new entry is added, or an existing
entry is removed.

### Listing entries

Entries should be sorted by last name and listed with 20 elements pagination. It means that every 20
entries a user should confirm to see next 20 (or less). E.g.

```txt
ID              Last Name                First Name                Phone#
-------------------------------------------------------------------------
4               Bar                      Foo                       +33333333333
7               Buba                     Pupkin                    +11111111111
...
3               Foo                      Bar                       +22222222222
Press <ENTER> to continue...
ID              Last Name                First Name                Phone#
-------------------------------------------------------------------------
1               Samcuks                  Jaroslavs                 +41123456789
2               Zaichenkov               Pavel                     +41987654321
```

Note: IDs are auto-incremented for every new entry and aren't necessary sorted.

After the last entry the main menu should be displayed again.

### Adding new entries (optional)

- IDs should be auto-assigned and greater than any existing ID.
- Names, last names and phone numbers should be entered from the keyboard.
- Names and last names could be duplicated.

### Removing entries (optional)

- A user is asked to enter entry ID, which is then removed from the database.
- Optional confirmation messages could be shown.

## JSON

JSON stands for JavaScript Object Notation and is a popular formating allowing to store or transmit data.
Go provides an [encoding/json](https://pkg.go.dev/encoding/json) package that implements encoding and
decoding operations.

In the examples below we use the following structure:

```go
type Entry struct {
  ID          uint32
  FirstName   string
  LastName    string
  PhoneNumber string
}
```

and the following constants:

```go
const phonebookFile = "phonebook.json"
```

### Encoding to a file

```go
entries := []Entry{
  {1, "Jaroslavs", "Samcuks", "+41123456789"},
  {2, "Pavel", "Zaichenkov", "+41987654321"},
}

f := MustCreateFile(phonebookFile)
defer f.Close()
w := bufio.NewWriter(f)
defer w.Flush()
if err := json.NewEncoder(w).Encode(entries); err != nil {
  log.Fatal(err)
}
```

### Decoding from a file

```go
var entries []Entry

f := MustOpenFile(phonebookFile)
defer f.Close()
if err := json.NewDecoder(bufio.NewReader(f)).Decode(&entries); err != nil {
  log.Fatal(err)
}

fmt.Println(entries)
```

## Extra hints

### Main menu

```go
func mainMenu() (choice int) {
	fmt.Println(`Choose your action:
1) List entries
2) Add new
3) Remove by ID
4) Quit`)
	fmt.Scanln(&choice)
	return choice
}
```

### List entries

```go
func listEntries() {
	entries, err := loadEntries()
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].LastName < entries[j].LastName })
	fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
	fmt.Println("-------------------------------------------------------------------------")
	for i, e := range entries {
		fmt.Printf("%10d %20s %20s %20s\n", e.ID, e.LastName, e.FirstName, e.PhoneNumber)
		if (i+1)%pageLines == 0 && i < len(entries)-1 {
			fmt.Print("Press <ENTER> to continue...")
			fmt.Scanln()
			fmt.Printf("%10s %20s %20s %20s\n", "ID", "Last Name", "First Name", "Phone Number")
			fmt.Println("-------------------------------------------------------------------------")
		}
	}
}
```

### Add new entry

```go
func addNew() {
	var e Entry
	fmt.Print("Enter last name: ")
	fmt.Scan(&e.LastName)
	fmt.Print("Enter first name: ")
	fmt.Scan(&e.FirstName)
	fmt.Print("Enter phone number: ")
	fmt.Scan(&e.PhoneNumber)
	entries, _ := loadEntries()
  ...                  // Find largest available ID
	e.ID = largestID + 1 // And increment it.
	if err := saveEntries(append(entries, e)); err != nil {
		fmt.Println("ERR:", err)
	}
}
```

### Main loop

```go
func main() {
	for {
		choice := mainMenu()
		if choice == 1 {
			listEntries()
		} else if choice == 2 {
			addNew()
		} else if choice == 3 {
			removeByID()
		} else if choice == 4 {
			break
		} else {
			fmt.Println("ERR: wrong choice", choice)
		}
	}
}
```
