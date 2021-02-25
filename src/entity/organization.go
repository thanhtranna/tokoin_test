package entity

import "tokoin-challenge/src/common"

// Organization struct
type Organization struct {
	ID            int      `json:"_id"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}

// Organizations : slice of Organization
type Organizations []*Organization

// ToString : stringtify Organization
func (o Organization) ToString() string {
	return common.Jsonify(o)
}
