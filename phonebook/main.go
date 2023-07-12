package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

func main() {
	phonebook := make(map[string]Contact)

	for {
		fmt.Println("\nPhonebook:")

		fmt.Println("\n1. Add a contact")
		fmt.Println("2. Lookup a contact")
		fmt.Println("3. List all contacts")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addContact(phonebook)
		case 2:
			lookupContact(phonebook)
		case 3:
			printPhonebook(phonebook)
		case 4:
			fmt.Println("Exiting the program...")
			return
		default:
			fmt.Println("Invalid choice, Please try again.")
		}
	}
}

func addContact(phonebook map[string]Contact) {
	var name, phone string

	fmt.Print("Enter contact name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter phone number: ")
	fmt.Scanln(&phone)

	contact := Contact{
		Name:  name,
		Phone: phone,
	}

	phonebook[name] = contact

	fmt.Println("Contact is added successfully!")
}

func lookupContact(phonebook map[string]Contact) {
	var name string

	fmt.Print("Enter contact name: ")
	fmt.Scanln(&name)

	contact, ok := phonebook[name]
	if ok == true {
		fmt.Println("Contact found:")
		fmt.Println("Name:", contact.Name)
		fmt.Println("Phone Number:", contact.Phone)
	} else {
		fmt.Println("Contact not found.")
	}
}

func printPhonebook(phonebook map[string]Contact) {
	if len(phonebook) == 0 {
		fmt.Println("Phonebook is empty")
		return
	}

	fmt.Println("All contacts:")

	for _, value := range phonebook {
		fmt.Println("Name:", value.Name)
		fmt.Println("Phone Number:", value.Phone)
	}
}

// students := map[string]Student{
// 	"smith": Student{
// 		Name:   "smith",
// 		Grades: []int{45, 66, 99},
// 	},

// 	"dave": Student{
// 		Name:   "dave",
// 		Grades: []int{77, 88, 99, 56},
// 	},
// }
