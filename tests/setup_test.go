package tests

import (
	"tokoin-challenge/src/repositories"
	"tokoin-challenge/src/services"
)

const (
	sampleFilePath = "data/sample.json"
)

type SearchArgs struct {
	Key   string
	Value string
}

type TestCase struct {
	Name           string
	Args           interface{}
	ExpectedResult interface{}
	ExpectedError  bool
}

type SearchTestCase struct {
	Name           string
	Args           SearchArgs
	ExpectedResult interface{}
	ExpectedError  bool
}

var mockOrgRepo repositories.IOrgRepository
var mockTicketRepo repositories.ITicketRepository
var mockUserRepo repositories.IUserRepository

var mockOrgService services.IOrgService
var mockTicketService services.ITicketService
var mockUserService services.IUserService

func NewOrgRepo() repositories.IOrgRepository {
	mockOrgRepo := &repositories.OrganizationRepo{}
	mockOrgRepo.LoadDataFromBytes([]byte(mockDataOrg))

	return mockOrgRepo
}

func NewTicketRepo() repositories.ITicketRepository {
	mockTicketRepo := &repositories.TicketRepo{}
	mockTicketRepo.LoadDataFromBytes([]byte(mockDataTickets))

	return mockTicketRepo
}

func NewUserRepo() repositories.IUserRepository {
	mockUserRepo := &repositories.UserRepo{}
	mockUserRepo.LoadDataFromBytes([]byte(mockDataUsers))

	return mockUserRepo
}

func init() {
	mockOrgRepo = NewOrgRepo()
	mockTicketRepo = NewTicketRepo()
	mockUserRepo = NewUserRepo()

	mockOrgService = services.NewOrgService(mockOrgRepo, mockTicketRepo, mockUserRepo)
	mockTicketService = services.NewTicketService(mockOrgRepo, mockTicketRepo, mockUserRepo)
	mockUserService = services.NewUserService(mockOrgRepo, mockTicketRepo, mockUserRepo)
}
