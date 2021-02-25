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

// IOrgRepository : organization interface
type IOrgRepository interface {
	Retrieve(id int) (*entity.Organization, error)
	List(key, value string) (*entity.Organizations, error)
}

type OrganizationRepo struct {
	organizations entity.Organizations
	orgIDMap      map[int]*entity.Organization
}

func NewOrgRepository() IOrgRepository {
	orgRepo := OrganizationRepo{}
	err := orgRepo.LoadDataFromFile(config.Config.Data.Organization)
	if err != nil {
		fmt.Println("Cannot load data, error: ", err)
	}

	return &orgRepo
}

func (r *OrganizationRepo) LoadDataFromFile(path string) error {
	data, err := common.ReadJSONFile(path)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("cannot load data from json file %s", path))
	}
	r.LoadDataFromBytes(data)

	return nil
}

func (r *OrganizationRepo) LoadDataFromBytes(data []byte) error {
	var orgs entity.Organizations
	err := json.Unmarshal(data, &orgs)
	if err != nil {
		return err
	}
	r.organizations = orgs
	r.orgIDMap = map[int]*entity.Organization{}
	for _, org := range orgs {
		r.orgIDMap[org.ID] = org
	}

	return nil
}

func (r *OrganizationRepo) Retrieve(id int) (*entity.Organization, error) {
	return r.orgIDMap[id], nil
}

func (r *OrganizationRepo) List(key, value string) (*entity.Organizations, error) {
	results := entity.Organizations{}
	switch key {
	case "_id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return &results, errors.New("input _id is invalid")
		}
		org, _ := r.Retrieve(id)
		if org != nil {
			results = append(results, org)
		}
	case "url":
		for _, org := range r.organizations {
			if org.URL == value {
				results = append(results, org)
			}
		}
	case "external_id":
		for _, org := range r.organizations {
			if org.ExternalID == value {
				results = append(results, org)
			}
		}
	case "name":
		for _, org := range r.organizations {
			if org.Name == value {
				results = append(results, org)
			}
		}
	case "domain_names":
		for _, org := range r.organizations {
			for _, d := range org.DomainNames {
				if d == value {
					results = append(results, org)
					break
				}
			}
		}
	case "created_at":
		for _, org := range r.organizations {
			if org.CreatedAt == value {
				results = append(results, org)
			}
		}
	case "details":
		for _, org := range r.organizations {
			if org.Details == value {
				results = append(results, org)
			}
		}
	case "shared_tickets":
		v, err := common.StringToBoolean(value)
		if err != nil {
			return &results, err
		}

		for _, org := range r.organizations {
			if org.SharedTickets == v {
				results = append(results, org)
			}
		}
	case "tags":
		for _, org := range r.organizations {
			for _, tag := range org.Tags {
				if tag == value {
					results = append(results, org)
					break
				}
			}
		}
	default:
		return &results, errors.New("key is invalid")
	}

	return &results, nil
}
