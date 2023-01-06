package main

import (
	"github.com/VanLavr/GCK/CLI"
	"fmt"
)

func main() {
	var a CLIresps.Editor
	contact := CLIresps.Contact{}
	a = &contact



	a.WriteContact("Ivan", "112343", "gg", "1")
	a.WriteContact("Vasya", "243213", "wp", "2")
	a.WriteContact("Asya", "666323", "lol", "3")
	a.WriteContact("Van", "666313", "lmao", "4")
	a.WriteContact("Nastya", "666353", "rofl", "5")
	a.WriteContact("Gosha", "666383", "l", "6")
	a.ShowAllContacts()
	a.DeleteContact("2")
	a.WriteContact("Vasya", "243213", "wp", "2")
	a.ShowContactInfo("3")
	fmt.Println()
	fmt.Println()
	a.ShowAllContacts()
	a.WriteContact("Vasya", "243213", "wp", "16")
	a.DeleteContact("2")
	a.ShowAllContacts()
}