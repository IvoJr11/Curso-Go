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

func listContacts() {
	if len(contactList) == 0 {
		fmt.Println("Contact List is empty.")
	} else {
		for _, contact := range contactList {
			fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
		}
	}
}

func main() {
	addContact("Mario", "mario@gmail.com", "299432094")
	addContact("Juan", "juan@gmail.com", "299453421")
	addContact("Luis", "luis@gmail.com", "29976532")
	addContact("Osvaldo", "osvaldo@gmail.com", "29921362")
	addContact("Lucas", "lucas@gmail.com", "2995465413")

	listContacts()
}
