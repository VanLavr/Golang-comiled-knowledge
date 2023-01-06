package main

import (
	"github.com/VanLavr/GCK/CLI"
	"fmt"
	"time"
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

			out1 := make(chan string)
			out2 := make(chan string)
			out3 := make(chan string)
			out4 := make(chan string)

			go IDontKnowHowToUseConcurrencyHere(out1, "saving data...")
			go IDontKnowHowToUseConcurrencyHere(out2, "checking storage...")
			go IDontKnowHowToUseConcurrencyHere(out3, "translating...")
			go IDontKnowHowToUseConcurrencyHere(out4, "i'm dumb...")

			for {
				message1, opened1 := <- out1
				fmt.Println(message1)
				time.Sleep(time.Millisecond * 250)

				for {
					message2, opened2 := <- out2
					fmt.Println(message2)
					time.Sleep(time.Millisecond * 250)
					if !opened2 {
						break
					}
				}

				for {
					message3, opened3 := <- out3
					fmt.Println(message3)
					time.Sleep(time.Millisecond * 250)
					if !opened3 {
						break
					}
				}

				for {
					message4, opened4 := <- out4
					fmt.Println(message4)
					time.Sleep(time.Millisecond * 250)
					if !opened4 {
						break
					}
				}

				if !opened1 {
					break
				}
			}

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

func IDontKnowHowToUseConcurrencyHere(out chan string, processName string) {
	defer close(out)
	out <- fmt.Sprintf("process %s started", processName)
	out <- fmt.Sprintf("process %s finished", processName)
}