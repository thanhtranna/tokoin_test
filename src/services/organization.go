package services

import (
	"strconv"

	"github.com/jinzhu/copier"

	"tokoin-challenge/src/entity"
	"tokoin-challenge/src/repositories"
)

// IOrgService : organization interface service
type IOrgService interface {
	List(key, value string) (*entity.OrganizationsResponse, error)
}

type OrgService struct {
	orgRepo    repositories.IOrgRepository
	ticketRepo repositories.ITicketRepository
	userRepo   repositories.IUserRepository
}

func NewOrgService(orgRepo repositories.IOrgRepository, ticketRepo repositories.ITicketRepository,
	userRepo repositories.IUserRepository) IOrgService {
	return &OrgService{
		orgRepo:    orgRepo,
		ticketRepo: ticketRepo,
		userRepo:   userRepo,
	}
}

func (s *OrgService) List(key, value string) (*entity.OrganizationsResponse, error) {
	orgs, err := s.orgRepo.List(key, value)
	if err != nil {
		return nil, err
	}
	results := entity.OrganizationsResponse{}
	for _, org := range *orgs {
		var rs entity.OrganizationResponse
		copier.Copy(&rs, &org)
		strOrgID := strconv.Itoa(org.ID)

		// Get tickets of organization
		rs.TicketSubjects, _ = s.ticketRepo.ListSubjects("organization_id", strOrgID)

		// Get user names of organization
		rs.UserNames, _ = s.userRepo.ListNames("organization_id", strOrgID)

		results = append(results, &rs)
	}

	return &results, nil
}
