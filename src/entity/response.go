package entity

type ListResult interface {
	ToString() string
	Length() int
}

// OrganizationResponse struct
type OrganizationResponse struct {
	Organization   `json:",inline"`
	TicketSubjects []string `json:"ticket_subjects"`
	UserNames      []string `json:"user_names"`
}

// OrganizationsResponse : slice OrganizationResponse
type OrganizationsResponse []*OrganizationResponse

func (o OrganizationsResponse) ToString() string {
	rs := ""
	for _, e := range o {
		if e != nil {
			rs += e.ToString() + "\n-------------------------------------------\n"
		}
	}
	return rs
}

func (o OrganizationsResponse) Length() int {
	return len(o)
}

// TicketResponse struct
type TicketResponse struct {
	Ticket           `json:",inline"`
	AssigneeName     string `json:"assignee_name"`
	SubmitterName    string `json:"submitter_name"`
	OrganizationName string `json:"organization_name"`
}

// TicketsResponse : slice TicketResponse
type TicketsResponse []*TicketResponse

func (o TicketsResponse) ToString() string {
	rs := ""
	for _, e := range o {
		if e != nil {
			rs += e.ToString() + "\n-------------------------------------------\n"
		}
	}
	return rs
}

func (o TicketsResponse) Length() int {
	return len(o)
}

// UserResponse struct
type UserResponse struct {
	User                    `json:",inline"`
	AssigneeTicketSubjects  []string `json:"assignee_ticket_subjects"`
	SubmittedTicketSubjects []string `json:"submitted_ticket_subjects"`
	OrganizationName        string   `json:"organization_name"`
}

// UsersResponse : slice UserResponse
type UsersResponse []*UserResponse

func (o UsersResponse) ToString() string {
	rs := ""
	for _, e := range o {
		if e != nil {
			rs += e.ToString() + "\n-------------------------------------------\n"
		}
	}
	return rs
}

func (o UsersResponse) Length() int {
	return len(o)
}
