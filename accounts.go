package form3

import (
	"context"
	"fmt"
)

const path = "/v1/organisation/accounts"

type AccountService service

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

type AccountRoot struct {
	Data  *AccountData `json:"data"`
	Links *Links       `json:"links,omitempty"`
}

// Get a single account using the account ID.
// https://api-docs.form3.tech/api.html#organisation-accounts-fetch
func (s *AccountService) Fetch(ctx context.Context, id string) (*AccountRoot, *Response, error) {
	path := fmt.Sprintf("%s/%s", path, id)

	req, err := s.client.GET(ctx, path, nil)

	if err != nil {
		return nil, nil, err
	}

	root := new(AccountRoot)
	resp, err := s.client.Do(ctx, req, root)

	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Register an existing bank account with Form3 or create a new one.
// https://api-docs.form3.tech/api.html#organisation-accounts-create
func (s *AccountService) Create(ctx context.Context, account *AccountData) (*AccountRoot, *Response, error) {
	reqBody := AccountRoot{Data: account}
	req, err := s.client.POST(path, reqBody)

	if err != nil {
		return nil, nil, err
	}

	root := new(AccountRoot)
	resp, err := s.client.Do(ctx, req, root)

	return root, resp, err
}

// Delete an account by id and version
// https://api-docs.form3.tech/api.html#organisation-accounts-delete
func (s *AccountService) Delete(ctx context.Context, id string, version int64) (*Response, error) {

	path := fmt.Sprintf("%s/%s?version=%d", path, id, version)
	req, err := s.client.DELETE(path)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}
