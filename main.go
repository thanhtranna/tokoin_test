package main

import (
	"bufio"
	"fmt"
	"os"

	"tokoin-challenge/src/entity"
	"tokoin-challenge/src/repositories"
	"tokoin-challenge/src/services"
)

const welcomeMessage = `Type 'quit' to exit at any time. Select search options:
 	* Enter '1' to search.	
 	* Enter '2' to view a list of searchable fields.
`

const chooseModelMessage = `Please choose model you want to search. Select options:
	 * Enter '1' for Organization.
	 * Enter '2' for Ticket.
	 * Enter '3' for User.
`

const inputFieldMessage = `Enter search term: `
const inputValueMessage = `Enter search value: `

var scanner = bufio.NewScanner(os.Stdin)

func Run(orgSrv services.IOrgService, userSrv services.IUserService, ticketSrv services.ITicketService) {
	for {
		fmt.Print(welcomeMessage)
		cmd, quit := readInput()
		if quit {
			break
		}

		if quit = commandHandler(orgSrv, userSrv, ticketSrv, cmd); quit {
			break
		}
	}
}

func commandHandler(orgSrv services.IOrgService, userSrv services.IUserService, ticketSrv services.ITicketService, cmd string) bool {
	switch cmd {
	case "1":
		return searchHandler(orgSrv, userSrv, ticketSrv)
	case "2":
		showSearchableFields()
		return false
	}

	return false
}

func searchHandler(orgSrv services.IOrgService, userSrv services.IUserService, ticketSrv services.ITicketService) bool {
	var (
		searchModel string
		searchField string
		searchValue string
		quit        bool
	)

	fmt.Print(chooseModelMessage)
	for {
		line, quit := readInput()
		if quit {
			return true
		}

		if line == "1" || line == "2" || line == "3" {
			searchModel = line
			break
		}
		fmt.Print("Please enter '1', '2' or '3':")
	}

	fmt.Print(inputFieldMessage)
	searchField, quit = readInput()
	if quit {
		return true
	}

	fmt.Print(inputValueMessage)
	searchValue, quit = readInput()
	if quit {
		return true
	}

	var results entity.ListResult
	var err error

	switch searchModel {
	case "1":
		results, err = orgSrv.List(searchField, searchValue)
	case "2":
		results, err = ticketSrv.List(searchField, searchValue)
	case "3":
		results, err = userSrv.List(searchField, searchValue)
	}

	if err != nil {
		fmt.Println("Error occurred: ", err)
		return false
	}

	fmt.Print("-------------------------------------------")
	fmt.Println(results.ToString())
	fmt.Printf("Found %d record(s).\n", results.Length())

	return false
}

func showSearchableFields() {
	userSearchableFields := []string{"_id", "url", "external_id", "name", "alias", "created_at",
		"active", "verified", "shared", "locale", "timezone", "last_login_at", "email",
		"phone", "signature", "organization_id", "tags", "suspended", "role",
	}
	ticketSearchableFields := []string{"_id", "url", "external_id", "created_at",
		"type", "subject", "description", "priority", "status", "submitter_id",
		"assignee_id", "organization_id", "tags", "has_incidents", "due_at", "via",
	}
	orgSearchableFields := []string{
		"_id", "url", "external_id", "name", "domain_names", "created_at", "details", "shared_tickets", "tags",
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Search user with")
	for _, e := range userSearchableFields {
		fmt.Println(e)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Search ticket with")
	for _, e := range ticketSearchableFields {
		fmt.Println(e)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Search organization with")
	for _, e := range orgSearchableFields {
		fmt.Println(e)
	}

	fmt.Println("------------------------------------------")
}

func readInput() (string, bool) {
	scanner.Scan()
	line := scanner.Text()
	if line == "quit" {
		return "", true
	}

	return line, false
}

func main() {
	fmt.Println("Hello World!")

	// import repository
	userRepo := repositories.NewUserRepository()
	orgRepo := repositories.NewOrgRepository()
	ticketRepo := repositories.NewTicketRepository()

	// import services

	userSrv := services.NewUserService(orgRepo, ticketRepo, userRepo)
	orgSrv := services.NewOrgService(orgRepo, ticketRepo, userRepo)
	ticketSrv := services.NewTicketService(orgRepo, ticketRepo, userRepo)

	Run(orgSrv, userSrv, ticketSrv)
}
