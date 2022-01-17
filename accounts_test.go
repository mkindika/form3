//go:build integration
// +build integration

package form3

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCreateNewAccount(t *testing.T) {
	c := NewClient()
	ctx := context.Background()

	accountData := setupTestAcountData()
	uuid := accountData.ID
	version := accountData.Version

	t.Cleanup(func() {
		c.Account.Delete(ctx, uuid, *version)
	})

	response, _, _ := c.Account.Create(ctx, accountData)

	assert.NotNil(t, response.Data, "expecting non-nil result")
	if !reflect.DeepEqual(response.Data, accountData) {
		t.Errorf("Account.Get returned %+v, expected %+v", response.Data, accountData)
	}
}

func TestInvalidAccountID(t *testing.T) {
	c := NewClient()
	ctx := context.Background()

	accountData := setupTestAcountData()
	accountData.ID = "invalid uuid"

	_, _, err := c.Account.Create(ctx, accountData)

	expected := "validation failure list:\nvalidation failure list:\nid in body must be of type uuid: \"invalid uuid\""
	actual := err.(*ErrorResponse).Message

	assert.NotNil(t, err, "expecting non-nil result")
	assert.NotNil(t, err, "expecting non-nil result")
	assert.Equal(t, actual, expected, "Invalid validation message")
}

func TestFetchAccountById(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	accountData := setupTestAcountData()
	uuid := accountData.ID
	version := accountData.Version

	t.Cleanup(func() {
		c.Account.Delete(ctx, uuid, *version)
	})

	// Create an account first
	c.Account.Create(ctx, accountData)

	response, _, _ := c.Account.Fetch(ctx, uuid)

	assert.NotNil(t, response, "expecting non nil result")
	if !reflect.DeepEqual(response.Data, accountData) {
		t.Errorf("Account.Get returned %+v, expected %+v", response.Data, accountData)
	}
}

func TestDeleteAccountByIdAndVersion(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	accountData := setupTestAcountData()
	uuid := accountData.ID
	version := accountData.Version

	// Create an account first
	c.Account.Create(ctx, accountData)

	response, _ := c.Account.Delete(ctx, uuid, *version)

	assert.NotNil(t, response, "expecting non-nil result")
	assert.Equal(t, response.StatusCode, 204, "Status code should be 204")
}

func setupTestAcountData() *AccountData {

	accountData := &AccountData{
		Attributes: &AccountAttributes{
			AccountClassification:   func() *string { x := "Personal"; return &x }(),
			AccountMatchingOptOut:   func() *bool { x := false; return &x }(),
			AccountNumber:           "8053530741",
			AlternativeNames:        []string{"Hugo Boss AG"},
			BankID:                  "400300",
			BankIDCode:              "GBDSC",
			BaseCurrency:            "GBP",
			Bic:                     "NWBKGB22",
			Country:                 func() *string { x := "GB"; return &x }(),
			Iban:                    "AG12",
			JointAccount:            func() *bool { x := false; return &x }(),
			Name:                    []string{"Hugo Boss"},
			SecondaryIdentification: "AB46ERR",
			Status:                  func() *string { x := "pending"; return &x }(),
			Switched:                func() *bool { x := false; return &x }(),
		},
		ID:             uuid.NewString(),
		OrganisationID: "055fff08-76a7-11ec-90d6-0242ac120003",
		Type:           "accounts",
		Version:        new(int64),
	}

	return accountData
}
