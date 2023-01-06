package main

import (
	"github.com/VanLavr/GCK/CLI"
	"fmt"
)

func main() {
	for {
		var (
			choice string
			manager CLIresps.Editor
		)
		contact := CLIresps.Contact{}
		manager = &contact

		fmt.Println("_____________________________________________________________________________________________|_________________________________________")
		fmt.Println("0 - Exit | 1 - Show all | 2 - Show by ID | 3 - Create a contact | 4 - Delete a contact by ID | please use ONE (1) word for a comment...")
		fmt.Println("`````````````````````````````````````````````````````````````````````````````````````````````|`````````````````````````````````````````")
		fmt.Scan(&choice)
		if choice == "0" {
			break
		} else {
			switch choice {
			case "1":
				manager.ShowAllContacts()
			case "2":
				var id string
				fmt.Println("Please, type an ID:")
				fmt.Scan(&id)
				manager.ShowContactInfo(id)
				fmt.Println()
			case "3":
				var (
					name string
					number string
					comment string
					id string
				)
				fmt.Println("Please, type Name, number, comment, ID:")
				fmt.Scan(&name, &number, &comment, &id)
				manager.WriteContact(name, number, comment, id)
			case "4":
				var id string
				fmt.Println("Please, type an ID:")
				fmt.Scan(&id)
				manager.DeleteContact(id)
				fmt.Println()
			default:
				fmt.Println("/^\\                        /^\\")
				fmt.Println(" | please type number (0-4) |")
				fmt.Println(" |                          |")
			}
		}
	}
}