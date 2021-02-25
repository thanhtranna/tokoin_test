package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"tokoin-challenge/config"
	"tokoin-challenge/src/common"
	"tokoin-challenge/src/entity"
)

// ITicketRepository : ticket interface
type ITicketRepository interface {
	List(key, value string) (*entity.Tickets, error)
	ListSubjects(key, value string) ([]string, error)
}

type TicketRepo struct {
	tickets entity.Tickets
}

func NewTicketRepository() ITicketRepository {
	ticketRepo := TicketRepo{}
	ticketRepo.LoadDataFromFile(config.Config.Data.Ticket)
	return &ticketRepo
}

func (r *TicketRepo) LoadDataFromFile(path string) error {
	data, err := common.ReadJSONFile(path)
	if err != nil {
		return fmt.Errorf("cannot load data from json file %s", path)
	}
	return r.LoadDataFromBytes(data)
}

func (r *TicketRepo) LoadDataFromBytes(data []byte) error {
	var tickets entity.Tickets
	err := json.Unmarshal(data, &tickets)
	if err != nil {
		return err
	}
	r.tickets = tickets

	return nil
}

func (r *TicketRepo) List(key, value string) (*entity.Tickets, error) {
	results := entity.Tickets{}
	switch key {
	case "_id":
		for _, ticket := range r.tickets {
			if ticket.ID == value {
				results = append(results, ticket)
			}
		}
	case "url":
		for _, org := range r.tickets {
			if org.URL == value {
				results = append(results, org)
			}
		}
	case "external_id":
		for _, org := range r.tickets {
			if org.ExternalID == value {
				results = append(results, org)
			}
		}
	case "created_at":
		for _, org := range r.tickets {
			if org.CreatedAt == value {
				results = append(results, org)
			}
		}
	case "type":
		for _, org := range r.tickets {
			if org.Type == value {
				results = append(results, org)
			}
		}
	case "subject":
		for _, org := range r.tickets {
			if org.Subject == value {
				results = append(results, org)
			}
		}
	case "description":
		for _, org := range r.tickets {
			if org.Description == value {
				results = append(results, org)
			}
		}
	case "priority":
		for _, org := range r.tickets {
			if org.Priority == value {
				results = append(results, org)
			}
		}
	case "status":
		for _, org := range r.tickets {
			if org.Status == value {
				results = append(results, org)
			}
		}
	case "submitter_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return &results, errors.New("input submitter_id is invalid")
		}

		for _, org := range r.tickets {
			if org.SubmitterID == id {
				results = append(results, org)
			}
		}
	case "assignee_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return &results, errors.New("input assignee_id is invalid")
		}

		for _, org := range r.tickets {
			if org.AssigneeID == id {
				results = append(results, org)
			}
		}
	case "organization_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return &results, errors.New("input organization_id is invalid")
		}

		for _, org := range r.tickets {
			if org.OrganizationID == id {
				results = append(results, org)
			}
		}
	case "tags":
		for _, org := range r.tickets {
			for _, tag := range org.Tags {
				if tag == value {
					results = append(results, org)
					break
				}
			}
		}
	case "has_incidents":
		v, err := common.StringToBoolean(value)
		if err != nil {
			return &results, err
		}

		for _, org := range r.tickets {
			if org.HasIncidents == v {
				results = append(results, org)
			}
		}
	case "due_at":
		for _, org := range r.tickets {
			if org.DueAt == value {
				results = append(results, org)
			}
		}
	case "via":
		for _, org := range r.tickets {
			if org.Via == value {
				results = append(results, org)
			}
		}
	default:
		return &results, errors.New("key is invalid")
	}

	return &results, nil
}

func (r *TicketRepo) ListSubjects(key, value string) ([]string, error) {
	rs := []string{}
	tickets, err := r.List(key, value)
	if err != nil {
		return nil, err
	}

	for _, t := range *tickets {
		rs = append(rs, t.Subject)
	}

	return rs, nil
}
