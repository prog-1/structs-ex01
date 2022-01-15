# Phone Book

## Description

Create a program which implements a phone book, allowing to list and add phone book entries.
We recommend to store the phone book in files using [JSON](https://en.wikipedia.org/wiki/JSON) format.

TODO(yarcat): Explain UI.

## JSON

JSON stands for JavaScript Object Notation and is a popular formating allowing to store or transmit data.
Go provides an [encoding/json](https://pkg.go.dev/encoding/json) package that implements encoding and
decoding operations.

In the examples below we use the following structure:

```go
type Entry struct {
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
  {"Jaroslavs", "Samcuks", "+41123456789"},
  {"Pavel", "Zaichenkov", "+41987654321"},
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
