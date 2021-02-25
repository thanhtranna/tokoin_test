package tests

import (
	"testing"

	"tokoin-challenge/src/entity"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationEntityToString(t *testing.T) {
	testcases := []TestCase{
		{"Organization entity to string", nil, 1, false},
	}

	orgEntity := entity.Organization{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := orgEntity.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestTicketEntityToString(t *testing.T) {
	testcases := []TestCase{
		{"Ticket entity to string", nil, 1, false},
	}

	ticketEntity := entity.Ticket{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := ticketEntity.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestUserEntityToString(t *testing.T) {
	testcases := []TestCase{
		{"Ticket entity to string", nil, 1, false},
	}

	userEntity := entity.User{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := userEntity.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestOrganizationResponseToString(t *testing.T) {
	testcases := []TestCase{
		{"OrganizationResponse entity to string", nil, 1, false},
	}

	orgResponse := entity.OrganizationResponse{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := orgResponse.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestOrganizationsResponseToString(t *testing.T) {
	testcases := []TestCase{
		{"OrganizationsResponse entity to string", nil, 1, false},
	}

	orgsResponse := entity.OrganizationsResponse{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := orgsResponse.ToString()
			assert.Empty(t, str)
		})
	}
}

func TestTicketResponseToString(t *testing.T) {
	testcases := []TestCase{
		{"TicketResponse entity to string", nil, 1, false},
	}

	ticketResponse := entity.TicketResponse{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := ticketResponse.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestTicketsResponseToString(t *testing.T) {
	testcases := []TestCase{
		{"TicketResponse entity to string", nil, 1, false},
	}

	ticketsResponse := entity.TicketsResponse{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := ticketsResponse.ToString()
			assert.Empty(t, str)
		})
	}
}

func TestUserResponseToString(t *testing.T) {
	testcases := []TestCase{
		{"UserResponse entity to string", nil, 1, false},
	}

	userResponse := entity.UserResponse{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := userResponse.ToString()
			assert.NotEmpty(t, str)
		})
	}
}

func TestUsersResponseToString(t *testing.T) {
	testcases := []TestCase{
		{"TicketResponse entity to string", nil, 1, false},
	}

	usersResponse := entity.UsersResponse{}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			str := usersResponse.ToString()
			assert.Empty(t, str)
		})
	}
}
