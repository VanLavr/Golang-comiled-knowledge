package CLIresps

import (
	"fmt"
	"os"
	"encoding/json"
	"strings"
)

var (
	path string = "C:\\Users\\Ivan\\GoLang\\Golang-compiled-knowledge\\dataBase\\data.json"
)

type Editor interface {
	WriteContact()
	DeleteContact()
	ChangeContact()
}

type Shower interface {
	ShowContactInfo()
	ShowAllContacts()
}


type Contact struct {
	Name    string `json:"NAME"`
	Number  string `json:"NUM"`
	Comment string `json:"COMM"`
	ID      string `json:"ID"`
}

func NewContact() *Contact {
	return new(Contact)
}

func (c *Contact) WriteContact(Nm string, Nmbr string, Cmmnt string, Id string) {
	c.Name = Nm
	c.Number = Nmbr
	c.Comment = Cmmnt
	c.ID = Id

	JSONdata, err1 := json.Marshal(c)
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	file, err2 := os.OpenFile(path, os.O_APPEND | os.O_WRONLY, 0644)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	defer file.Close()

	content, e := os.ReadFile(path)
	if e != nil {
		fmt.Println(e.Error())
	}
	if string(content) != "" {
		_, er := file.Write([]byte("\n"))
		if er != nil {
			fmt.Println(er.Error())
		}
	
		_, err3 := file.Write([]byte(JSONdata))
		if err3 != nil {
			fmt.Println(err3.Error())
		}
	} else {
		_, err3 := file.Write([]byte(JSONdata))
		if err3 != nil {
			fmt.Println(err3.Error())
		}
	}

	fmt.Println("Contact was written...")
}

func (c *Contact) DeleteContact(Id string) {
	var substr string = "\"ID\":\""
	substr += Id
	substr += "\""
	var stringArray []string
	var str string
	
	bytes, err1 := os.ReadFile(path)
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	str = string(bytes)
	stringArray = strings.Split(str, "\n")

	for i := 0; i < len(stringArray); i++ {
		if strings.Contains(stringArray[i], substr) {
			stringArray = append(stringArray[:i], stringArray[i+1:]...)
		}
	}

	wroteString := ""
	for i := 0; i < len(stringArray); i++ {
		wroteString += stringArray[i]
		wroteString += "\n"
	}

	file, err1 := os.Create(path)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	defer file.Close()

	_, err2 := file.WriteString(wroteString)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	fmt.Println("Contact was deleted...")
}

func (c *Contact) ShowContactInfo(Id string) {
	var substr string = "\"ID\":\""
	substr += Id
	substr += "\""
	var stringArray []string
	var str string
	
	bytes, err1 := os.ReadFile(path)
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	str = string(bytes)
	stringArray = strings.Split(str, "\n")

	for i := 0; i < len(stringArray); i++ {
		if strings.Contains(stringArray[i], substr) {
			fmt.Printf(stringArray[i])
		}
	}
}

func (c *Contact) ShowAllContacts() {
	var stringArray []string
	var str string
	
	bytes, err1 := os.ReadFile(path)
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	str = string(bytes)
	stringArray = strings.Split(str, "\n")

	for i := 0; i < len(stringArray); i++ {
		fmt.Println(stringArray[i])
	}
}



func (c *Contact) ChangeContact() {}
