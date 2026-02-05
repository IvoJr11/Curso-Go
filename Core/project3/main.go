package main

import "fmt"

type Contact struct {
	ID                 int
	Name, Email, Phone string
}

var (
	contactList            []Contact
	contactListIndexByName map[string]int
	nextID                 int = 1
)

func init() {
	contactList = make([]Contact, 0)
	contactListIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, ok := contactListIndexByName[name]; ok {
		fmt.Println("Contact already exists")
		return
	}
	newContact := Contact{nextID, name, email, phone}
	nextID++
	contactList = append(contactList, newContact)
	contactListIndexByName[name] = newContact.ID
	fmt.Printf("Contact added successfully: %v\n", name)
}

func findContact(name string) *Contact {
	index, ok := contactListIndexByName[name]
	var ret *Contact = nil
	if ok {
		ret = &contactList[index]
	}
	return ret
}

func main() {
}

func menu() {
	for true {
		switch expression {
		case condition:
		}
	}
}
