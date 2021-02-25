package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tokoin-challenge/src/repositories"
)

const invalidDataOrg = `org`
const mockDataOrg = `
[
	{
        "_id": 124,
        "url": "http://initech.tokoin.io.com/api/v2/organizations/124.json",
        "external_id": "15c21605-cbc6-440f-8da2-6e1601aed5fa",
        "name": "Bitrex",
        "domain_names": [
            "unisure.com",
            "boink.com",
            "quinex.com",
            "poochies.com"
        ],
        "created_at": "2016-05-11T12:16:15 -10:00",
        "details": "Non profit",
        "shared_tickets": true,
        "tags": [
            "Lott",
            "Hunter",
            "Beasley",
            "Glass"
        ]
    },
	{
        "_id": 125,
        "url": "http://initech.tokoin.io.com/api/v2/organizations/125.json",
        "external_id": "42a1a845-70cf-40ed-a762-acb27fd606cc",
        "name": "Strezzö",
        "domain_names": [
            "techtrix.com",
            "teraprene.com",
            "corpulse.com",
            "flotonic.com"
        ],
        "created_at": "2016-02-21T06:11:51 -11:00",
        "details": "MegaCorp",
        "shared_tickets": false,
        "tags": [
            "Vance",
            "Ray",
            "Jacobs",
            "Frank"
        ]
    }
]`

func TestOrganizationRepoLoadDataFromFile(t *testing.T) {
	testcases := []TestCase{
		{"Load from existed file", sampleFilePath, nil, false},
		{"Load from not existed file", "", nil, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			repo := &repositories.OrganizationRepo{}
			err := repo.LoadDataFromFile(testcase.Args.(string))
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestOrganizationRepoLoadDataFromBytes(t *testing.T) {
	testcases := []TestCase{
		{"Load from valid json bytes", mockDataOrg, nil, false},
		{"Load from invalid json bytes", invalidDataOrg, nil, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			repo := &repositories.OrganizationRepo{}
			err := repo.LoadDataFromBytes([]byte(testcase.Args.(string)))
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestOrganizationRepoListExistedRecord(t *testing.T) {
	mockOrgRepo := &repositories.OrganizationRepo{}
	mockOrgRepo.LoadDataFromBytes([]byte(mockDataOrg))

	testcases := []SearchTestCase{
		// search existed record.
		{"Search by existed _id", SearchArgs{"_id", "124"}, 1, false},
		{"Search by existed url", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/organizations/124.json"}, 1, false},
		{"Search by existed external_id", SearchArgs{"external_id", "15c21605-cbc6-440f-8da2-6e1601aed5fa"}, 1, false},
		{"Search by existed name", SearchArgs{"name", "Bitrex"}, 1, false},
		{"Search by existed domain_names", SearchArgs{"domain_names", "unisure.com"}, 1, false},
		{"Search by existed created_at", SearchArgs{"created_at", "2016-05-11T12:16:15 -10:00"}, 1, false},
		{"Search by existed details", SearchArgs{"details", "Non profit"}, 1, false},
		{"Search by existed shared_tickets", SearchArgs{"shared_tickets", "true"}, 1, false},
		{"Search by existed tags", SearchArgs{"tags", "Hunter"}, 1, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockOrgRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestOrganizationRepoListNotExistedRecord(t *testing.T) {
	mockOrgRepo := &repositories.OrganizationRepo{}
	mockOrgRepo.LoadDataFromBytes([]byte(mockDataOrg))

	testcases := []SearchTestCase{
		// search not existed record..
		{"Search by not existed _id", SearchArgs{"_id", "111"}, 0, false},
		{"Search by not existed url", SearchArgs{"url", "http://initech.tokoin.io.com/api/v2/organizations/111.json"}, 0, false},
		{"Search by not existed external_id", SearchArgs{"external_id", "9270ed79-35eb-4a38-a46f-35725197ea11"}, 0, false},
		{"Search by not existed name", SearchArgs{"name", "Enthaze11"}, 0, false},
		{"Search by not existed domain_names", SearchArgs{"domain_names", "zentix11.com"}, 0, false},
		{"Search by not existed created_at", SearchArgs{"created_at", "2016-05-21T11:11:11 -10:00"}, 0, false},
		{"Search by not existed details", SearchArgs{"details", "Non profit11"}, 0, false},
		{"Search by not existed shared_tickets", SearchArgs{"shared_tickets", "true"}, 1, false},
		{"Search by not existed tags", SearchArgs{"tags", "Collier11"}, 0, false},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockOrgRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}

func TestOrganizationRepoListInvalidInput(t *testing.T) {
	mockOrgRepo := &repositories.OrganizationRepo{}
	mockOrgRepo.LoadDataFromBytes([]byte(mockDataOrg))

	testcases := []SearchTestCase{
		// search by invalid input.
		{"Search by invalid _id", SearchArgs{"_id", "id"}, 0, true},
		{"Search by invalid shared_tickets", SearchArgs{"shared_tickets", "fasdf"}, 0, true},
		{"Search by invalid key", SearchArgs{"key", "value"}, 0, true},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			results, err := mockOrgRepo.List(testcase.Args.Key, testcase.Args.Value)
			assert.NotNil(t, results, err)
			assert.Equal(t, testcase.ExpectedResult, len(*results), err)
			assert.Equal(t, testcase.ExpectedError, err != nil, err)
		})
	}
}
