package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	command := os.Args[1]
	userBookName := os.Args[2:]
	commandList := []string{
		"help",
		"list",
		"search",
	}
	ifCommandExists := containsCommand(commandList, command)
	if ifCommandExists == false {
		fmt.Println("Enter an available command from list:", commandList)
	} else if command == "help" {
		fmt.Println("Available commands: ", commandList[1:])
		fmt.Println("\nDescriptions")
		fmt.Println("------------")
		fmt.Println("list : Lists all the books the application has.")
		fmt.Println("search : Searches if the wanted book is in the application. Use as -> search <bookName>")
	}

	switch command {
	case "list":
		listAll(bookList)
	case "search":
		res, _, bookRes := searchForBook(strings.Join(userBookName, " "), &bookList)
		if res == true {
			fmt.Println("Application has the book: ", bookRes.name, "by", bookRes.author)
		} else {
			fmt.Println("Book not found in application!")
		}
	}

}

//----------------------------------------------------------------------------------------------------------------------

// containsCommand function checks if a string is present in a slice of strings. Returns bool
func containsCommand(l []string, t string) bool {
	for _, v := range l {
		if v == t {
			return true
		}
	}
	return false
}

type book struct {
	name   string
	author string
}

// bookTag method function prints the name and author fields of the struct.
func (b *book) bookTag() {
	fmt.Println("Book Name: ", b.name)
	fmt.Println("Author: ", b.author)
}

// listAll function lists all the books in a slice of book structs.
func listAll(l []book) {
	for _, v := range l {
		v.bookTag()
		fmt.Println("----------")
	}
}

// searchForBook function looks for the query sting in a slice of book structs to return a bool result, integer index and the struct book.
func searchForBook(query string, bookList *[]book) (bool, int, *book) {
	for i, v := range *bookList {
		if strings.ToUpper(query) == strings.ToUpper(v.name) {
			fmt.Println("You have a match!")
			return true, i, &v
		}
	}
	return false, -1, nil
}

var bookList = []book{
	{
		name:   "Lord of the Rings",
		author: "J.R.R. Tolkien"},
	{
		name:   "Dispossessed",
		author: "Ursula K. Le Guin"},
	{
		name:   "The Devil Inside Us",
		author: "Sabahattin Ali"},
	{
		name:   "Prisoners of Ourselves",
		author: "Gunduz Vassaf"},
	{
		name:   "The Odyssey",
		author: "Homer"},
	{
		name:   "Cosmos",
		author: "Carl Sagan"},
	{
		name:   "Animal Farm",
		author: "George Orwell"},
	{
		name:   "Martin Eden",
		author: "Jack London"},
}
